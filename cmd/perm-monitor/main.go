package main

import (
	"net"
	"os"

	"github.com/cactus/go-statsd-client/statsd"
	flags "github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"code.cloudfoundry.org/lager"

	"crypto/x509"

	"context"

	"strconv"

	"sync"

	"time"

	"code.cloudfoundry.org/perm-go"
	"code.cloudfoundry.org/perm/cmd"
	"code.cloudfoundry.org/perm/messages"
	"code.cloudfoundry.org/perm/monitor"
)

type options struct {
	Perm permOptions `group:"Perm" namespace:"perm"`

	StatsD statsDOptions `group:"StatsD" namespace:"statsd"`

	Logger cmd.LagerFlag

	QueryProbe probeOptions `group:"Query probe" namespace:"query-probe"`

	AdminProbe probeOptions `group:"Admin probe" namespace:"admin-probe"`
}

type permOptions struct {
	Hostname      string                 `long:"hostname" description:"Hostname used to resolve the address of Perm" required:"true"`
	Port          int                    `long:"port" description:"Port used to connect to Perm" required:"true"`
	CACertificate []cmd.FileOrStringFlag `long:"ca-certificate" description:"File path of Perm's CA certificate"`
}

type statsDOptions struct {
	Hostname string `long:"hostname" description:"Hostname used to connect to StatsD server" required:"true"`
	Port     int    `long:"port" description:"Port used to connect to StatsD server" required:"true"`
}

type probeOptions struct {
	Interval time.Duration `long:"interval" description:"Frequency with which the probe is issued" default:"5s"`
	Timeout  time.Duration `long:"timeout" description:"Time after which the probe is considered to have failed" default:"100ms"`
}

func main() {
	parserOpts := &options{}
	parser := flags.NewParser(parserOpts, flags.Default)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	logger, _ := parserOpts.Logger.Logger("perm-monitor")

	logger.Debug(messages.Starting)
	defer logger.Debug(messages.Finished)

	//////////////////////
	// Setup StatsD Client
	statsDAddr := net.JoinHostPort(parserOpts.StatsD.Hostname, strconv.Itoa(parserOpts.StatsD.Port))
	statsDClient, err := statsd.NewBufferedClient(statsDAddr, "", 0, 0)
	if err != nil {
		logger.Fatal(messages.FailedToConnectToStatsD, err, lager.Data{
			"addr": statsDAddr,
		})
		os.Exit(1)
	}
	defer statsDClient.Close()
	//////////////////////

	//////////////////////
	// Setup Perm GRPC Client
	//
	//// Setup TLS Credentials
	pool := x509.NewCertPool()

	for _, certPath := range parserOpts.Perm.CACertificate {
		caPem, e := certPath.Bytes(cmd.InjectableOS{}, cmd.InjectableIOReader{})
		if e != nil {
			logger.Fatal(messages.FailedToReadCertificate, e, lager.Data{
				"location": certPath,
			})
			os.Exit(1)
		}

		if ok := pool.AppendCertsFromPEM(caPem); !ok {
			logger.Fatal(messages.FailedToAppendCertToPool, e, lager.Data{
				"location": certPath,
			})
			os.Exit(1)
		}
	}

	addr := net.JoinHostPort(parserOpts.Perm.Hostname, strconv.Itoa(parserOpts.Perm.Port))
	creds := credentials.NewClientTLSFromCert(pool, parserOpts.Perm.Hostname)

	//// Setup GRPC connection
	g, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Fatal(messages.FailedToGRPCDial, err, lager.Data{
			"addr": addr,
		})
		os.Exit(1)
	}
	defer g.Close()

	roleServiceClient := protos.NewRoleServiceClient(g)
	permissionServiceClient := protos.NewPermissionServiceClient(g)
	//////////////////////

	ctx := context.Background()

	adminProbe := &monitor.AdminProbe{
		RoleServiceClient: roleServiceClient,
	}

	queryProbe := &monitor.QueryProbe{
		RoleServiceClient:       roleServiceClient,
		PermissionServiceClient: permissionServiceClient,
	}

	queryProbeHistogram := monitor.NewThreadSafeHistogram(
		QueryProbeHistogramWindow,
		3,
	)
	statter := &monitor.Statter{
		StatsD:    statsDClient,
		Histogram: queryProbeHistogram,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go RunAdminProbe(ctx, logger.Session("admin-probe"), &wg, adminProbe, statter, parserOpts.AdminProbe.Interval, parserOpts.AdminProbe.Timeout)
	go RunQueryProbe(ctx, logger.Session("query-probe"), &wg, queryProbe, statter, parserOpts.QueryProbe.Interval, parserOpts.AdminProbe.Timeout)

	wg.Wait()
	os.Exit(0)
}
