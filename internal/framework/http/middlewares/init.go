package middlewares

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_custom} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out} "${user_agent}"` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.00",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if len(reqBody) != 0 {
			fmt.Println("--- Request ---")
			fmt.Println(string(reqBody))
		}
		if len(resBody) != 0 {
			fmt.Println("--- Response ---")
			fmt.Println(string(resBody))
		}
	}))

	e.Use(middleware.Secure())
	e.Use(middleware.Recover())

	skipList := []string{
		"/swagger",
		"/health",
	}

	e.Use(authorizeUser(localConfig{
		Skipper: func(c echo.Context) bool {
			reqURI := c.Request().RequestURI

			for _, skip := range skipList {
				if strings.HasPrefix(reqURI, skip) {
					return true
				}
			}
			return false
		},
	}))
}
