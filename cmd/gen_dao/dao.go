package main

import (
	"gorm.io/gen"
	"server/config"
	"server/db"
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

	g.Execute()
}
