package memory

import (
	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fuaa/core"
)

func NewServices() *core.Services {
	tokens := NewTokenService()
	users := NewUserService(tokens)

	return &core.Services{
		Tokens: tokens,
		Users:  users,
		Logger: func(e echo.Context) {},
	}
}
