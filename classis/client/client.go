package classis

import (
	"errors"
	"github.com/hashicorp/go-uuid"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	request *resty.Request
	base    string
	token   string
}

func NewClientWith(url string, username string, password string) (*Client, error) {

	loginResponse := LoginResponse{}
	apiError := APIError{}
	_, err := resty.R().
		SetBody(Login{EmailAddress: username, Password: password}).
		SetResult(&loginResponse).
		SetError(&apiError).
		Post(url + "/users/login")

	if apiError.Error != 0 {
		return nil, errors.New(apiError.Reason)
	}
	if err != nil {
		return nil, err
	}

	request := resty.R().
		SetHeader("Authorization", "Bearer "+loginResponse.Token)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{netClient, request, url, loginResponse.Token}, nil
}

func (c *Client) CreateSpotGroup(spotGroup SpotGroup) (string, error) {
	generatedUID, _ := uuid.GenerateUUID()
	apiError := APIError{}
	_, err := c.request.
		SetBody(spotGroup).
		SetError(&apiError).
		Post(c.base + "/spot-groups/" + generatedUID)
	if err != nil {
		return "", err
	}
	if apiError.Error != 0 {
		return "", errors.New(apiError.Reason)
	}

	return generatedUID, err
}

func (c *Client) DeleteSpotGroup(groupId string) error {
	apiError := APIError{}
	_, err := c.request.
		SetError(&apiError).
		Delete(c.base + "/spot-groups/" + groupId)
	if apiError.Error != 0 {
		return errors.New(apiError.Reason)
	}
	return err
}
