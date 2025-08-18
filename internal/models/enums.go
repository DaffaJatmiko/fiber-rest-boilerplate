package models

// User Roles
const (
	RoleUser  = "USER"
	RoleAdmin = "ADMIN"
)

var ValidRoles = []string{RoleUser, RoleAdmin}

// helper function
func IsValidRole(role string) bool {
	for _, validRole := range ValidRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
