package usecase

import (
	"test-lion-superindo/lib/helper"
	"test-lion-superindo/package/auth/model"
	"test-lion-superindo/package/auth/repository"
)

type authUsecase struct {
	repo repository.AuthRepo
}

func NewAuthUsecase(repo repository.AuthRepo) AuthUsecase {
	return &authUsecase{repo}
}

type AuthUsecase interface {
	Registration(user model.UserAuth) error
	Login(user model.UserAuth) (string, bool, error)
}

func (uc *authUsecase) Registration(user model.UserAuth) error {
	return uc.repo.AddUserAuth(user)
}
func (uc *authUsecase) Login(user model.UserAuth) (string, bool, error) {
	isMatch := false
	users, err := uc.repo.CheckUser(user)
	token := ""

	if isMatch = len(users) > 0; isMatch {
		token, err = helper.TokenConfig(users[0].Username)
	} else {
		return "", false, err
	}

	return token, isMatch, err
}
