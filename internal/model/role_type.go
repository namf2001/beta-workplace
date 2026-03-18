package model

// OrgRole represents the role of a member within an organization
type OrgRole string

const (
	OrgRoleAdmin    OrgRole = "admin"
	OrgRoleSubAdmin OrgRole = "sub_admin"
	OrgRoleMember   OrgRole = "member"
)

// String converts to string value
func (r OrgRole) String() string {
	return string(r)
}

// IsValid checks if OrgRole is valid
func (r OrgRole) IsValid() bool {
	return r == OrgRoleAdmin || r == OrgRoleSubAdmin || r == OrgRoleMember
}

// WorkplaceRole represents the role of a member within a workplace
type WorkplaceRole string

const (
	WorkplaceRoleAdmin  WorkplaceRole = "admin"
	WorkplaceRoleMember WorkplaceRole = "member"
)

// String converts to string value
func (r WorkplaceRole) String() string {
	return string(r)
}

// IsValid checks if WorkplaceRole is valid
func (r WorkplaceRole) IsValid() bool {
	return r == WorkplaceRoleAdmin || r == WorkplaceRoleMember
}

// ProjectRole represents the role of a member within a project
type ProjectRole string

const (
	ProjectRoleOwner  ProjectRole = "owner"
	ProjectRoleMember ProjectRole = "member"
)

// String converts to string value
func (r ProjectRole) String() string {
	return string(r)
}

// IsValid checks if ProjectRole is valid
func (r ProjectRole) IsValid() bool {
	return r == ProjectRoleOwner || r == ProjectRoleMember
}
