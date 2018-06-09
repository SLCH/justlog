package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server api server
type Server struct {
	port     string
	logPath  string
	channels []string
}

// NewServer create Server
func NewServer() Server {
	return Server{
		logPath:  "/var/twitch_logs",
		channels: []string{},
		port:     ":8025",
	}
}

// AddChannel to in-memory store to serve joined channels
func (s *Server) AddChannel(channel string) {
	s.channels = append(s.channels, channel)
}

// Init api server
func (s *Server) Init() {

	e := echo.New()
	e.HideBanner = true

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/channel/:channel/user/:username", s.getCurrentUserLogs)
	e.GET("/channel", s.getAllChannels)
	e.GET("/channel/:channel", s.getCurrentChannelLogs)
	e.GET("/channel/:channel/:year/:month/:day", s.getDatedChannelLogs)
	e.GET("/channel/:channel/user/:username/:year/:month", s.getDatedUserLogs)
	e.GET("/channel/:channel/user/:username/random", s.getRandomQuote)

	fmt.Println("Starting API on port " + s.port)
	e.Logger.Fatal(e.Start(s.port))
}
