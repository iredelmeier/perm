package perm_test

import (
	"crypto/tls"
	"crypto/x509"
	"net"

	"code.cloudfoundry.org/perm/pkg/api"
	"code.cloudfoundry.org/perm/pkg/api/db"
	"code.cloudfoundry.org/perm/pkg/sqlx"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MySQL server", func() {
	var (
		conn     *sqlx.DB
		listener net.Listener

		subject *api.Server
	)

	BeforeEach(func() {
		var err error

		conn, err = testMySQLDB.Connect()
		Expect(err).NotTo(HaveOccurred())

		// Port 0 should find a random open port
		listener, err = net.Listen("tcp", "localhost:0")
		Expect(err).NotTo(HaveOccurred())

		cert, err := tls.X509KeyPair([]byte(testCert), []byte(testCertKey))
		Expect(err).NotTo(HaveOccurred())

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		store := db.NewDataService(conn)
		subject = api.NewServer(store, api.WithTLSConfig(tlsConfig))

		go func() {
			err = subject.Serve(listener)
			Expect(err).NotTo(HaveOccurred())
		}()
	})

	AfterEach(func() {
		subject.Stop()

		err := testMySQLDB.Truncate(
			"DELETE FROM role",
			"DELETE FROM action",
		)
		Expect(err).NotTo(HaveOccurred())

		err = conn.Close()
		Expect(err).NotTo(HaveOccurred())
	})

	testAPI(func() serverConfig {
		rootCAPool := x509.NewCertPool()

		ok := rootCAPool.AppendCertsFromPEM([]byte(testCA))
		Expect(ok).To(BeTrue())

		return serverConfig{
			addr: listener.Addr().String(),
			tlsConfig: &tls.Config{
				RootCAs: rootCAPool,
			},
		}
	})
})
