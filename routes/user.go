package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misterclayt0n/story-api/models"
	"gorm.io/gorm"
)

func InitializeUserRoutes(router *gin.Engine, db *gorm.DB) {
	userRoutes := router.Group("/users")
	{
		// @Summary Register a new user
		// @Description Register a new user with username and password
		// @ID create-user
		// @Accept  json
		// @Produce  json
		// @Param   user     body    models.User     true        "User data"
		// @Success 200 {object} models.User
		// @Router /users [post]
		userRoutes.POST("", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Create(&user)
			c.JSON(http.StatusOK, user)
		})

		// @Summary List all users
		// @Description Get all users
		// @ID list-users
		// @Produce  json
		// @Success 200 {array} models.User
		// @Router /users [get]
		userRoutes.GET("", Authenticate, func(c *gin.Context) {
			var users []models.User
			db.Find(&users)
			c.JSON(http.StatusOK, users)
		})

		// @Summary Get a user by ID
		// @Description Get a user by its ID
		// @ID get-user
		// @Produce  json
		// @Param   id     path    string     true        "User ID"
		// @Success 200 {object} models.User
		// @Router /users/{id} [get]
		userRoutes.GET("/:id", Authenticate, func(c *gin.Context) {
			var user models.User
			if err := db.First(&user, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		// @Summary Update a user by ID
		// @Description Update a user's username and password by its ID
		// @ID update-user
		// @Accept  json
		// @Produce  json
		// @Param   id     path    string     true        "User ID"
		// @Param   user     body    models.User     true        "User data"
		// @Success 200 {object} models.User
		// @Router /users/{id} [put]
		userRoutes.PUT("/:id", Authenticate, func(c *gin.Context) {
			var user models.User
			if err := db.First(&user, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		})

		// @Summary Delete a user by ID
		// @Description Delete a user by its ID
		// @ID delete-user
		// @Produce  json
		// @Param   id     path    string     true        "User ID"
		// @Success 200 {object} map[string]string
		// @Router /users/{id} [delete]
		userRoutes.DELETE("/:id", Authenticate, func(c *gin.Context) {
			if err := db.Delete(&models.User{}, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		})
	}
}
