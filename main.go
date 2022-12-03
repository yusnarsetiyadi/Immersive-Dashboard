package main

import (
	"api-alta-dashboard/config"
	"api-alta-dashboard/factory"
	"api-alta-dashboard/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := posgresql.InitDB(cfg)

	e := echo.New()

	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
