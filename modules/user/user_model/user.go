package user_model

import "delivery-golang/common"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"role" gorm:"column:role;"`
	Salt            string `json:"-" gorm:"column:salt;"`
	Phone           string `json:"phone" gorm:"column:phone;"`
}
