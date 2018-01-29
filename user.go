package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/jinzhu/now"
)

func submitUserFoodChoiceHandler(context *gin.Context) {
	var payload UpdateFoodChoiceUserPL
	err := context.ShouldBindJSON(&payload)
	if err == nil {
		session, _ = mgo.Dial(ConnectionString)
		defer session.Close()

		col = session.DB(DatabaseName).C("UserFoodChoices")



		for _, day := range payload.Days {
			email, _ := ParseToken(context.GetHeader("Authorization"))
			var user = UserFoodChoice{
				day.Date,
				day.Food,
				email,
			}
			// remove current one
			var result = []UserFoodChoice{}
			CheckError(col.Find(bson.M{
				"date": day.Date,
				"email": email,
			}).All(&result))
			if(len(result)) > 0{
				CheckError(col.Remove(bson.M{
					"date": day.Date,
					"email": email,
				}))
			}
			// insert
			err := col.Insert(user)
			if err != nil {
				context.JSON(200, gin.H{
					"status": "Error",
					"msg":    "Cant insert into database",
				})
			}
		}
		context.JSON(200, gin.H{
			"status": "Success",
			"msg":    payload,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getUserFoodChoiceHandler(context *gin.Context) {
	session, _ = mgo.Dial(ConnectionString)
	defer session.Close()

	col = session.DB(DatabaseName).C("UserFoodChoices")
	email, _ := ParseToken(context.GetHeader("Authorization"))

	beginingOfWeek := now.BeginningOfWeek()
	var result = []UserFoodChoice{}
	CheckError(col.Find(bson.M{
		"date": bson.M{
			"$gt": beginingOfWeek,
		},
		"email": email,
	}).All(&result))

	context.JSON(200, gin.H{
		"status":         "Success",
		"result":         result,
		"beginingOfWeek": beginingOfWeek,
	})
}

func getCurrentWeekFoodListHandler(context *gin.Context) {

	now.FirstDayMonday = false

	session, _ = mgo.Dial(ConnectionString)
	defer session.Close()

	col = session.DB(DatabaseName).C("FoodChoices")

	beginingOfWeek := now.BeginningOfWeek()
	var result = []UpdateFoodChoice{}
	CheckError(col.Find(bson.M{
		"date": bson.M{
			"$gt": beginingOfWeek,
		},
	}).All(&result))

	context.JSON(200, gin.H{
		"status":         "Success",
		"result":         result,
		"beginingOfWeek": beginingOfWeek,
	})
}
