package main

import (
	"github.com/DaianaServin/ProyectoApi/cmd/config"
	"github.com/DaianaServin/ProyectoApi/cmd/server/routes"
	"github.com/DaianaServin/ProyectoApi/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func main() {
	config.LoadConfigFromFile()
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}