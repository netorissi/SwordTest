package model

type Role string

const (
	RoleTech    Role = "tech"
	RoleManager Role = "manager"
)

type User struct {
	ID   int  `json:"id" db:"id"`
	Role Role `json:"role" db:"role"`
}
