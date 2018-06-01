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

func IndexHandler(c *gin.Context) {
	docPath := c.Param("path")
	if len(docPath) == 0 || (docPath[len(docPath)-1:] != "/" && docPath[len(docPath)-1:] != "\\") {
		docPath += "/"
	}

	fullPath := os.Getenv(ENV_GOPATH) + "/" + GOPATH_SRC + docPath
	logger.Printf("fullPath: %s\n", fullPath)

	files, err := get_file_list(fullPath, EXT_MD)
	if err != nil {
		logger.Printf("Err: %s\n", err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}

	type Doc struct {
		Path     string
		Filename string
	}
	var docs []*Doc
	for _, v := range files {
		docs = append(docs, &Doc{
			Path:     docPath[1:] + v,
			Filename: v,
		})
	}
	if len(docs) == 0 {
		c.JSON(http.StatusOK, fmt.Sprintf("no %s files in dir: %s", EXT_MD, docPath[1:]))
		return
	}

	c.HTML(http.StatusOK, TPL_INDEX, gin.H{
		"docs": docs,
	})
}

func get_file_list(dir string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(p string, f os.FileInfo, err error) error {
		if f == nil || err != nil {
			logger.Printf("Err: %s\n", err.Error())
			return err
		}
		if p == dir {
			return nil
		}

		if strings.Contains(p, "\\.idea\\") || strings.Contains(p, "/.idea/") {
			return nil
		}
		if strings.Contains(p, "\\.git\\") || strings.Contains(p, "/.git/") {
			return nil
		}
		if strings.Contains(p, "\\vendor\\") || strings.Contains(p, "/vendor/") {
			return nil
		}

		filename := f.Name()
		if strings.HasPrefix(filename, ".") || filename == "vendor" {
			return nil
		}

		if !f.IsDir() && path.Ext(filename) == ext && len(p) >= len(dir) {
			files = append(files, p[len(dir):])
		}

		return nil
	})
	return files, err
}
