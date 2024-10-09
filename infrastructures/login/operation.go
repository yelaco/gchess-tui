package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"
)

type operation struct {
	serviceUrl *url.URL
}

func NewOperation(serviceUrl *url.URL) domainlogin.Operation {
	return operation{
		serviceUrl: serviceUrl,
	}
}

const LoginPath = "/login"

func (o operation) RequestLogin(info domainlogin.Login) (domainlogin.User, error) {
	u := o.serviceUrl.JoinPath(LoginPath)
	loginJson, err := json.Marshal(mapLoginDomainToDao(info))
	if err != nil {
		return domainlogin.User{}, err
	}

	resp, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(loginJson),
	)
	if err != nil {
		return domainlogin.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domainlogin.User{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result UserDao
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return domainlogin.User{}, err
	}
	return mapUserDaoToDomain(result), nil
}
