package lua_integration

import (
	"database/sql"
	"encoding/json"

	"github.com/Shopify/go-lua"
	"github.com/ushieru/sqlirest/utils"
)

func RegisterCallRawQuery(l *lua.State, db *sql.DB) {
	l.Register("callRawQuery", func(l *lua.State) int {
		stm, isOk := l.ToString(1)
		if !isOk || stm == "" {
			l.PushNil()
			return 1
		}
		r, err := db.Query(stm)
		if err != nil {
			l.PushNil()
			return 1
		}
		mapRows, err := utils.SQLRowsToMaps(r)
		if err != nil {
			l.PushNil()
			return 1
		}
		mJson, err := json.Marshal(mapRows)
		if err != nil {
			l.PushNil()
			return 1
		}
		l.PushString(string(mJson))
		return 1
	})
}
