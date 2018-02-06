package models

import "code.cloudfoundry.org/perm-go"

type PermissionName string

type PermissionDefinition struct {
	Name PermissionName
}

type PermissionResourcePattern string

type Permission struct {
	Name            PermissionName
	ResourcePattern PermissionResourcePattern
}

func (p *Permission) ToProto() *protos.Permission {
	return &protos.Permission{
		Name:            string(p.Name),
		ResourcePattern: string(p.ResourcePattern),
	}
}
