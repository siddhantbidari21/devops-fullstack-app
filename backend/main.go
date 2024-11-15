package main

import (
    "employees/controller"
    "employees/repository"
    "employees/routes"
    "employees/service"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error loading .env file")
    }

    app := fiber.New()
    app.Use(cors.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: os.Getenv("ALLOWED_ORIGINS"),
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
    db := initializeDatabaseConnection()
    repository.RunMigrations(db)
    employeeRepository := repository.NewEmployeeRepository(db)
    employeeService := service.NewEmployeeService(employeeRepository)
    employeeController := controller.NewEmployeeController(employeeService)
    routes.RegisterRoute(app, employeeController)

    go func() {
        err := app.Listen("0.0.0.0:8080")
        if err != nil {
            log.Fatalln(fmt.Sprintf("error starting the server %s", err.Error()))
        }
    }()

    // Block forever to keep the container running
    select {}
}

func initializeDatabaseConnection() *gorm.DB {
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  createDsn(),
        PreferSimpleProtocol: true,
    }), &gorm.Config{})
    if err != nil {
        log.Fatalln(fmt.Sprintf("error connecting with database %s", err.Error()))
    }
    return db
}

func createDsn() string {
    dsnFormat := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    return fmt.Sprintf(dsnFormat, dbHost, dbUser, dbPassword, dbName, dbPort)
}

