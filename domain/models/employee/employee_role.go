package employee

import (
	"database/sql"
	"strings"
)

type EmployeeRole struct {
	RoleId    sql.NullInt64
	Role      sql.NullString
	Hierarchy sql.NullInt32
}

type EmployeeRoles []EmployeeRole

func (roles EmployeeRoles) FindRoleByName(roleName string) EmployeeRole {
	for i := range roles {
		roleName = strings.ToLower(roleName)
		currentRole := strings.ToLower(roles[i].Role.String)
		if currentRole == roleName {
			return roles[i]
		}
	}

	return EmployeeRole{}
}

func (roles EmployeeRoles) FindHierarchByRoleId(roleId int32) int32 {
	for i := range roles {
		if roles[i].RoleId.Int64 == int64(roleId) {
			return roles[i].Hierarchy.Int32
		}
	}

	return 0
}

func (role *EmployeeRole) GetRoleId() int64 {
	return role.RoleId.Int64
}
