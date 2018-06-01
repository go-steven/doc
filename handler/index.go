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
	logger.Printf("docPath: %s\n", docPath)

	fullPath := os.Getenv("GOPATH") + "/src" + docPath + "/"
	logger.Printf("fullPath: %s\n", fullPath)
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
		//fmt.Printf("file.Name(): %s\n", file.Name())
		if !file.IsDir() && path.Ext(file.Name()) == ".md" {
			logger.Printf("file.Name(): %s\n", file.Name())

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
