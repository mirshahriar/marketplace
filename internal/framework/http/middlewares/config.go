package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
)

type localConfig struct {
	Skipper middleware.Skipper
}
