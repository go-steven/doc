package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func DocHandler(c *gin.Context) {
	doc := c.Param("doc")
	logger.Printf("input doc: %s\n", doc)

	if check(doc == "" || path.Ext(doc) != EXT_MD, fmt.Sprintf("[%s]not a %s file.", EXT_MD, doc), c) {
		return
	}

	f, err := ioutil.ReadFile(basePath() + doc)
	if checkErr(err, c) {
		return
	}

	c.HTML(http.StatusOK, "doc.tpl", gin.H{
		"Title":   doc,
		"Content": template.HTML(blackfriday.MarkdownCommon([]byte(f))),
	})
}
