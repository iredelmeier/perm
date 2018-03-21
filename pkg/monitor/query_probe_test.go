package monitor_test

import (
	permgofakes "code.cloudfoundry.org/perm-go/perm-gofakes"
	. "code.cloudfoundry.org/perm/pkg/monitor"

	"context"

	"errors"

	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/perm-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("QueryProbe", func() {
	var (
		p *QueryProbe

		fakeRoleServiceClient        *permgofakes.FakeRoleServiceClient
		fakePermissionsServiceClient *permgofakes.FakePermissionServiceClient
		fakeLogger                   *lagertest.TestLogger
		fakeContext                  context.Context

		uniqueSuffix string

		someError error
	)

	BeforeEach(func() {
		fakeRoleServiceClient = new(permgofakes.FakeRoleServiceClient)
		fakePermissionsServiceClient = new(permgofakes.FakePermissionServiceClient)

		fakeLogger = lagertest.NewTestLogger("query-probe")
		fakeContext = context.Background()

		uniqueSuffix = "foobar"

		p = &QueryProbe{
			RoleServiceClient:       fakeRoleServiceClient,
			PermissionServiceClient: fakePermissionsServiceClient,
		}

		someError = errors.New("some-error")
	})

	Describe("Setup", func() {
		It("creates a role with a permission and assigns it to a test user", func() {
			err := p.Setup(fakeContext, fakeLogger, uniqueSuffix)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeRoleServiceClient.CreateRoleCallCount()).To(Equal(1))
			_, createRoleRequest, _ := fakeRoleServiceClient.CreateRoleArgsForCall(0)
			Expect(createRoleRequest.GetName()).To(Equal("system.query-probe.foobar"))
			permissions := createRoleRequest.GetPermissions()
			Expect(permissions).To(HaveLen(1))
			Expect(permissions[0].GetName()).To(Equal("system.query-probe.assigned-permission.name"))
			Expect(permissions[0].GetResourcePattern()).To(Equal("system.query-probe.assigned-permission.resource-id.foobar"))

			Expect(fakeRoleServiceClient.AssignRoleCallCount()).To(Equal(1))
			_, assignRoleRequest, _ := fakeRoleServiceClient.AssignRoleArgsForCall(0)
			Expect(assignRoleRequest.GetRoleName()).To(Equal("system.query-probe.foobar"))
			Expect(assignRoleRequest.GetActor().GetIssuer()).To(Equal("system"))
			Expect(assignRoleRequest.GetActor().GetID()).To(Equal("query-probe"))
		})

		Context("when creating the role", func() {

			Context("when the role already exists", func() {
				BeforeEach(func() {
					fakeRoleServiceClient.CreateRoleReturns(nil, status.Error(codes.AlreadyExists, "role-already-exists"))
				})
				It("swallows the error", func() {
					err := p.Setup(fakeContext, fakeLogger, uniqueSuffix)
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when any other grpc error occurs", func() {
				BeforeEach(func() {
					fakeRoleServiceClient.CreateRoleReturns(nil, status.Error(codes.Unavailable, "server-not-available"))
				})

				It("errors", func() {
					err := p.Setup(fakeContext, fakeLogger, uniqueSuffix)
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("when assigning the role", func() {

			Context("when the role has already been assigned", func() {
				BeforeEach(func() {
					fakeRoleServiceClient.AssignRoleReturns(nil, status.Error(codes.AlreadyExists, "role-assignment-already-exists"))
				})

				It("swallows the error", func() {
					err := p.Setup(fakeContext, fakeLogger, uniqueSuffix)
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when any other grpc error occurs", func() {
				BeforeEach(func() {
					fakeRoleServiceClient.AssignRoleReturns(nil, status.Error(codes.Unavailable, "server-not-available"))
				})

				It("errors", func() {
					err := p.Setup(fakeContext, fakeLogger, uniqueSuffix)
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Cleanup", func() {
		It("deletes the role", func() {
			err := p.Cleanup(fakeContext, fakeLogger, uniqueSuffix)
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeRoleServiceClient.DeleteRoleCallCount()).To(Equal(1))
			_, deleteRoleRequest, _ := fakeRoleServiceClient.DeleteRoleArgsForCall(0)
			Expect(deleteRoleRequest.GetName()).To(Equal("system.query-probe.foobar"))
		})

		Context("when the role doesn't exist", func() {
			BeforeEach(func() {
				fakeRoleServiceClient.DeleteRoleReturns(nil, status.Error(codes.NotFound, "role-not-found"))
			})

			It("swallows the error", func() {
				err := p.Cleanup(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when any other grpc error occurs", func() {
			BeforeEach(func() {
				fakeRoleServiceClient.DeleteRoleReturns(nil, status.Error(codes.Unavailable, "server-not-available"))
			})

			It("errors", func() {
				err := p.Cleanup(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Run", func() {
		BeforeEach(func() {
			fakePermissionsServiceClient.HasPermissionReturnsOnCall(0, &protos.HasPermissionResponse{HasPermission: true}, nil)
			fakePermissionsServiceClient.HasPermissionReturnsOnCall(1, &protos.HasPermissionResponse{HasPermission: false}, nil)
		})

		It("asks if the actor has a permission it should have, and a permission it shouldn't", func() {
			correct, durations, err := p.Run(fakeContext, fakeLogger, uniqueSuffix)
			Expect(err).NotTo(HaveOccurred())
			Expect(correct).To(BeTrue())
			Expect(durations).To(HaveLen(2))

			Expect(fakePermissionsServiceClient.HasPermissionCallCount()).To(Equal(2))

			_, hasPositivePermissionRequest, _ := fakePermissionsServiceClient.HasPermissionArgsForCall(0)
			Expect(hasPositivePermissionRequest.GetActor().GetIssuer()).To(Equal("system"))
			Expect(hasPositivePermissionRequest.GetActor().GetID()).To(Equal("query-probe"))
			Expect(hasPositivePermissionRequest.GetPermissionName()).To(Equal("system.query-probe.assigned-permission.name"))
			Expect(hasPositivePermissionRequest.GetResourceId()).To(Equal("system.query-probe.assigned-permission.resource-id.foobar"))

			_, hasNegativePermissionRequest, _ := fakePermissionsServiceClient.HasPermissionArgsForCall(1)
			Expect(hasNegativePermissionRequest.GetActor().GetIssuer()).To(Equal("system"))
			Expect(hasNegativePermissionRequest.GetActor().GetID()).To(Equal("query-probe"))
			Expect(hasNegativePermissionRequest.GetPermissionName()).To(Equal("system.query-probe.unassigned-permission.name"))
			Expect(hasNegativePermissionRequest.GetResourceId()).To(Equal("system.query-probe.unassigned-permission.resource-id.foobar"))
		})

		Context("when checking for the permission it should have errors", func() {
			BeforeEach(func() {
				fakePermissionsServiceClient.HasPermissionReturnsOnCall(0, nil, someError)
			})

			It("errors and does not ask for the next permission", func() {
				_, durations, err := p.Run(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).To(MatchError(someError))
				Expect(durations).To(HaveLen(1))

				Expect(fakePermissionsServiceClient.HasPermissionCallCount()).To(Equal(1))
			})
		})

		Context("when checking for the permission it should have returns that the actor doesn't have permission", func() {
			BeforeEach(func() {
				fakePermissionsServiceClient.HasPermissionReturnsOnCall(0, &protos.HasPermissionResponse{HasPermission: false}, nil)
			})

			It("returns that it's incorrect and does not ask for the next permission", func() {
				correct, durations, err := p.Run(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).NotTo(HaveOccurred())
				Expect(correct).To(BeFalse())
				Expect(durations).To(HaveLen(1))

				Expect(fakePermissionsServiceClient.HasPermissionCallCount()).To(Equal(1))
			})
		})

		Context("when checking for the permission it should not have errors", func() {
			BeforeEach(func() {
				fakePermissionsServiceClient.HasPermissionReturnsOnCall(1, nil, someError)
			})

			It("errors", func() {
				_, durations, err := p.Run(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).To(MatchError(someError))
				Expect(durations).To(HaveLen(2))

				Expect(fakePermissionsServiceClient.HasPermissionCallCount()).To(Equal(2))
			})
		})

		Context("when checking for the permission it should not have returns that the actor does have permission", func() {
			BeforeEach(func() {
				fakePermissionsServiceClient.HasPermissionReturnsOnCall(1, &protos.HasPermissionResponse{HasPermission: true}, nil)
			})

			It("returns that it's incorrect", func() {
				correct, durations, err := p.Run(fakeContext, fakeLogger, uniqueSuffix)
				Expect(err).NotTo(HaveOccurred())
				Expect(correct).To(BeFalse())
				Expect(durations).To(HaveLen(2))

				Expect(fakePermissionsServiceClient.HasPermissionCallCount()).To(Equal(2))
			})
		})
	})
})