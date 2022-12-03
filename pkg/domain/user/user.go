package user

type Role string
type User struct {
	Id    int
	Roles []Role
}
