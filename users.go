package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (user *User) GetCampaigns() ([]Campaign, error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Id) + "/campaigns"

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []Campaign{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Campaign{}, err
	}

	var csr CampaignsResponse
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return []Campaign{}, err
	}

	//TODO fix this crap
	for c := range csr.Data {
		fmt.Println(c)
	}
	return csr.Data, nil
}

func (user *User) GetCampaign(campaignId int) (*Campaign, error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Id) + "/campaigns/" + strconv.Itoa(campaignId)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return &Campaign{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Campaign{}, err
	}

	var ucr CampaignResponse
	err = json.Unmarshal(b, &ucr)
	if err != nil {
		return &Campaign{}, err
	}

	return ucr.Data, nil
}

// GET /users/1/teams

func (user *User) GetTeams() (teams []Team, err error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Id) + "/teams"

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []Team{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Team{}, err
	}

	var tsr TeamsResponse
	err = json.Unmarshal(b, &tsr)
	if err != nil {
		return []Team{}, err
	}

	return tsr.Data, nil
}

// GET /users/:id/owned-teams

func (user *User) GetOwnedTeams() (teams []Team, err error) {
	url := baseURL + "/users/" + strconv.Itoa(user.Id) + "/owned-teams"

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []Team{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Team{}, err
	}

	var tsr TeamsResponse
	err = json.Unmarshal(b, &tsr)
	if err != nil {
		return []Team{}, err
	}

	return tsr.Data, nil
}
