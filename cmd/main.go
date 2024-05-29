package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/misterclayt0n/story-api/docs"
	"github.com/misterclayt0n/story-api/models"
	"github.com/misterclayt0n/story-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Story API
// @version 1.0
// @description API for managing stories and generating content.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	db, err := gorm.Open(sqlite.Open("story_api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Story{}, &models.User{})

	r := gin.Default()

	routes.InitializeStoryRoutes(r, db)

	// swagger
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/register", func(c *gin.Context) {
		routes.Register(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		routes.Login(c, db)
	})

	r.Run(":8080")
}
