package schemas

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	id             string
	clientId       string
	balance        int64
	virtualBalance int64
}

type Client struct {
	gorm.Model
	id    string
	name  string
	email string
}
