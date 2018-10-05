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

func (tc *TiltifyClient) GetUser(userId int) (*User, error) {
	url := baseURL + "/users"
	if userId >= 0 {
		url += "/" + strconv.Itoa(userId)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &User{}, err
	}
	req.Header.Add("Authorization", "Bearer "+tc.AuthorizationToken)

	resp, err := tc.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return &User{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &User{}, err
	}

	var ur UserResponse
	err = json.Unmarshal(b, &ur)
	if err != nil {
		return &User{}, err
	}

	return ur.Data, nil
}

// curl https://tiltify.com/api/v3/campaigns/42

func (tc *TiltifyClient) GetCampaign(campaignId int) (*Campaign, error) {
	url := baseURL + "/campaigns/" + strconv.Itoa(campaignId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &Campaign{}, err
	}
	req.Header.Add("Authorization", "Bearer "+tc.AuthorizationToken)

	resp, err := tc.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return &Campaign{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Campaign{}, err
	}

	var ur CampaignResponse
	err = json.Unmarshal(b, &ur)
	if err != nil {
		return &Campaign{}, err
	}

	return ur.Data, nil
}
