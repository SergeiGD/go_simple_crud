package entity

type UserCreateRawEntity struct {
	Username    string
	Email       string
	RawPassword string
}

type UserCreateEntity struct {
	Username       string
	Email          string
	HashedPassword string
	PasswordSalt   string
}

type UserDetailEntity struct {
	ID       int
	Username string
	Email    string
}
