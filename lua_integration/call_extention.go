package lua_integration

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/yuin/gopher-lua"
)

func CallExtention(name string, db *sql.DB) (*string, error) {
	L := lua.NewState()
	Loader(L, db)
	if val, isOk := getExtentions()[name]; isOk {
		if err := L.DoFile("./extentions/" + val); err != nil {
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
