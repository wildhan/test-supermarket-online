package main

import (
	"fmt"
	"net/http"
	"os"
	"test-lion-superindo/config/database"
	"test-lion-superindo/lib/helper"
	"test-lion-superindo/lib/log"

	authHandler "test-lion-superindo/package/auth/handler"
	authRepository "test-lion-superindo/package/auth/repository"
	authUsecase "test-lion-superindo/package/auth/usecase"

	merchandiseHandler "test-lion-superindo/package/merchandise/handler"
	merchandiseRepository "test-lion-superindo/package/merchandise/repository"
	merchandiseUsecase "test-lion-superindo/package/merchandise/usecase"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Error(fmt.Sprintf("Failed load .env: %v", err.Error()))
		os.Exit(2)
	}

	dbConn := database.CreateConnection(os.Getenv("DB_STRING"))

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	// AUTH
	authRepo := authRepository.NewAuthRepo(dbConn)
	authUC := authUsecase.NewAuthUsecase(authRepo)
	authHandler.NewAuthHandler(authUC).Mount(e.Group("/auth"))

	// MERCHANDISE
	merchandiseGroup := e.Group("/merchandise")
	merchandiseGroup.Use(echojwt.WithConfig(helper.JWTAuth()))
	merchandiseRepo := merchandiseRepository.NewMerchandiseRepo(dbConn)
	merchandiseUC := merchandiseUsecase.NewMerchandiseUsecase(merchandiseRepo)
	merchandiseHandler.NewMerchandiseHandler(merchandiseUC).Mount(merchandiseGroup)

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		log.Error(fmt.Sprintf("Failed start echo: %v", err.Error()))
	}
}
