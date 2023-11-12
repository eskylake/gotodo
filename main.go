package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eskylake/go-todo/config"
	"github.com/eskylake/go-todo/database"
	"github.com/eskylake/go-todo/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	conf config.Config
	err  error
)

func dd(arg ...any) {
	log.Fatalln(arg...)
	panic("Die")
}

func init() {
	conf, err = config.LoadConfig(".")
	if err != nil {
		dd("Failed to load environment variables! \n", err.Error())
	}

	database.ConnectDB(&conf)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(routers.Cors())

	routers.SetupRoutes(micro)

	port, err := strconv.Atoi(conf.APIPort)
	if err != nil {
		dd(err.Error())
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
