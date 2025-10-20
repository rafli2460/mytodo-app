package main

import (
	"encoding/json"
	"os"
	"time"

	swagger "github.com/gofiber/swagger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/rafli024/mytodo-app/internal/config"
	"github.com/rafli024/mytodo-app/internal/constant"
	"github.com/rafli024/mytodo-app/internal/contract"
	"github.com/rafli024/mytodo-app/internal/datasources"
	"github.com/rafli024/mytodo-app/internal/handler"
	"github.com/rafli024/mytodo-app/internal/middlewares"
	"github.com/rafli024/mytodo-app/internal/router"
	"github.com/rafli024/mytodo-app/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	_ "github.com/rafli024/mytodo-app/docs"
)

func NewApp() *contract.App {
	os.Setenv("TZ", "Asia/Jakarta")

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: handler.HttpError,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	customLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	fiberApp.Use(fiberzerolog.New(
		fiberzerolog.Config{
			Logger: &customLogger,
		}),
	)

	crs := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Access-Control-Allow-Origin, Accept, content-type, X-Requested-With, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Screen",
	})

	fiberApp.Use(crs)

	cfg := config.Init()

	if cfg[constant.ServerEnv] == constant.EnvDevelopment {
		fiberApp.Use(pprof.New())
	}

	app := &contract.App{
		Fiber:  fiberApp,
		Logger: &customLogger,
		Config: cfg,
	}

	app.Ds = datasources.Init(cfg)
	app.Services = service.Init(app)
	middlewares.Init(app)
	handler.Init(app)
	router.InitRoutes(app)

	return app
}


// @title Todo List API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app := NewApp()

	app.Fiber.Get("/swagger/*", swagger.HandlerDefault)

	if err := app.Fiber.Listen(":" + app.Config[constant.ServerPort]); err != nil {
		app.Logger.Fatal().Err(err).Msg("Fiber App Error")
	}
}
