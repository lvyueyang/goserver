package user

type CreateUserBodyDto struct {
	Name string `json:"name" binding:"required"`
}
