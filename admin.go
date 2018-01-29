package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func updateFoodChoiceListHandler(context *gin.Context) {
	var payload UpdateFoodChoiceAdminPL
	err := context.ShouldBindJSON(&payload)
	if err == nil {
		session, _ = mgo.Dial(ConnectionString)
		defer session.Close()

		col = session.DB(DatabaseName).C("FoodChoices")
		for _, day := range payload.Days {
			err := col.Insert(day)
			if err != nil {
				context.JSON(200, gin.H{
					"status": "Error",
					"msg":    "Cant insert into database",
				})
			}
		}
		context.JSON(200, gin.H{
			"status": "Success",
			"msg":    "Updated food choice.",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getTodaySummaryHandler(context *gin.Context) {
	session, _ = mgo.Dial(ConnectionString)
	defer session.Close()

	col = session.DB(DatabaseName).C("UserFoodChoices")
	toDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.UTC)

	var result = []UserFoodChoice{}
	CheckError(col.Find(bson.M{
		"date":  bson.M{
			"$gt": time.Now().AddDate(0, 0, -1),
			"$lt": toDate,
		},
	}).All(&result))

	context.JSON(200, gin.H{
		"status": "Success",
		"result": result,
	})
}
