package domain

import (
	"context"
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	//ID     uint   `json:"id" gorm:"primary_key"`
	Email        string `bson:"email" json:"email"`
	PasswordHash string `bson:"password_hash" json:"password_hash"`
}

type UserEntity interface {
	Fetch(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
}

type UserRepository interface {
	Fetch(ctx context.Context) (res []User, err error)
	GetByID(ctx context.Context, id string) (User, error)
}
