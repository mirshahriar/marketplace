// Package middlewares contains all the middlewares used by the API
// nolint: wrapcheck
package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mirshahriar/marketplace/internal/ports"
)

// AuthorizeUser is the middleware for authorizing a user
func AuthorizeUser(adapter ports.APIPort) echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		// TODO: Add caching here
		user, found, cErr := adapter.GetUserByToken(key)
		if cErr != nil {
			return false, cErr
		}

		if !found {
			return false, nil
		}

		// Setting the user in the context
		c.Set("user", user)

		return true, nil
	})
}
