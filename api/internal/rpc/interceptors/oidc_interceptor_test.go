package interceptors_test

import (
	"context"

	"code.cloudfoundry.org/perm/logx/logxfakes"
	"code.cloudfoundry.org/perm/oidcx/oidcxfakes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	. "code.cloudfoundry.org/perm/api/internal/rpc/interceptors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth Server Interceptor", func() {
	var (
		interceptor        grpc.UnaryServerInterceptor
		fakeProvider       *oidcxfakes.FakeProvider
		sampleHandler      func(context.Context, interface{}) (interface{}, error)
		fakeSecurityLogger *logxfakes.FakeSecurityLogger
	)

	BeforeEach(func() {
		fakeProvider = new(oidcxfakes.FakeProvider)
		fakeSecurityLogger = new(logxfakes.FakeSecurityLogger)
		interceptor = OIDCInterceptor(fakeProvider, fakeSecurityLogger)
		sampleHandler = func(c context.Context, r interface{}) (interface{}, error) { return nil, nil }
	})

	It("errors out when context contains no metadata", func() {
		_, err := interceptor(context.Background(), nil, nil, sampleHandler)

		Expect(err).To(MatchError("perm: unauthenticated"))
	})

	It("errors out when context does not contain a token field", func() {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"key": "value"}))
		_, err := interceptor(ctx, nil, nil, sampleHandler)

		Expect(err).To(MatchError("perm: unauthenticated"))
	})

	It("errors out when token isn't valid", func() {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"token": "hello"}))
		_, err := interceptor(ctx, nil, nil, sampleHandler)

		Expect(err).To(MatchError("perm: unauthenticated"))
	})
})
