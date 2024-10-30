package entity

type UserCreateEntity struct {
	Username       string
	Email          string
	HashedPassword string
}

type UserDetailEntity struct {
	ID       int
	Username string
	Email    string
}
