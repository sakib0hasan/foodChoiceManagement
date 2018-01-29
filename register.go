package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session
var col *mgo.Collection

const (
	ConnectionString = "localhost:27017"
	DatabaseName = "foodChoiceManagement"
	CollectionName = "Users"
)

func Register(context *gin.Context) {
	session, _ = mgo.Dial(ConnectionString)
	defer session.Close()

	col = session.DB(DatabaseName).C(CollectionName)

	var payload RegisterPL
	err := context.ShouldBindJSON(&payload)
	if err == nil {
		// check email
		var result = []User{}
		CheckError(col.Find(bson.M{"email": payload.Email}).All(&result))

		if len(result) > 0 {
			context.JSON(200, gin.H{
				"status": "Error",
				"msg":    "Email Already Exist.",
			})
		} else {
			password, _ := HashPassword(payload.Password)
			userPayload := User{payload.Name, payload.Email, password, "user"}
			err := col.Insert(userPayload)
			if err != nil {
				log.Fatal(err)
			}
			context.JSON(200, gin.H{
				"status": "Success",
				"msg":    "Account Created Successfully.",
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}