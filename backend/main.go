package main

import (
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.NoRoute(func(c *gin.Context) {
    dir, file := path.Split(c.Request.RequestURI)
    ext := filepath.Ext(file)
    if file == "" || ext == "" {
      c.File("./static/index.html")
      } else {
        c.File("./static/" + path.Join(dir, file))
      }
    })
  router.GET("/api/v1/comics", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "Welcome Gin Server")
  })
  router.Run("localhost:8080")
}
