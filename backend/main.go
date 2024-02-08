package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/morsok/comic-handler/ent"
	"github.com/spf13/viper"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
  handler := log.New(os.Stderr)
  logger := slog.New(handler)
  handler.SetTimeFormat(time.StampMilli)
  // Load Config
  logger.Debug("Loading Configuration")
  viper.SetDefault("LogLevel", "info")
  viper.SetDefault("directories.config", "/config")
  viper.SetDefault("directories.comics", "/comics")
  viper.SetDefault("directories.watch", "/watch")
  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath("/config")
  if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
      log.Info("No config file found, creating default one")
      viper.SafeWriteConfig()
    } else {
      log.Fatal(fmt.Errorf("fatal error config file: %w", err))
    }
  }
  log.Info("Configuration loaded")
  // Config file found and successfully parsed
  level, err := log.ParseLevel(viper.GetString("LogLevel"))
  if err != nil {
    logger.Error(fmt.Sprintf("failed to parse log level, reverting to INFO: %v", err))
    level = log.InfoLevel
  }
  handler.SetLevel(level)
  // Database
  // Create an ent.Client with in-memory SQLite database.
  logger.Debug("Opening internal sqlite DB")
  dbClient, err := ent.Open(dialect.SQLite, "file:comichandler.db?mode=memory&cache=shared&_fk=1") // TODO: use file DB
  if err != nil {
    logger.Error(fmt.Sprintf("Failed opening connection to internal sqlite DB: %v", err))
    os.Exit(1)
  }
  defer dbClient.Close()
  logger.Debug("DB opened with success")
  ctx := context.Background()
  // Run the automatic migration tool to create all schema resources.
  logger.Debug("Initializing the DB and/or running migrations")
  if err := dbClient.Schema.Create(ctx, schema.WithAtlas(true), migrate.WithGlobalUniqueID(true)); err != nil {
    logger.Error(fmt.Sprintf("Database initialization: Failed creating schema resources: %v", err))
    os.Exit(1)
  }
  logger.Debug("DB init and/or migrations success")

  // TMP
  serieTest, err := dbClient.Serie.Create().Save(ctx)
  if err != nil {
    logger.Error(fmt.Sprintf("Failed creating a test serie: %v", err))
  }
  logger.Debug("New task added to db", "task", serieTest)
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
  logger.Debug("Starting web server")
  router.Run("0.0.0.0:9999")
}
