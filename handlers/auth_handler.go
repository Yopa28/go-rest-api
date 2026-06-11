package handlers

import (
	"database/sql"
	"os"
	"time"

	"go-rest-api/models"
	"go-rest-api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	
)

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if newUser.Name == "" {
		c.JSON(400, gin.H{
			"message": "Name is required",
		})
		return
	}

	if newUser.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}

	if newUser.Password == "" {
		c.JSON(400, gin.H{
			"message": "Password is required",
		})
		return
	}

	if len(newUser.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "Password must be at least 6 characters",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	newUser.Password = string(hashedPassword)

	if err := repositories.CreateUser(&newUser); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":    newUser.ID,
			"name":  newUser.Name,
			"email": newUser.Email,
		},
	})
}

func Login(c *gin.Context) {
	var loginInput models.User

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if loginInput.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}
	user, err := repositories.GetUserByEmail(loginInput.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{
				"message": "Invalid email or password",
			})
			return
		}

		c.JSON(500, gin.H{
			"message": "Failed to login",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role" :  user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	

	secretKey := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to generate token",
		})
		return
	}


	c.JSON(200, gin.H{
		"message": "Login successfully",
		"token":   tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role" : user.Role,	
		},
	})
}

func Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	c.JSON(200, gin.H{
		"message": "Profile accessed successfully",
		"user": gin.H{
			"user_id": userID,
			"email":   email,
		},
	})
}

