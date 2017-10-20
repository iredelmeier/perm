package db

import "code.cloudfoundry.org/perm/db/migrations"

var Migrations = []Migration{
	{
		Name: "create_roles_table",
		Up:   migrations.CreateRolesTableUp,
		Down: migrations.CreateRolesTableDown,
	},
	{
		Name: "create_actors_table",
		Up:   migrations.CreateActorsTableUp,
		Down: migrations.CreateActorsTableDown,
	},
	{
		Name: "create_role_assignments_table",
		Up:   migrations.CreateRoleAssignmentsTableUp,
		Down: migrations.CreateRoleAssignmentsTableDown,
	},
}
