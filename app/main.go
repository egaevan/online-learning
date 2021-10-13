package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/egaevan/online-learning/config"
	"github.com/egaevan/online-learning/delivery/rest"
	"github.com/egaevan/online-learning/repository"
	"github.com/egaevan/online-learning/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// Init config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Init echo framework
	e := echo.New()

	// Init DB
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)

	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		fmt.Println("Error connect to db")
		log.Fatal(err)
	}

	defer db.Close()

	// Init repository
	courseRepo := repository.NewCourseRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Init usecase
	courseUsecae := usecase.NewCourse(courseRepo)
	userUsecae := usecase.NewUser(userRepo)

	// Init handler
	rest.NewHandler(e, courseUsecae, userUsecae)

	e.Logger.Fatal(e.Start(":8080"))
}
