package doc

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/doc/handler"
)

func DocRouter(r *gin.Engine, allowedHosts []string, templateDir string, skip int) {
	if len(allowedHosts) > 0 && !is_allowed_host(allowedHosts) {
		return
	}

	if templateDir == "" {
		templateDir = curr_path(skip) + "/templates"
	}
	r.Delims("{{", "}}")
	r.LoadHTMLGlob(templateDir + "/*.tpl")

	handler.SetLogger(logger)
	g := r.Group("/doc")
	{
		g.GET("/list/*path", handler.IndexHandler)
		g.GET("/md/*doc", handler.DocHandler)
	}
}
