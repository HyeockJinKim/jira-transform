package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func APIKeyMiddleware(expectedAPIKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("Authorization")
			if apiKey != "Bearer "+expectedAPIKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
			}
			return next(c)
		}
	}
}
