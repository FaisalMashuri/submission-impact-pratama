package main

import (
	"fmt"
	"log"
	"time"

	"github.com/FaisalMashuri/submission-golang/app"
	"github.com/FaisalMashuri/submission-golang/app/controller"
	"github.com/FaisalMashuri/submission-golang/app/repository"
	"github.com/FaisalMashuri/submission-golang/app/usecase"
	"github.com/FaisalMashuri/submission-golang/config"
	"github.com/FaisalMashuri/submission-golang/driver"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal(err)
		log.Fatal("Error loading .env file")
	}
	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
		fmt.Println("error")
	}
	db := driver.InitDB(cfg.DBUSER, cfg.DBPASSWORD, cfg.HOSTDB, cfg.DBPORT, cfg.DBName)
	timeout := time.Duration(cfg.SERVER_TIMEOUT) * time.Second

	//init product
	productRepo := repository.NewProductRepositoryImpl(db)
	productUseCase := usecase.NewProductUsecase(productRepo, timeout)
	productControllerInit := controller.NewProductController(*productUseCase)
	routesInit := app.RouteParams{
		Productcontroller: productControllerInit,
	}
	routesInit.Routes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.APPPORT)))
}
