package contract

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type App struct {
	Logger   *zerolog.Logger
	Fiber    *fiber.App
	Config   map[string]string
	Ds       *Datasources
	Services *Service
}
