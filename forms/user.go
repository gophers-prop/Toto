package forms

//CreateUserCommand : struct to create user
type CreateUserCommand struct {
	Name   string  `json:"name" binding:"required"`
	ID   string  `json:"id" binding:"required"`
	Age int `json:"age" binding:"required"`
}

//UpdateUserCommand : struct to update user
type UpdateUserCommand struct {
	Name   string  `json:"name" binding:"required"`
	ID   string  `json:"id" binding:"required"`
	Age int `json:"age" binding:"required"`
}

// ResponseMsg represents response message
type ResponseMsg struct {
	Message string `json:"message" binding:"required"`
}