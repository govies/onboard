package repository

import (
	"context"
	"github.com/govies/framework/logger"
	"github.com/govies/onboard/domain"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	Logger *logger.Logger
}

func NewUserRepository(l *logger.Logger) *UserRepository {
	return &UserRepository{
		Logger: l,
	}
}

func (m *UserRepository) Fetch(ctx context.Context) (res []domain.User, err error) {
	err = mgm.Coll(&domain.User{}).SimpleFind(&res, bson.M{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *UserRepository) GetByID(ctx context.Context, id string) (res domain.User, err error) {
	err = mgm.Coll(&res).FindByID(id, &res)
	if err != nil {
		return res, err
	}
	return res, err
}