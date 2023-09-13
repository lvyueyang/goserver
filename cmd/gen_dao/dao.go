package main

import (
	"gorm.io/gen"
	"server/config"
	"server/dal/model"
	"server/db"
)

func main() {
	config.New()
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	database := db.Connect()

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(database)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(model.User{}, model.Captcha{}, model.Account{})

	// Generate the code
	g.Execute()
}
