package handler

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func DocHandler(c *gin.Context) {
	doc := c.Param("doc")
	logger.Printf("doc: %s\n", doc)

	if doc != "" || path.Ext(doc) == EXT_MD {
		f, err := ioutil.ReadFile(os.Getenv(ENV_GOPATH) + "/" + GOPATH_SRC + doc)
		if err != nil {
			logger.Printf("Err: %s\n", err.Error())
			c.JSON(http.StatusOK, err.Error())
			return
		}
		c.HTML(http.StatusOK, TPL_DOC, gin.H{
			"Title":   doc,
			"Content": template.HTML(blackfriday.MarkdownCommon([]byte(f))),
		})
	} else {
		c.JSON(http.StatusOK, fmt.Sprintf("[%s]not a %s file.", EXT_MD, doc))
	}
}
