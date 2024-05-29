package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misterclayt0n/story-api/gemini"
	"github.com/misterclayt0n/story-api/models"
	"gorm.io/gorm"
)

func InitializeStoryRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	storyRoutes := router.Group("/stories")
	{
		// @Summary Create a new story
		// @Description Create a new story with title, description, and category
		// @ID create-story
		// @Accept  json
		// @Produce  json
		// @Param   story     body    models.Story     true        "Story data"
		// @Success 200 {object} models.Story
		// @Router /stories [post]
		storyRoutes.POST("", func(c *gin.Context) {
			var story models.Story
			if err := c.ShouldBindJSON(&story); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			db.Create(&story)
			c.JSON(http.StatusOK, story)
		})

		// @Summary Get all stories
		// @Description Get all stories
		// @ID get-stories
		// @Produce  json
		// @Success 200 {array} models.Story
		// @Router /stories [get]
		storyRoutes.GET("", func(c *gin.Context) {
			var stories []models.Story
			db.Find(&stories)
			c.JSON(http.StatusOK, stories)
		})

		// @Summary Get a story by ID
		// @Description Get a story by its ID
		// @ID get-story-by-id
		// @Produce  json
		// @Param   id     path    string     true        "Story ID"
		// @Success 200 {object} models.Story
		// @Router /stories/{id} [get]
		storyRoutes.GET("/:id", func(c *gin.Context) {
			var story models.Story

			if err := db.First(&story, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
				return
			}

			c.JSON(http.StatusOK, story)
		})

		// @Summary Update a story by ID
		// @Description Update a story by its ID
		// @ID update-story-by-id
		// @Accept  json
		// @Produce  json
		// @Param   id     path    string     true        "Story ID"
		// @Param   story     body    models.Story     true        "Story data"
		// @Success 200 {object} models.Story
		// @Router /stories/{id} [put]
		storyRoutes.PUT("/:id", func(c *gin.Context) {
			var story models.Story

			if err := db.First(&story, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
				return
			}

			if err := c.ShouldBindJSON(&story); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			db.Save(&story)
			c.JSON(http.StatusOK, story)
		})

		// @Summary Delete a story by ID
		// @Description Delete a story by its ID
		// @ID delete-story-by-id
		// @Produce  json
		// @Param   id     path    string     true        "Story ID"
		// @Success 200 {object} map[string]string
		// @Router /stories/{id} [delete]
		storyRoutes.DELETE("/:id", func(c *gin.Context) {
			if err := db.Delete(&models.Story{}, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Story deleted"})
		})

		// @Summary Generate content for a story
		// @Description Generate content for a story based on a prompt
		// @ID generate-story-content
		// @Accept  json
		// @Produce  json
		// @Param   id     path    string     true        "Story ID"
		// @Param   prompt     body    map[string]string     true        "Prompt data"
		// @Success 200 {object} models.Story
		// @Router /stories/{id}/generate [post]
		storyRoutes.POST("/:id/generate", func(c *gin.Context) {
			var prompt struct {
				Prompt string `json:"prompt"`
			}

			if err := c.ShouldBindJSON(&prompt); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			var story models.Story
			if err := db.First(&story, "id = ?", c.Param("id")).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
				return
			}

			generatedContent, err := gemini.GenerateStory(prompt.Prompt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate story content"})
				return
			}

			story.Description += " " + generatedContent
			db.Save(&story)

			c.JSON(http.StatusOK, story)
		})
	}
}
