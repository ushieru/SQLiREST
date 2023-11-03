package main

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/ushieru/sqlirest/handlers"
	"github.com/ushieru/sqlirest/lua_integration"
	"github.com/ushieru/sqlirest/utils"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.SetDefault("DATABASE_URL", ":memory:")
	viper.SetDefault("PORT", "8080")
	database := viper.GetString("DATABASE_URL")
	sqlFile := viper.GetString("SQL_FILE")
	port := viper.GetString("PORT")

	app := fiber.New()
	db, err := sql.Open("sqlite", database)
	if err != nil {
		log.Fatal(err)
	}

	lua_integration.GetExtentions()
	utils.LoadSQL(sqlFile, db)

	handlers.SetupGetHandler(app, db)
	handlers.SetupPostHandler(app, db)
	handlers.SetupPatchHandler(app, db)
	handlers.SetupDeleteHandler(app, db)

	log.Fatal(app.Listen(":" + port))
}
