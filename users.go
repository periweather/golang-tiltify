package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (user *User) GetCampaigns() ([]CampaignData, error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Data.Id) + "/campaigns"

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var csr Campaigns
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return nil, err
	}

	//TODO fix this crap
	for c := range csr.Data {
		fmt.Println(c)
	}
	return csr.Data, nil
}

func (user *User) GetCampaign(campaignId int) (*Campaign, error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Data.Id) + "/campaigns/" + strconv.Itoa(campaignId)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ucr Campaign
	err = json.Unmarshal(b, &ucr)
	if err != nil {
		return nil, err
	}

	return &ucr, nil
}

// GET /users/1/teams

func (user *User) GetTeams() (*Teams, error) {
	b, err := user.tc.makeRequest("/users/" + strconv.Itoa(user.Data.Id) + "/teams")
	if err != nil {
		return nil, err
	}

	var teams Teams
	err = json.Unmarshal(b, &teams)
	if err != nil {
		return nil, err
	}

	return &teams, nil
}

// GET /users/:id/owned-teams

func (user *User) GetOwnedTeams() (*Teams, error) {
	b, err := user.tc.makeRequest("/users/" + strconv.Itoa(user.Data.Id) + "/owned-teams")
	if err != nil {
		return nil, err
	}

	var teams Teams
	err = json.Unmarshal(b, &teams)
	if err != nil {
		return nil, err
	}

	return &teams, nil
}
