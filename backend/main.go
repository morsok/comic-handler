package main

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/morsok/comic-handler/ent"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
  log.SetTimeFormat(time.StampMilli)
  // Load Config
  log.Info("Loading Configuration")
  // TODO Load configuration
  log.SetLevel(log.DebugLevel) // TODO Set from conf
  // Database
  // Create an ent.Client with in-memory SQLite database.
  log.Debug("Opening internal sqlite DB")
  dbClient, err := ent.Open(dialect.SQLite, "file:comichandler.db?mode=memory&cache=shared&_fk=1") // TODO: use file DB
  if err != nil {
      log.Fatal(fmt.Sprintf("Failed opening connection to internal sqlite DB: %v", err))
  }
  defer dbClient.Close()
  log.Debug("DB openned with success")
  ctx := context.Background()
  // Run the automatic migration tool to create all schema resources.
  log.Debug("Initializing the DB and/or running migrations")
  if err := dbClient.Schema.Create(ctx); err != nil {
      log.Fatal(fmt.Sprintf("Database initialization: Failed creating schema resources: %v", err))
  }
  log.Debug("DB init and/or migrations success")

  // TMP
  serieTest, err := dbClient.Serie.Create().Save(ctx)
  if err != nil {
      log.Fatal(fmt.Sprintf("Failed creating a test serie: %v", err))
  }
  log.Debug("New task added to db", "task", serieTest)
  // END TMP

  // HTTP Server
  router := gin.Default()
  router.SetTrustedProxies(nil)
  router.NoRoute(func(c *gin.Context) {
    dir, file := path.Split(c.Request.RequestURI)
    ext := filepath.Ext(file)
    if file == "" || ext == "" {
      c.File("./static/index.html")
      } else {
        c.File("./static/" + path.Join(dir, file))
      }
    })
  v1 := router.Group("/api/v1") // TODO Add authorization
	{
		v1.GET("/comics", func(c *gin.Context) {
      c.String(http.StatusOK, "Welcome Gin Server v1")
    })
	}
  router.Run("localhost:9999")
}
