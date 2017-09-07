package classis

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"bytes"
	"encoding/json"
	"github.com/hashicorp/go-uuid"
	"io"
	"io/ioutil"
	"time"
)

type Client struct {
	client *http.Client
	base   string
	token  string
}

func NewClientWith(url string, username string, password string) (*Client, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	login := Login{
		EmailAddress: username,
		Password:     password,
	}
	log.Errorf("Wtf %s", login)
	loginBytes, err := json.Marshal(login)
	if err != nil {
		return nil, err
	}
	loginReader := bytes.NewReader(loginBytes)

	response, err := netClient.Post(url+"/users/login", "application/json", loginReader)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	loginResponse := LoginResponse{}
	err = json.Unmarshal(contents, &loginResponse)
	if err != nil {
		return nil, err
	}
	return &Client{netClient, url, loginResponse.Token}, nil
}

func (c *Client) CreateSpotGroup(spotGroup SpotGroup) (string, error) {
	var a [2]interface{}
	generatedUID, err := uuid.GenerateUUID()
	a[0] = generatedUID
	a[1] = spotGroup
	spotBytes, _ := json.Marshal(a)
	spotReader := bytes.NewReader(spotBytes)
	response, err := c.PostCustom(c.base+"/methods/sgUpsert", spotReader)

	defer response.Body.Close()
	//contents, err := ioutil.ReadAll(response.Body)
	//spotGroupResponse := SpotGroup{}
	//err = json.Unmarshal(contents, &spotGroupResponse)

	if err != nil {
		log.Error(err)
	}

	return generatedUID, err
}

func (c *Client) DeleteSpotGroup(groupId string) error {
	var a [1]interface{}
	a[0] = groupId
	spotBytes, _ := json.Marshal(a)
	spotReader := bytes.NewReader(spotBytes)
	response, err := c.PostCustom(c.base+"/methods/spotGroup.remove", spotReader)

	defer response.Body.Close()
	//contents, err := ioutil.ReadAll(response.Body)
	//spotGroupResponse := SpotGroup{}
	//err = json.Unmarshal(contents, &spotGroupResponse)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (c *Client) PostCustom(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	return c.client.Do(req)
}
