package main

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const APP_KEY = "randomtokenstringrandomtokenstringrandomtokenstringrandomtokenstringrandomtokenstring"

func GenerateToken(email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	return tokenString, err
}

func ParseToken(myToken string) (string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(APP_KEY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	} else {
		return "", err
	}
}

func TokenMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != ""{
		_, err := ParseToken(token)
		if err != nil{
			c.AbortWithStatusJSON(401, gin.H{
				"status": "Error",
				"msg":    "Token invalid",
			})
		}else{
			c.Next()
		}
	}else{
		c.AbortWithStatusJSON(401, gin.H{
			"status": "Error",
			"msg":    "Token not given",
		})
	}
	c.Next()
}