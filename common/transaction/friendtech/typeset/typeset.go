package typeset

import (
	"math/big"
	"time"
)

type FriendTechKeyTransaction struct {
	ID             int                `json:"id"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	Hash           string             `json:"hash"`
	Timestamp      int64              `json:"timestamp"`
	BlockNumber    int64              `json:"blockNumber"`
	FromAddress    string             `json:"fromAddress"`
	SubjectAddress string             `json:"subjectAddress"`
	IsBuy          bool               `json:"isBuy"`
	Amount         int64              `json:"amount"`
	Cost           big.Int            `json:"cost"`
	FromAccount    *FriendTechAccount `json:"fromAccount"`
	SubjectAccount *FriendTechAccount `json:"subjectAccount"`
}

type FriendTechAccount struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Address         string    `json:"address"`
	TwitterUsername string    `json:"twitterUsername"`
	TwitterPfpURL   string    `json:"twitterPfpUrl"`
	ProfileChecked  bool      `json:"profileChecked"`
	Price           float64   `json:"price"`
	Supply          int       `json:"supply"`
	Holders         int       `json:"holders"`
}
