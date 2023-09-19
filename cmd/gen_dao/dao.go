package main

import (
	"gorm.io/gen"
	"server/config"
	"server/db"
	"server/types"
)

func main() {
	config.New()
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/dao",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	database := db.Connect()

	g.UseDB(database)

	g.ApplyBasic(db.Models...)
	g.ApplyInterface(func(CommonFindQuery) {}, db.Models...)

	g.Execute()
}

type CommonFindQuery interface {
	// FindList // 查询列表
	//
	// SELECT *
	// FROM @@table
	// {{ if order.OrderKey != "" }}
	//  {{ if order.OrderType == "desc"}}
	//    ORDER BY @@order.OrderKey DESC
	//  {{ else }}
	//    ORDER BY @@order.OrderKey
	//  {{ end }}
	// {{ end }}
	// LIMIT @limit
	// OFFSET @offset
	FindList(order types.Order, offset, limit int) (list []*gen.T, err error)

	// FindByID // 根据 ID 查询
	//
	// SELECT * FROM @@table WHERE id=@id
	FindByID(id uint) (*gen.T, error)
}
