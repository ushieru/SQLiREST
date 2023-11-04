package lua_integration

import (
	"os"
	"strings"
)

func getLuaScripts(method string) map[string]string {
	files, err := os.ReadDir("./scripts")
	extentions := make(map[string]string)
	if err != nil {
		return nil
	}
	for _, file := range files {
		fullName := file.Name()
		nameSplit := strings.Split(fullName, ".")
		if len(nameSplit) == 3 {
			methodScript := nameSplit[0]
			if methodScript != method {
				continue
			}
			name := nameSplit[1]
			extentions[name] = fullName
		}
		name := nameSplit[0]
		extentions[name] = fullName
	}
	return extentions
}
