package model

import "time"

type News struct {
	BaseModel
	Cover     string    `json:"cover"` // 封面
	Title     string    `json:"title" gorm:"unique"`
	Desc      string    `json:"desc"`
	Content   string    `json:"content" gorm:"type:longtext"`
	PushDate  time.Time `json:"push_date"` // 发布日期
	Recommend uint      `json:"recommend"` // 推荐等级
	AuthorID  uint      `json:"author_id"` // 作者
}

func (*News) TableName() string {
	return "news"
}
