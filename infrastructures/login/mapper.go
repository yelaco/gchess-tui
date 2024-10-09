package login

import domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"

func mapUserDaoToDomain(userDao UserDao) domainlogin.User {
	return domainlogin.User{
		Username:  userDao.Username,
		Email:     userDao.Email,
		Rating:    userDao.Rating,
		CreatedAt: userDao.CreatedAt,
	}
}

func mapLoginDomainToDao(login domainlogin.Login) LoginDao {
	return LoginDao{
		Username: login.Username,
		Password: login.Password,
	}
}
