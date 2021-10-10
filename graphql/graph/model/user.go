package model

import "github.com/jackc/pgtype"

type User struct {
	ID    pgtype.UUID `json:"id"`
	email string      `json:"email"`
}
