package perm_test

import (
	"crypto/tls"
	"crypto/x509"
	"net"

	"code.cloudfoundry.org/perm/pkg/api"
	"code.cloudfoundry.org/perm/pkg/api/db"
	"code.cloudfoundry.org/perm/pkg/sqlx"
	"code.cloudfoundry.org/perm/pkg/sqlx/sqlxtest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MySQL server", func() {
	var (
		testDB   *sqlxtest.TestMySQLDB
		conn     *sqlx.DB
		listener net.Listener

		subject *api.Server
	)

	BeforeSuite(func() {
		var err error

		testDB = sqlxtest.NewTestMySQLDB()

		err = testDB.Create(db.Migrations...)
		Expect(err).NotTo(HaveOccurred())

		conn, err = testDB.Connect()
		Expect(err).NotTo(HaveOccurred())

		// Port 0 should find a random open port
		listener, err = net.Listen("tcp", "localhost:0")
		Expect(err).NotTo(HaveOccurred())

		cert, err := tls.X509KeyPair([]byte(testCert), []byte(testCertKey))
		Expect(err).NotTo(HaveOccurred())

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
		subject = api.NewServer(conn, api.WithTLSConfig(tlsConfig))

		go func() {
			err = subject.Serve(listener)
			Expect(err).NotTo(HaveOccurred())
		}()
	})

	AfterEach(func() {
		err := testDB.Truncate(
			"DELETE FROM role",
			"DELETE FROM actor",
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		subject.Stop()

		err := conn.Close()
		Expect(err).NotTo(HaveOccurred())

		err = testDB.Drop()
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
