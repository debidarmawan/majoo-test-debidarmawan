package http

import (
	"fmt"

	"majoo-test-debidarmawan/config"
	_ "majoo-test-debidarmawan/docs"
	"majoo-test-debidarmawan/handlers"
	"majoo-test-debidarmawan/repositories"
	"majoo-test-debidarmawan/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func Routes(f *fiber.App, dbConn *config.DbConnection) {
	f.Get("/swagger/*", fiberSwagger.WrapHandler)

	routerGroup := f.Group("/v1")
	routerGroup.Use(recover.New())
	routerGroup.Use(logger.New())

	routerGroupNoAuth := f.Group("/v1")
	routerGroupNoAuth.Use(recover.New())
	routerGroupNoAuth.Use(logger.New())

	routerGroup.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: fmt.Sprintf("%s,%s,%s,%s,%s", fiber.MethodGet, fiber.MethodPost, fiber.MethodPut, fiber.MethodPatch, fiber.MethodDelete),
	}))

	userRepo := repositories.NewUserRepo(dbConn)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)
	userHandler.Routes(routerGroupNoAuth)

	merchantRepo := repositories.NewMerchantRepo(dbConn)
	merchantUseCase := usecases.NewMerchantUseCase(merchantRepo)
	merchantHandler := handlers.NewMerchantHandler(merchantUseCase)
	merchantHandler.Routes(routerGroup)
}
