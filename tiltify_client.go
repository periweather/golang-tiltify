package tiltify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var baseURL = "https://tiltify.com/api/v3"

type TiltifyClient struct {
	Client             *http.Client
	AuthorizationToken string
}

// Top Level User functions

func (tc *TiltifyClient) GetCurrentUser() (*User, error) {
	b, err := tc.makeRequest("/user")
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, err
	}

	user.tc = tc

	return &user, nil
}

func (tc *TiltifyClient) GetUser(userId int) (*User, error) {
	b, err := tc.makeRequest("/users/" + strconv.Itoa(userId))

	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return &User{}, err
	}

	user.tc = tc

	return &user, nil
}

func (tc *TiltifyClient) GetUsers() (*Users, error) {
	b, err := tc.makeRequest("/users")
	if err != nil {
		return nil, err
	}

	var users Users
	err = json.Unmarshal(b, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// curl https://tiltify.com/api/v3/campaigns/42

func (tc *TiltifyClient) GetCampaign(campaignId int) (*Campaign, error) {
	b, err := tc.makeRequest("/campaigns/" + strconv.Itoa(campaignId))
	if err != nil {
		return nil, err
	}

	var campaign Campaign
	err = json.Unmarshal(b, &campaign)
	if err != nil {
		return nil, err
	}

	campaign.tc = tc

	return &campaign, nil
}

// Helpers

func (tc *TiltifyClient) makeRequest(requestURL string) (b []byte, err error) {
	// TODO why is this url returning 404 instead of first 100 users?
	req, err := http.NewRequest(http.MethodGet, baseURL+requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+tc.AuthorizationToken)

	resp, err := tc.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
