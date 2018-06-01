package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	docPath := c.Param("path")

	fullPath := os.Getenv("GOPATH") + "/src" + docPath + "/"
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}

	type Doc struct {
		Path     string
		Filename string
	}
	var docs []*Doc
	for _, file := range files {
		if !file.IsDir() && path.Ext(file.Name()) == ".md" {
			docs = append(docs, &Doc{
				Path:     docPath[1:] + "/" + file.Name(),
				Filename: file.Name(),
			})
		}
	}
	if len(docs) == 0 {
		c.JSON(http.StatusOK, fmt.Sprintf("no .md files in dir: %s", docPath[1:]))
		return
	}

	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"docs": docs,
	})
}
