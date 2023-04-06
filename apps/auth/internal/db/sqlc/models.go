// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Session struct {
	ID           pgtype.UUID        `json:"id"`
	Username     string             `json:"username"`
	RefreshToken string             `json:"refresh_token"`
	UserAgent    string             `json:"user_agent"`
	ClientIp     string             `json:"client_ip"`
	IsBlocked    bool               `json:"is_blocked"`
	ExpiresAt    pgtype.Timestamptz `json:"expires_at"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	Username          string             `json:"username"`
	HashedPassword    string             `json:"hashed_password"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}