package main

import (
  "path"
  "path/filepath"

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
  router.Run("localhost:8080")
}
