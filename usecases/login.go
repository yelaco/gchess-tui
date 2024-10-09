package usecases

import (
	"github.com/yelaco/gchess-tui/domains/dtos"
	domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"
)

type LoginUsecaseInterface interface {
	LoginUser(dtos.Login) (domainlogin.User, error)
}

type LoginUsecase struct {
	operation domainlogin.Operation
}

func NewLoginUsecase(loginOperation domainlogin.Operation) LoginUsecaseInterface {
	return &LoginUsecase{
		operation: loginOperation,
	}
}

func (u *LoginUsecase) LoginUser(info dtos.Login) (domainlogin.User, error) {
	user, err := u.operation.RequestLogin(domainlogin.Login{
		Username: info.Username,
		Password: info.Password,
	})
	if err != nil {
		return domainlogin.User{}, err
	}

	return user, nil
}
