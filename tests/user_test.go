package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/misterclayt0n/story-api/routes"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	r := SetUpRouter()
	db := SetUpDatabase()
	r.POST("/register", func(c *gin.Context) {
		routes.Register(c, db)
	})

	user := `{"username":"testuser","password":"testpass"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(user))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestLoginUser(t *testing.T) {
	r := SetUpRouter()
	db := SetUpDatabase()
	r.POST("/register", func(c *gin.Context) {
		routes.Register(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		routes.Login(c, db)
	})

	// Register the user first
	user := `{"username":"testuser","password":"testpass"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(user))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Now login with the same user
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(user))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
