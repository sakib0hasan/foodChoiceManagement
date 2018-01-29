package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func Login(context *gin.Context) {
	session, _ = mgo.Dial(ConnectionString)
	defer session.Close()

	col = session.DB(DatabaseName).C(CollectionName)

	var payload LoginPL
	err := context.ShouldBindJSON(&payload)
	if err == nil {
		// check email
		var result = []User{}
		CheckError(col.Find(bson.M{"email": payload.Email}).Limit(1).All(&result))

		if len(result) > 0 {
			if CheckPasswordHash(payload.Password, result[0].Password) {
				token, _ := GenerateToken(payload.Email, result[0].Role)
				context.JSON(200, gin.H{
					"status": "Success",
					"msg":    "Logged in.",
					"role": result[0].Role,
					"token": token,
				})
			} else {
				context.JSON(200, gin.H{
					"status": "Error",
					"msg":    "Email or password is wrong.",
				})
			}
		}else{
			context.JSON(200, gin.H{
				"status": "Error",
				"msg":    "Email or password is wrong.",
			})
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
