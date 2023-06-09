// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int32          `json:"ID"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Birthday  time.Time      `json:"birthday"`
	Height    float64        `json:"height"`
	Status    sql.NullString `json:"status"`
	Gender    sql.NullBool   `json:"gender"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
