package handlers

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/ushieru/sqlirest/lua_integration"
)

func SetupPostHandler(app *fiber.App, db *sql.DB) {
	app.Post("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		body := make(map[string]any)
		json.Unmarshal(c.Body(), &body)
		res, _ := lua_integration.CallLuaScript("post", table, body, db)
		if res != nil {
			var v any
			if err := json.Unmarshal([]byte(*res), &v); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "...")
			}
			return c.JSON(v)
		}
		cols := make([]string, len(body))
		values := make([]any, len(body))
		i := 0
		for k, v := range body {
			cols[i] = k
			values[i] = v
			i++
		}
		sql, args := sqlbuilder.InsertInto(table).Cols(cols...).Values(values...).Build()
		_, err := db.Exec(sql, args...)
		if err != nil {
			log.Default().Printf(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		}
		return c.SendStatus(fiber.StatusCreated)
	})
}
