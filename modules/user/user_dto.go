package user

type CreateUserBodyDto struct {
	Name string `json:"name" binding:"required" label:"姓名"` // 姓名
	Sex  string `json:"sex" binding:"required" label:"性别"`  // 性别
}
