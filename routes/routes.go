package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	InitializeUserRoutes(router, db)
	InitializeStoryRoutes(router, db)
}
