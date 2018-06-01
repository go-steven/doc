package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const EXT_MD = ".md"

type Doc struct {
	Path     string
	FileName string
}

func IndexHandler(c *gin.Context) {
	docPath := c.Param("path")
	if len(docPath) == 0 || (docPath[len(docPath)-1:] != "/" && docPath[len(docPath)-1:] != "\\") {
		docPath += "/"
	}
	logger.Printf("doc path: %s\n", docPath)

	files, err := listFiles(basePath()+docPath, EXT_MD)
	if checkErr(err, c) {
		return
	}

	var docs []*Doc
	for _, v := range files {
		docs = append(docs, &Doc{
			Path:     docPath[1:] + v,
			FileName: v,
		})
	}
	if check(len(docs) == 0, fmt.Sprintf("no %s files in dir: %s", EXT_MD, docPath[1:]), c) {
		return
	}

	c.HTML(http.StatusOK, "index.tpl", gin.H{
		"docs": docs,
	})
}
