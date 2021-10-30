package entity

import (
	"context"
	"github.com/govies/onboard/domain"
)

type UserEntity struct {
	userRepo domain.UserRepository
}

func NewUserEntity(a domain.UserRepository) domain.UserEntity {
	return &UserEntity{
		userRepo: a,
	}
}
func (a *UserEntity) Fetch(c context.Context) (res []domain.User, err error) {
	res, err = a.userRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}

func (a *UserEntity) GetByID(c context.Context, id string) (res domain.User, err error) {
	res, err = a.userRepo.GetByID(c, id)
	return
}