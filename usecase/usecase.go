package usecase

import (
	"casbin/models"
	"casbin/repository"
	"context"
	"time"
)

type (
	casbinUsecase struct {
		mysql          repository.Repository
		contextTimeOut time.Duration
	}

	Usecase interface {
		GetPersons(c context.Context) ([]*models.Person, error)
		GetUser(c context.Context, username, password string) (*models.User, error)
	}
)

func NewUsecaseConfig(mysql repository.Repository, timeout time.Duration) Usecase {
	return &casbinUsecase{mysql, timeout}
}

func (u *casbinUsecase) GetPersons(c context.Context) ([]*models.Person, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeOut)
	defer cancel()
	persons, err := u.mysql.GetPersons(ctx)
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (u *casbinUsecase) GetUser(c context.Context, username, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeOut)
	defer cancel()

	user, err := u.mysql.GetUser(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
