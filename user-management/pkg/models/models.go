package models

import "time"

type User struct {
	User_ID      uint      `json:"userid" gorm:"primaryKey;unique"`
	First_Name   string    `json:"firstname" gorm:"not null" validate:"min=5,max=20"`
	Last_Name    string    `json:"lastname" gorm:"not null" validate:"min=3"`
	User_Name    string    `json:"username" gorm:"not null;unique" validate:"min=5,max=20"`
	Password     string    `json:"password" gorm:"not null" validate:"min=5,max=20"`
	Email        string    `json:"email" gorm:"not null;unique" validate:"email"`
	IsBlocked    bool      `json:"isblocked" gorm:"default:false"`
	Phone_Number string    `json:"phonenumber" gorm:"not null;unique"`
	Otp          uint      `json:"otp" gorm:"not null"`
	Validate     bool      `json:"validate" gorm:"default:false"`
	Created_at   time.Time `json:"createdat"`
	Referal_Code string    `json:"referalcode" gorm:"not null"`
	Wallet       uint      `json:"wallet" gorm:"default:0"`
}
