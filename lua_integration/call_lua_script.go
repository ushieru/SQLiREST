package lua_integration

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/yuin/gopher-lua"
)

func CallLuaScript(prefix, name string, args map[string]any, db *sql.DB) (*string, error) {
	L := lua.NewState()
	Loader(L, db, args)
	if val, isOk := getLuaScripts(prefix)[name]; isOk {
		if err := L.DoFile("./scripts/" + val); err != nil {
			e := strings.Split(err.Error(), ":")
			if len(e) == 4 {
				return &e[3], nil
			}
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
	return nil, nil
}
