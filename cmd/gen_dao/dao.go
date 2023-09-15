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
	g.ApplyInterface(func(FindQuery) {}, db.Models...)

	g.Execute()
}

type FindQuery interface {
	// FindList
	// SELECT COUNT(*) OVER() AS total_count, *
	// FROM @@table
	// {{ if order.OrderKey != "" }}
	// ORDER BY order.OrderKey order.OrderType
	// {{ end }}
	// LIMIT page.PageSize
	// OFFSET (page.Current - 1) * page.PageSize
	FindList(order types.Order, page types.Pagination) ([]gen.T, error)
}
