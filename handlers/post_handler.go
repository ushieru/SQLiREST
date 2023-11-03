package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
)

func SetupPostHandler(app *fiber.App, db *sql.DB) {
	app.Post("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		body := make(map[string]interface{})
		json.Unmarshal(c.Body(), &body)
		cols := make([]string, len(body))
		values := make([]interface{}, len(body))
		i := 0
		for k, v := range body {
			cols[i] = k
			values[i] = v
			i++
		}
		sql, args := sqlbuilder.InsertInto(table).Cols(cols...).Values(values...).Build()
		db.Exec(sql, args...)
		return c.SendStatus(fiber.StatusCreated)
	})
}
