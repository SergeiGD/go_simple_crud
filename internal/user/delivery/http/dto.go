package http

type CreateUserDOT struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserByIDDOT struct {
	ID int `uri:"id" binding:"required"`
}
