package controller

import (
	"recordkeeping/lib/util"

	"go.uber.org/zap"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
)

// Environment struct to pass around "global" config
type Environment struct {
	DB  *pg.DB
	Rtr *echo.Echo
	Val *util.CustomValidator
	Log *zap.Logger
}

// Handler interface for methods that handle requests
type Handler interface {
	Init(*Environment, string) error
}
