package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	id             string `json: "id" gorm: "primaryKey"`
	clientId       string `json: "clientId"`
	balance        int64  `json: "balance"`
	virtualBalance int64  `json: "virtualBalance"`
}

type Clients struct {
	gorm.Model
	id    string `json: "id" gorm: "primaryKey"`
	name  string `json: "name"`
	email string `json: "email"`
}

type Transactions struct {
	gorm.Model
	id               string `json: "id" gorm: "primaryKey"`
	transaction_type string `json: "transaction_type"`
	amount           int64  `json: "amount"`
	account_id       string `json: "account_id"`
}

type ScheduledTransactions struct {
	gorm.Model
	id               string    `json: "id" gorm: "primaryKey"`
	transaction_type string    `json: "transaction_type"`
	amount           int64     `json: "amount"`
	account_id       string    `json: "account_id"`
	scheduled_for    time.Time `json: "scheduled_for"`
}
