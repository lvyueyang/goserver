package model

type {{.Name}} struct {
	BaseModel
	Cover     string `json:"cover"` // 封面
	Title     string `json:"title" gorm:"unique"`
	Desc      string `json:"desc"`
	Content   string `json:"content" gorm:"type:longtext"`
	PushDate  string `json:"wx_open_id"`
	Recommend string `json:"recommend"` // 推荐等级
	AuthorID  uint   `json:"author_id"` // 作者
}

func (*{{.name}}) TableName() string {
	return "news"
}
