package rpc

import (
	"context"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/messages"
	"code.cloudfoundry.org/perm/models"
	"code.cloudfoundry.org/perm/protos"
)

type RoleServiceServer struct {
	logger      lager.Logger
	assignments map[models.Actor][]string

	roleService           models.RoleService
	roleAssignmentService models.RoleAssignmentService
}

func NewRoleServiceServer(logger lager.Logger, roleService models.RoleService, roleAssignmentService models.RoleAssignmentService) *RoleServiceServer {
	return &RoleServiceServer{
		logger:                logger,
		assignments:           make(map[models.Actor][]string),
		roleService:           roleService,
		roleAssignmentService: roleAssignmentService,
	}
}

func (s *RoleServiceServer) CreateRole(ctx context.Context, req *protos.CreateRoleRequest) (*protos.CreateRoleResponse, error) {
	logger := s.logger.Session("create-role")
	name := req.GetName()
	logData := lager.Data{"role.name": name}
	logger.Debug(messages.Starting, logData)

	role, err := s.roleService.CreateRole(ctx, logger, name)

	if err != nil {
		return nil, togRPCErrorNew(err)
		//err := togRPCError(codes.AlreadyExists, errors.New(messages.ErrRoleAlreadyExists))
		//logger.Error(messages.ErrRoleAlreadyExists, err, logData)
		//return nil, err
	}

	logger.Debug(messages.Success, logData)
	return &protos.CreateRoleResponse{
		Role: role.ToProto(),
	}, nil
}

func (s *RoleServiceServer) GetRole(ctx context.Context, req *protos.GetRoleRequest) (*protos.GetRoleResponse, error) {
	logger := s.logger.Session("get-role")
	name := req.GetName()
	logData := lager.Data{"role.name": name}

	role, err := s.roleService.FindRole(ctx, logger, models.RoleQuery{Name: name})
	if err != nil {
		return nil, togRPCErrorNew(err)
	}

	logger.Debug(messages.Success, logData)
	return &protos.GetRoleResponse{
		Role: role.ToProto(),
	}, nil
}

func (s *RoleServiceServer) DeleteRole(ctx context.Context, req *protos.DeleteRoleRequest) (*protos.DeleteRoleResponse, error) {
	name := req.GetName()
	logger := s.logger.Session("delete-role").WithData(lager.Data{
		"role.name": name,
	})

	err := s.roleService.DeleteRole(ctx, logger, models.RoleQuery{Name: name})
	if err != nil {
		return nil, togRPCErrorNew(err)
	}

	return &protos.DeleteRoleResponse{}, nil
}

func (s *RoleServiceServer) AssignRole(ctx context.Context, req *protos.AssignRoleRequest) (*protos.AssignRoleResponse, error) {
	roleName := req.GetRoleName()
	pActor := req.GetActor()

	domainID := pActor.GetID()
	issuer := pActor.GetIssuer()
	logger := s.logger.Session("assign-role").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
		"role.name":    roleName,
	})

	err := s.roleAssignmentService.AssignRole(ctx, logger, roleName, domainID, issuer)
	if err != nil {
		return nil, togRPCErrorNew(err)
	}

	return &protos.AssignRoleResponse{}, nil
}

func (s *RoleServiceServer) UnassignRole(ctx context.Context, req *protos.UnassignRoleRequest) (*protos.UnassignRoleResponse, error) {
	roleName := req.GetRoleName()
	pActor := req.GetActor()

	domainID := pActor.GetID()
	issuer := pActor.GetIssuer()
	actor := models.Actor{
		DomainID: domainID,
		Issuer:   issuer,
	}
	logger := s.logger.Session("unassign-role").WithData(lager.Data{
		"actor.id":     actor.DomainID,
		"actor.issuer": actor.Issuer,
		"role.name":    roleName,
	})

	err := s.roleAssignmentService.UnassignRole(ctx, logger, roleName, domainID, issuer)
	if err != nil {
		return nil, togRPCErrorNew(err)
	}

	return &protos.UnassignRoleResponse{}, nil
}

func (s *RoleServiceServer) HasRole(ctx context.Context, req *protos.HasRoleRequest) (*protos.HasRoleResponse, error) {
	roleName := req.GetRoleName()
	pActor := req.GetActor()
	domainID := pActor.GetID()
	issuer := pActor.GetIssuer()

	logger := s.logger.Session("has-role").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
		"role.name":    roleName,
	})

	query := models.RoleAssignmentQuery{
		ActorQuery: models.ActorQuery{
			DomainID: domainID,
			Issuer:   issuer,
		},
		RoleQuery: models.RoleQuery{
			Name: roleName,
		},
	}

	found, err := s.roleAssignmentService.HasRole(ctx, logger, query)
	if err != nil {
		if err == models.ErrRoleNotFound || err == models.ErrActorNotFound {
			return &protos.HasRoleResponse{HasRole: false}, nil
		} else {
			return nil, togRPCErrorNew(err)
		}
	}

	return &protos.HasRoleResponse{HasRole: found}, nil
}

func (s *RoleServiceServer) ListActorRoles(ctx context.Context, req *protos.ListActorRolesRequest) (*protos.ListActorRolesResponse, error) {
	pActor := req.GetActor()
	domainID := pActor.GetID()
	issuer := pActor.GetIssuer()
	logger := s.logger.Session("list-actor-roles").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
	})

	roles, err := s.roleAssignmentService.ListActorRoles(ctx, logger, models.ActorQuery{DomainID: domainID, Issuer: issuer})
	if err != nil {
		if err == models.ErrActorNotFound {
			roles = []*models.Role{}
		} else {
			return nil, togRPCErrorNew(err)
		}
	}

	var pRoles []*protos.Role

	for _, r := range roles {
		pRoles = append(pRoles, r.ToProto())
	}

	return &protos.ListActorRolesResponse{
		Roles: pRoles,
	}, nil
}
