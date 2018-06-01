package handler

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func logErr(err string) {
	logger.Printf("Err: %s\n", err)
}

func check(flag bool, err string, c *gin.Context) (ret bool) {
	ret = flag
	if ret {
		logErr(err)
		c.JSON(http.StatusOK, fmt.Sprintf("Err: %s", err))
	}
	return
}

func checkErr(err error, c *gin.Context) (ret bool) {
	ret = err != nil
	if ret {
		logErr(err.Error())
		c.JSON(http.StatusOK, fmt.Sprintf("Err: %s", err.Error()))
	}
	return
}

func listFiles(dir string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(p string, f os.FileInfo, err error) error {
		if f == nil || err != nil {
			logger.Printf("Err: %s\n", err.Error())
			return err
		}

		var ignoredDir = []string{
			".idea",
			".git",
		}
		for _, v := range ignoredDir {
			if strings.Contains(p, "\\"+v+"\\") || strings.Contains(p, "/"+v+"/") {
				return nil
			}
		}

		if !f.IsDir() && path.Ext(f.Name()) == ext && len(p) >= len(dir) {
			files = append(files, p[len(dir):])
		}

		return nil
	})

	return files, err
}

func basePath() string {
	return os.Getenv("GOPATH") + "/src"
}
