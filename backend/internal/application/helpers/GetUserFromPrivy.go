package helpers

import (
	"encoding/json"
	"errors"
	"ethglobal-nd-backend/internal/httpdtos"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PrivyConfig struct {
	BaseUri  string
	UserName string
	Password string
	AppId    string
}

func NewPrivyConfig(baseUri, userName, password, appId string) *PrivyConfig {
	return &PrivyConfig{
		BaseUri:  baseUri,
		UserName: userName,
		Password: password,
		AppId:    appId,
	}
}

func (pc *PrivyConfig) GetPrivyUser(userID string) (*httpdtos.UserDTO, error) {
	url := pc.BaseUri + userID

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(pc.UserName, pc.Password)
	req.Header.Set("privy-app-id", pc.AppId)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Privy API error response: %s\n", string(body))
		return nil, fmt.Errorf("failed to fetch user: status %d: %s", resp.StatusCode, string(body))
	}

	var user httpdtos.UserDTO
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, errors.New("unable to parse user response: " + err.Error())
	}

	return &user, nil
}
