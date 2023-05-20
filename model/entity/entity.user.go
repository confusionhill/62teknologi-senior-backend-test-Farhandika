package entity

type User struct {
	Id       uint64
	Email    string
	Password string
	FullName string
	Access   uint32
}
