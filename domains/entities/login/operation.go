package login

type Operation interface {
	RequestLogin(Login) (User, error)
}
