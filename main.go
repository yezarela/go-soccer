package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yezarela/go-soccer/module/player"
	"github.com/yezarela/go-soccer/module/team"
	"github.com/yezarela/go-soccer/pkg/conn"
)

func init() {
	// Load configuration
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB cluster
	c, err := conn.NewMongoDBConnection(ctx, os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect(ctx)

	db := c.Database(os.Getenv("MONGODB_DBNAME"))

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Repositories
	teamRepo := team.NewRepository(db)
	playerRepo := player.NewRepository(db)

	// Route & handlers
	team.NewHandler(e, teamRepo)
	player.NewHandler(e, playerRepo)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
