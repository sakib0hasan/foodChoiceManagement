package main

import (
	"io/ioutil"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(email string, role string) (string, error) {
	privateKey, err := ioutil.ReadFile("keys/app.rsa")
	if err != nil {
		fmt.Println("Error reading private key")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(privateKey)

	return tokenString, err
}

func ParseToken(myToken string) (string, error) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		privateKey, err := ioutil.ReadFile("keys/app.rsa")
		if err != nil {
			fmt.Println("Error reading private key")
			return "", err
		}
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	} else {
		return "", err
	}
}

func TokenMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	fmt.Println(token)
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