package servers

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MirrorHandler interface {
	Mirror(c echo.Context) error
}

type ServerArgs struct {
	Addr    string
	Handler MirrorHandler
}

type Server struct {
	addr string
	e    *echo.Echo
}

func NewServer(args ServerArgs) *Server {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/mirror/:destination", args.Handler.Mirror)
	return &Server{
		addr: args.Addr,
		e:    e,
	}
}

func (s *Server) Start(ctx context.Context) error {
	return s.e.Start(s.addr)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
