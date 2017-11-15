package monitor

import (
	"context"
	"errors"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/messages"
	"code.cloudfoundry.org/perm/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	QueryProbeRoleName = "system.query-probe"
)

var QueryProbeActor = &protos.Actor{
	ID:     "query-probe",
	Issuer: "system",
}

var QueryProbeAssignedPermission = &protos.Permission{
	Name:            "system.query-probe.assigned-permission.name",
	ResourcePattern: "system.query-probe.assigned-permission.resource-id",
}

var QueryProbeUnassignedPermission = &protos.Permission{
	Name:            "system.query-probe.unassigned-permission.name",
	ResourcePattern: "system.query-probe.unassigned-permission.resource-id",
}

var (
	ErrQueryProbeAssignedPermissionIncorrect   = errors.New("query probe found actor did not have permission, when it should have")
	ErrQueryProbeUnassignedPermissionIncorrect = errors.New("query probe found actor had permission, when it should not have")
)

//go:generate counterfeiter code.cloudfoundry.org/perm/protos.PermissionServiceClient

type QueryProbe struct {
	RoleServiceClient       protos.RoleServiceClient
	PermissionServiceClient protos.PermissionServiceClient
}

func (p *QueryProbe) Setup(ctx context.Context, logger lager.Logger) error {
	logger.Debug(messages.Starting)
	defer logger.Debug(messages.Finished)

	createRoleRequest := &protos.CreateRoleRequest{
		Name: QueryProbeRoleName,
		Permissions: []*protos.Permission{
			QueryProbeAssignedPermission,
		},
	}
	_, err := p.RoleServiceClient.CreateRole(ctx, createRoleRequest)
	s, ok := status.FromError(err)

	// Not a GRPC error
	if err != nil && !ok {
		logger.Error(messages.FailedToCreateRole, err, lager.Data{
			"roleName":    createRoleRequest.GetName(),
			"permissions": createRoleRequest.GetPermissions(),
		})
		return err
	}

	// GRPC error
	if err != nil && ok {
		switch s.Code() {
		case codes.AlreadyExists:

		default:
			logger.Error(messages.FailedToCreateRole, err, lager.Data{
				"roleName":    createRoleRequest.GetName(),
				"permissions": createRoleRequest.GetPermissions(),
			})
			return err
		}
	}

	assignRoleRequest := &protos.AssignRoleRequest{
		Actor:    QueryProbeActor,
		RoleName: QueryProbeRoleName,
	}

	_, err = p.RoleServiceClient.AssignRole(ctx, assignRoleRequest)
	s, ok = status.FromError(err)

	// Not a GRPC error
	if err != nil && !ok {
		logger.Error(messages.FailedToAssignRole, err, lager.Data{
			"roleName":     assignRoleRequest.GetRoleName(),
			"actor.id":     assignRoleRequest.GetActor().GetID(),
			"actor.issuer": assignRoleRequest.GetActor().GetIssuer(),
		})
		return err
	}

	// GRPC error
	if err != nil && ok {
		switch s.Code() {
		case codes.AlreadyExists:

		default:
			logger.Error(messages.FailedToAssignRole, err, lager.Data{
				"roleName":     assignRoleRequest.GetRoleName(),
				"actor.id":     assignRoleRequest.GetActor().GetID(),
				"actor.issuer": assignRoleRequest.GetActor().GetIssuer(),
			})
			return err
		}
	}

	return nil
}

func (p *QueryProbe) Cleanup(ctx context.Context, logger lager.Logger) error {
	logger.Debug(messages.Starting)
	defer logger.Debug(messages.Finished)

	deleteRoleRequest := &protos.DeleteRoleRequest{
		Name: QueryProbeRoleName,
	}
	_, err := p.RoleServiceClient.DeleteRole(ctx, deleteRoleRequest)
	s, ok := status.FromError(err)

	// Not a GRPC error
	if err != nil && !ok {
		logger.Error(messages.FailedToDeleteRole, err, lager.Data{
			"roleName": deleteRoleRequest.GetName(),
		})
		return err
	}

	// GRPC error
	if err != nil && ok {
		switch s.Code() {
		case codes.NotFound:

		default:
			logger.Error(messages.FailedToDeleteRole, err, lager.Data{
				"roleName": deleteRoleRequest.GetName(),
			})
			return err
		}
	}

	return nil
}

func (p *QueryProbe) Run(ctx context.Context, logger lager.Logger) (bool, error) {
	logger.Debug(messages.Starting)
	defer logger.Debug(messages.Finished)

	hasAssignedPermissionRequest := &protos.HasPermissionRequest{
		Actor:          QueryProbeActor,
		PermissionName: QueryProbeAssignedPermission.Name,
		ResourceId:     QueryProbeAssignedPermission.ResourcePattern,
	}

	hasAssignedPermissionResponse, err := p.PermissionServiceClient.HasPermission(ctx, hasAssignedPermissionRequest)
	if err != nil {
		logger.Error(messages.FailedToFindPermissions, err, lager.Data{
			"actor.id":                    QueryProbeActor.GetID(),
			"actor.issuer":                QueryProbeActor.GetIssuer(),
			"permission.name":             QueryProbeAssignedPermission.GetName(),
			"permission.resource_pattern": QueryProbeAssignedPermission.GetResourcePattern(),
		})
		return false, err
	}

	if hasAssignedPermissionResponse.GetHasPermission() == false {
		return false, nil
	}

	hadUnassignedPermissionRequest := &protos.HasPermissionRequest{
		Actor:          QueryProbeActor,
		PermissionName: QueryProbeUnassignedPermission.Name,
		ResourceId:     QueryProbeUnassignedPermission.ResourcePattern,
	}

	hasUnassignedPermissionResponse, err := p.PermissionServiceClient.HasPermission(ctx, hadUnassignedPermissionRequest)
	if err != nil {
		logger.Error(messages.FailedToFindPermissions, err, lager.Data{
			"actor.id":                    QueryProbeActor.GetID(),
			"actor.issuer":                QueryProbeActor.GetIssuer(),
			"permission.name":             QueryProbeUnassignedPermission.GetName(),
			"permission.resource_pattern": QueryProbeUnassignedPermission.GetResourcePattern(),
		})
		return false, err
	}

	if hasUnassignedPermissionResponse.GetHasPermission() == true {
		return false, nil
	}

	return true, nil
}