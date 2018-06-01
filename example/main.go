package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/doc"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()

	r := gin.New()
	r.Use(gin.Recovery())

	doc.DocRouter(r, []string{}, "", 1)
	r.Run() // listen and serve on 0.0.0.0:8080
}
