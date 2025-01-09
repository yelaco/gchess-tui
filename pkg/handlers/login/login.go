package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yelaco/gchess-tui/pkg/app"
	domains "github.com/yelaco/gchess-tui/pkg/ui/domains"
)

type LoginDao struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// TODO: seperate user id and player id
type UserDao struct {
	UserId    string    `json:"player_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Rating    int64     `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

const LoginPath = "/login"

func LoginUser(info domains.Login) (domains.User, error) {
	u := app.GetConfig().ServiceHttpUrl.JoinPath(LoginPath)
	loginJson, err := json.Marshal(mapLoginDomainToDao(info))
	if err != nil {
		return domains.User{}, err
	}

	resp, err := http.Post(
		u.String(),
		"application/json",
		bytes.NewBuffer(loginJson),
	)
	if err != nil {
		return domains.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domains.User{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result UserDao
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return domains.User{}, err
	}

	return mapUserDaoToDomain(result), nil
}
