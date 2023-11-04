package lua_integration

import (
	"database/sql"

	"github.com/ushieru/sqlirest/lua_integration/module_funcions"
	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState, db *sql.DB) {
	L.PreloadModule("database", loadDBFunctions(db))
}

func loadDBFunctions(db *sql.DB) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		functions := map[string]lua.LGFunction{
			"callRawQuery": module_funcions.CallRawQuery(db),
			"execQuery":    module_funcions.ExecQuery(db),
		}
		mod := L.SetFuncs(L.NewTable(), functions)
		L.Push(mod)
		return 1
	}
}
