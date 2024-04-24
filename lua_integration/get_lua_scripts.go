package lua_integration

import (
	"os"
	"path/filepath"
)

func getLuaScript(method string, name string) (string, error) {
	extentionPath := filepath.Join("srv", method+"."+name+".lua")
	file, err := os.ReadFile(extentionPath)
	if err != nil {
		path := filepath.Join("srv", name+".lua")
		file, err = os.ReadFile(path)
		if err == nil {
			return string(file), nil
		}
		return "", err
	}
	return string(file), nil
}
