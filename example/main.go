package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-steven/doc"
)

func main() {
	r := gin.New()
	doc.Router(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
