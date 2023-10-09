package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint64       `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt time.Time    `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time    `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt sql.NullTime `xorm:"TIMESTAMP deleted_at"`
	Email     string       `xorm:"not null VARCHAR(100) UNIQUE email"`
}

type UserLoginRequest struct {
	Email string `json:"email" validate:"required"`
	Code  string `json:"code" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func (User) TableName() string {
	return "users"
}
