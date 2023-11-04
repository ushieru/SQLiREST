package handlers

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/ushieru/sqlirest/lua_integration"
	"github.com/ushieru/sqlirest/utils"
)

func SetupGetHandler(app *fiber.App, db *sql.DB) {
	app.Get("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		res, err := lua_integration.CallExtention(table, db)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if res != nil {
			return c.SendString(*res)
		}
		selectParams := c.Query("select", "*")
		queries := c.Queries()
		selectParamsSlice := strings.Split(selectParams, ",")
		sb := sqlbuilder.NewSelectBuilder()
		sb.Select(selectParamsSlice...)
		sb.From(table)
		if params := utils.BuildWhereParams(&sb.Cond, queries); len(params) > 0 {
			sb.Where(params...)
		}
		sql, params := sb.Build()
		rows, err := db.Query(sql, params...)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		maps, err := utils.SQLRowsToMaps(rows)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.JSON(maps)
	})
}
