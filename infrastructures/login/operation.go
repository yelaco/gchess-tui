package login

import domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"

type operation struct{}

func NewOperation() domainlogin.Operation {
	return operation{}
}

func (o operation) RequestLogin(domainlogin.Login) (domainlogin.User, error) {
	// TODO: implement login logic
	user := UserDao{}

	return mapUserDaoToDomain(user), nil
}
