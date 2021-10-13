package usecase

import (
	"context"

	"github.com/egaevan/online-learning/model"
	"github.com/egaevan/online-learning/repository"
	log "github.com/sirupsen/logrus"
)

type User struct {
	UserRepo repository.UserRepository
}

func NewUser(userRepo repository.UserRepository) UserUsecae {
	return &User{
		UserRepo: userRepo,
	}
}

func (u *User) Login(ctx context.Context, user model.User) (model.User, error) {
	user, err := u.UserRepo.FindOne(ctx, user.Email, user.Password)
	if err != nil {
		log.Error(err)
		return user, err
	}

	return user, nil
}

func (u *User) CreateUser(ctx context.Context, user model.User) error {
	err := u.UserRepo.Store(ctx, user)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *User) DeleteUser(ctx context.Context, userID int) error {

	err := u.UserRepo.Delete(ctx, userID)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
