package tiltify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var baseURL = "https://tiltify.com/api/v3"

func GetUser() (*User, error) {
	return GetUserById(-1)
}

func GetUsers() []User {
	return nil
}

func GetUserById(userId int) (*User, error) {
	url := baseURL + "/user"
	if userId >= 0 {
		url += "/" + strconv.Itoa(userId)
	}
	resp, err := http.Get(url)
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

func GetCampaigns(userId int) []Campaign {
	return nil
}

func GetCampaign(userId int, campaignId int) Campaign {
	return Campaign{}
}

func GetTeams(userId int) []Team {
	return nil
}

func GetOwnedTeams(userId int) []Team {
	return nil
}
