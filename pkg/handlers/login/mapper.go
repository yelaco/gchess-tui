package login

import (
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

func mapUserDaoToDomain(userDao UserDao) domains.User {
	return domains.User{
		UserId:    userDao.UserId,
		Username:  userDao.Username,
		Rating:    userDao.Rating,
		Email:     userDao.Email,
		CreatedAt: userDao.CreatedAt,
	}
}

func mapLoginDomainToDao(login domains.Login) LoginDao {
	return LoginDao{
		Username: login.Username,
		Password: login.Password,
	}
}
