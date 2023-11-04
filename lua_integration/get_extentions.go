package lua_integration

import (
	"os"
	"strings"
)

func getExtentions() map[string]string {
	files, err := os.ReadDir("./extentions")
	extentions := make(map[string]string)
	if err != nil {
		return nil
	}
	for _, file := range files {
		fullName := file.Name()
		name := strings.Split(fullName, ".")[0]
		extentions[name] = fullName
	}
	return extentions
}
