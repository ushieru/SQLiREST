package lua_integration

import (
	"os"
	"strings"
)

var Extentions = make(map[string]string)

func GetExtentions() {
	files, err := os.ReadDir("./extentions")
	if err != nil {
		return
	}
	for _, file := range files {
		fullName := file.Name()
		name := strings.Split(fullName, ".")[0]
		Extentions[name] = fullName
	}
}
