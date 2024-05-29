package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/misterclayt0n/story-api/models"
	"gorm.io/gorm"
)

var jwtKey = []byte("very-secret-key-ma-gucci")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @ID register-user
// @Accept  json
// @Produce  json
// @Param   user     body    models.User     true        "User data"
// @Success 200 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context, db *gorm.DB) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db.Where("username = ?", creds.Username).First(&user)
	if user.Username != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	user = models.User{Username: creds.Username, Password: creds.Password}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user with username and password
// @ID login-user
// @Accept  json
// @Produce  json
// @Param   user     body    models.User     true        "User data"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context, db *gorm.DB) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db.Where("username = ?", creds.Username).First(&user)
	if user.Username == "" || user.Password != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
		c.Abort()
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}
