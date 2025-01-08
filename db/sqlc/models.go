// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID        int64              `json:"id"`
	Owner     string             `json:"owner"`
	Currency  string             `json:"currency"`
	Balance   int64              `json:"balance"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// only positive
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
