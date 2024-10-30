package http

type CreateUserDOT struct {
	Username string
	Email    string
	Password string
}

type GetUserByIDDOT struct {
	ID int `uri:"id" binding:"required"`
}
