package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	id             string `gorm:"primaryKey"`
	clientId       string
	balance        int64
	virtualBalance int64
}

type Clients struct {
	gorm.Model
	id    string `gorm:"primaryKey"`
	name  string
	email string
}

type Transactions struct {
	gorm.Model
	id               string `gorm:"primaryKey"`
	transaction_type string
	amount           int64
	account_id       string
}

type ScheduledTransactions struct {
	gorm.Model
	id               string `gorm:"primaryKey"`
	transaction_type string
	amount           int64
	account_id       string
	scheduled_for    time.Time
}
