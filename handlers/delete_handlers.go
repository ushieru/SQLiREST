package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-sqlbuilder"
	"github.com/ushieru/sqlirest/utils"
)

func SetupDeleteHandler(app *fiber.App, db *sql.DB) {
	app.Delete("/:table", func(c *fiber.Ctx) error {
		table := c.Params("table")
		queries := c.Queries()
		dlb := sqlbuilder.NewDeleteBuilder()
		dlb.DeleteFrom(table)
		if params := utils.BuildWhereParams(&dlb.Cond, queries); len(params) > 0 {
			dlb.Where(params...)
		}
		sql, params := dlb.Build()
		_, err := db.Exec(sql, params...)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})
}
