package lua_integration

import (
	"database/sql"
	"errors"

	"github.com/yuin/gopher-lua"
)

func CallLuaScript(prefix, name string, args map[string]any, db *sql.DB) (*string, error) {
	L := lua.NewState()
	Loader(L, db, args)
	script, err := getLuaScript(prefix, name)
	if err != nil {
		return nil, err
	}
	if err := L.DoString(script); err != nil {
		return nil, err
	}
	lv := L.Get(-1)
	lstr, ok := lv.(lua.LString)
	if !ok {
		return nil, errors.New("something wrong")
	}
	s := string(lstr)
	return &s, nil
}
