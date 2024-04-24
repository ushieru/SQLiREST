package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/ushieru/sqlirest/lua_integration"
	"github.com/ushieru/sqlirest/utils"
)

func SetupGetHandler(app *fiber.App, db *sql.DB) {
	app.Get("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		queries := c.Queries()
		args := make(map[string]any)
		res, _ := lua_integration.CallLuaScript("get", table, args, db)
		if res != nil {
			var v interface{}
			if err := json.Unmarshal([]byte(*res), &v); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "...")
			}
			return c.JSON(v)
		}
		selectParams := c.Query("select", "*")
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
			log.Default().Printf(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		}
		maps, err := utils.SQLRowsToMaps(rows)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.JSON(maps)
	})
}
