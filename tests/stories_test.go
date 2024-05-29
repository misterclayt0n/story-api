package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/misterclayt0n/story-api/models"
	"github.com/misterclayt0n/story-api/routes"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func SetUpDatabase() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&models.Story{}, &models.User{})
	return db
}

func TestAddStory(t *testing.T) {
	r := SetUpRouter()
	db := SetUpDatabase()
	routes.InitializeStoryRoutes(r, db)

	story := `{"title":"Test Story","description":"This is a test story","category":"Test"}`
	req, _ := http.NewRequest("POST", "/stories", strings.NewReader(story))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Story")
}

func TestGetStories(t *testing.T) {
	r := SetUpRouter()
	db := SetUpDatabase()
	routes.InitializeStoryRoutes(r, db)

	req, _ := http.NewRequest("GET", "/stories", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
