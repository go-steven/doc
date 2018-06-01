package doc

import (
	"os"
	"path"
	"runtime"
)

func curr_path(skip int) string {
	_, file, _, _ := runtime.Caller(skip)
	return path.Dir(file)
}

func is_allowed_host(allowedHosts []string) bool {
	host, _ := os.Hostname()
	for _, v := range allowedHosts {
		if v == host {
			return true
		}
	}
	return false
}
