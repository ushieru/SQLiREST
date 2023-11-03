package handlers

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/ushieru/sqlirest/utils"
)

func SetupPatchHandler(app *fiber.App, db *sql.DB) {
	app.Patch("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		queries := c.Queries()
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
		ub := sqlbuilder.NewUpdateBuilder()
		setValues := make([]string, len(values))
		for i := range setValues {
			setValues[i] = ub.Assign(cols[i], values[i])
		}
		ub.Update(table)
		ub.Set(setValues...)
		if params := utils.BuildWhereParams(&ub.Cond, queries); len(params) > 0 {
			ub.Where(params...)
		}
		sql, args := ub.Build()
		_, err := db.Exec(sql, args...)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

}
