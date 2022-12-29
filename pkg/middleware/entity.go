package middleware

import (
	"net/http"

	"github.com/francoganga/go_reno2/pkg/context"
	"github.com/francoganga/go_reno2/pkg/domain/user"

	"github.com/labstack/echo/v4"
)

// LoadUser loads the user based on the ID provided as a path parameter
func LoadUser(userRepo user.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Param("user")
			// if err != nil {
			// 	return echo.NewHTTPError(http.StatusNotFound)
			// }

			u, err := userRepo.Get(c.Request().Context(), userID)

			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound)
			}

			c.Set(context.UserKey, u)
			return next(c)

			// switch err.(type) {
			// case nil:
			// 	c.Set(context.UserKey, u)
			// 	return next(c)
			// case *ent.NotFoundError:
			// 	return echo.NewHTTPError(http.StatusNotFound)
			// default:
			// 	return echo.NewHTTPError(
			// 		http.StatusInternalServerError,
			// 		fmt.Sprintf("error querying user: %v", err),
			// 	)
			// }
		}
	}
}
