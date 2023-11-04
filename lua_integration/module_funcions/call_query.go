package module_funcions

import (
	"database/sql"
	"encoding/json"

	"github.com/ushieru/sqlirest/utils"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

func CallQuery(db *sql.DB) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		stm := L.ToString(1)
		args := L.ToTable(2)
		mapArgs := gluamapper.ToGoValue(args, gluamapper.Option{NameFunc: func(s string) string { return s }, TagName: "gluamapper"}).(map[any]any)
		sliceArgs := make([]any, 0)
		for _, value := range mapArgs {
			sliceArgs = append(sliceArgs, value)
		}
		rows, _ := db.Query(stm, sliceArgs...)
		mapRows, _ := utils.SQLRowsToMaps(rows)
		mJson, _ := json.Marshal(mapRows)
		L.Push(lua.LString(mJson))
		return 1
	}
}
