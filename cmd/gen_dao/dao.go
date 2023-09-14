package main

import (
	"gorm.io/gen"
	"server/config"
	"server/dal/model"
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

	g.ApplyBasic(model.User{}, model.Captcha{}, model.Account{})

	g.Execute()
}
