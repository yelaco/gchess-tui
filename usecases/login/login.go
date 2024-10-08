package login

import domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"

type LoginUsecaseInterface interface {
	Login(username, password string) error
}

type LoginUsecase struct {
	operation domainlogin.Operation
}

func (u *LoginUsecase) Login(username, password string) (domainlogin.User, error) {
	user, err := u.operation.RequestLogin(domainlogin.Login{
		Username: username,
		Password: password,
	})
	if err != nil {
		return domainlogin.User{}, err
	}

	return user, nil
}
