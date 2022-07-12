package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kebhr/mhz19"
	"log"
)

type App struct {
	Sensor *mhz19.MHZ19
	Fiber  *fiber.App
}

func (a *App) Initialize() {
	log.Println("Initializing app...")

	a.Sensor = &mhz19.MHZ19{}
	a.connectToSensor()

	a.Fiber = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ServerHeader:          "mhz19-pi (v0.1)",
	})
	a.setupRoutes()

	log.Println("Web server will be ready on port 3500")
	a.listen()
}

func (a *App) connectToSensor() {
	if err := a.Sensor.Connect(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) setupRoutes() {
	a.Fiber.Use(cors.New())
	a.Fiber.Get("/api/v1/co2", func(c *fiber.Ctx) error {
		val, err := a.Sensor.ReadCO2()
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		}
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		return c.SendString(fmt.Sprintf("{\"co2\": %d}", val))
	})
}

func (a *App) listen() {
	err := a.Fiber.Listen(":3500")
	if err != nil {
		log.Fatal(err)
	}
}
