package module_funcions

import (
	"database/sql"
	"encoding/json"

	"github.com/ushieru/sqlirest/utils"
	"github.com/yuin/gopher-lua"
)

func CallRawQuery(db *sql.DB) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		stm := L.ToString(1)
		r, err := db.Query(stm)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		mapRows, err := utils.SQLRowsToMaps(r)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		mJson, err := json.Marshal(mapRows)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LString(mJson))
		return 1
	}
}
