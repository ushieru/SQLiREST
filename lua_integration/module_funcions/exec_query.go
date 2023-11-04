package module_funcions

import (
	"database/sql"

	"github.com/yuin/gopher-lua"
)

func ExecQuery(db *sql.DB) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		stm := L.ToString(1)
		_, err := db.Exec(stm)
		if err != nil {
			L.Push(lua.LFalse)
			return 1
		}
		L.Push(lua.LTrue)
		return 1
	}
}
