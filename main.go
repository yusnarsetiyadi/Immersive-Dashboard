package main

import (
	"api-alta-dashboard/config"
	"api-alta-dashboard/factory"
	"api-alta-dashboard/middlewares"
	"api-alta-dashboard/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := posgresql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	// middleware
	middlewares.LogMiddlewares(e)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
