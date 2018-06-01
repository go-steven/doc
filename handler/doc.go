package handler

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"fmt"
)

func DocHandler(c *gin.Context) {
	doc := c.Param("doc")
	if doc != "" || path.Ext(doc) == ".md" {
		f, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src" + doc)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.HTML(http.StatusOK, "doc.tpl", gin.H{
			"Title":   doc,
			"Content": template.HTML(blackfriday.MarkdownCommon([]byte(f))),
		})
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("[%s]not a .md file.", doc))
	}
}
