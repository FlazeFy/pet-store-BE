package models

type (
	UserLogin struct {
		Username string `json:"customers_slug" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserRegister struct {
		Username     string `json:"customers_slug" binding:"required"`
		Password     string `json:"password" binding:"required"`
		Email        string `json:"email" binding:"required"`
		CustomerName string `json:"customers_name" binding:"required"`
		CreatedBy    string `json:"created_by" binding:"nullable"`
	}
)
