package http

import (
	"log"
	"majoo-test-debidarmawan/config"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const DefaultPort = 8000

func ServeHTTP(dbConn *config.DbConnection) {
	f := fiber.New()

	Routes(f, dbConn)
	var port string
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		port = portEnv
	} else {
		port = strconv.Itoa(DefaultPort)
	}
	log.Fatal((f.Listen(":" + port)))
}
