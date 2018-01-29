package main

import "time"

type RegisterPL struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginPL struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateFoodChoiceAdminPL struct {
	Days []struct {
		Date time.Time `json:"date" binding:"required" time_format:"2006-01-02"`
		Food []string  `json:"food" binding:"required"`
	} `json:"Days" binding:"required"`
}
type UpdateFoodChoiceUserPL struct {
	Days []struct {
		Date time.Time `json:"date" binding:"required" time_format:"2006-01-02"`
		Food string    `json:"food" binding:"required"`
	} `json:"Days" binding:"required"`
}

type UpdateFoodChoice struct {
	Date time.Time `json:"date" binding:"required" time_format:"2006-01-02"`
	Food []string  `json:"food" binding:"required"`
}

type UserFoodChoice struct {
	Date  time.Time `json:"date" binding:"required" time_format:"2006-01-02"`
	Food  string    `json:"food" binding:"required"`
	Email string    `json:"email"`
}
