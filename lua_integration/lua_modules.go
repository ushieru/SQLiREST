package lua_integration

import (
	"database/sql"
	"encoding/json"

	"github.com/ushieru/sqlirest/lua_integration/module_funcions"
	"github.com/yuin/gopher-lua"
	gopher_json "layeh.com/gopher-json"
)

func Loader(L *lua.LState, db *sql.DB, args map[string]any) {
	L.SetGlobal("getArgs", L.NewFunction(func(l *lua.LState) int {
		arsMap, err := json.Marshal(args)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		L.Push(lua.LString(arsMap))
		return 1
	}))
	L.PreloadModule("database", loadDBFunctions(db))
	gopher_json.Preload(L)
}

func loadDBFunctions(db *sql.DB) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		functions := map[string]lua.LGFunction{
			"callRawQuery": module_funcions.CallRawQuery(db),
			"callQuery":    module_funcions.CallQuery(db),
			"execRawQuery": module_funcions.ExecRawQuery(db),
		}
		mod := L.SetFuncs(L.NewTable(), functions)
		L.Push(mod)
		return 1
	}
}
