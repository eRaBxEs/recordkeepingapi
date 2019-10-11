package controller

import (
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

//Here we define the kind of request and the manner of the response

// H ... for mapping strings to any interface
type H map[string]interface{}

// DefineRoutes used to define all routes used in the application
func DefineRoutes(e *echo.Echo, db *pg.DB, log *zap.Logger, prefix string) error {

	handlers := []Handler{
		&Record{},
	}

	env := &Environment{
		DB:  db,
		Rtr: e,
		Log: log,
	}

	for _, h := range handlers {
		if err := h.Init(env, prefix); err != nil {
			return err
		}
	}

	return nil
}
