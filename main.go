package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LatLongPair struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type LatLongReading struct {
	LatLong LatLongPair `json:"LatLong"`
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	frontendHostname := os.Getenv("FRONTENDHOSTNAME")
	frontendPort := os.Getenv("FRONTENDPORT")

	// Defaults for dev environment
	if frontendHostname == "" {
		frontendHostname = "localhost"
	}
	if frontendPort == "" {
		frontendPort = "3000"
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://"+frontendHostname+":"+frontendPort},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// CondenserReading

	e.GET("/condenser/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, NewCondenserReading(id))
	}) // Endpoint to get CondenserReading

	// VolumeReading

	e.GET("/volume/:id", func(c echo.Context) error {
		liters := (rand.Float64() * 5) + 10 // generate random liter number
		return c.JSON(http.StatusOK, NewVolumeReading(liters))
	}) // Endpoint to get VolumeReading

	// HumidityReading

	e.GET("/humidity/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewHumidityReading())
	}) // Endpoint to get HumidityReading

	// LastEmptiedReading

	e.GET("/lastemptied/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewLastEmptiedReading())
	}) // Endpoint to get LastEmptiedReading

	// UptimeReading

	e.GET("/uptime/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewUptimeReading())
	}) // Endpoint to get LastEmptiedReading

	// LatLonReading

	e.GET("/latlon/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewLatLonReading())
	}) // Endpoint to get LastEmptiedReading

	// Set up server port

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "7777"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
