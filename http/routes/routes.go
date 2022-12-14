package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fuaa/core"
)

func Register(router *echo.Echo, services *core.Services, urls map[string]string) {
	NewGetLogin(services, urls).Register(router)
	NewPostOauthToken(services, urls).Register(router)
}
