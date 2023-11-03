package lua_integration

import (
	"database/sql"
	"errors"

	"github.com/Shopify/go-lua"
)

func CallExtention(name string, db *sql.DB) (string, error) {
	if val, isOk := Extentions[name]; isOk {
		l := lua.NewState()
		lua.OpenLibraries(l)
		RegisterCallRawQuery(l, db)
		if err := lua.DoFile(l, "./extentions/"+val); err != nil {
			return "", err
		}
		result, _ := l.ToString(-1)
		return result, nil
	}
	return "", errors.New("not found")
}
