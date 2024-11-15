package routes

import (
    "employees/controller"
    "github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App, employeeController *controller.Employee) {
    app.Post("/employees", employeeController.Create)
    app.Get("/employees", employeeController.GetAll)
    
    // Adding the /health route
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("OK")
    })
}
