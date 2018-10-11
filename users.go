package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//
//func GetUser() (*User, error) {
//	return GetUserById(-1)
//}

func GetUsers() []User {
	return nil
}

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

func (user *User) GetTeams(userId int) []Team {
	return nil
}

func (user *User) GetOwnedTeams(userId int) []Team {
	return nil
}
