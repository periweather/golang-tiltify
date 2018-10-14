package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"reflect"
)

// TODO according to the API, these don't just return the json object, but other stuff (meta) wrapped around it

// TODO error checking
func (campaign *Campaign) GetCampaignDonations() (donations []Donation, err error) {
	// GET /campaigns/:id/donations

	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/donations"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return []Donation{}, err
	}

	resp, err := campaign.tiltifyClient.DoRequest(req)

	//resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []Donation{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Donation{}, err
	}

	var dsr DonationsResponse
	err = json.Unmarshal(b, &dsr)
	if err != nil {
		return []Donation{}, err
	}

	fmt.Println(dsr.Links)

	return dsr.Data, nil
}

func (campaign *Campaign) GetCampaignRewards(campaignId int) (rewards []Reward, err error) {
	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/rewards"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return []Reward{}, err
	}

	resp, err := campaign.tiltifyClient.DoRequest(req)

	defer resp.Body.Close()
	if err != nil {
		return []Reward{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Reward{}, err
	}

	var rsr RewardsResponse
	err = json.Unmarshal(b, &rsr)
	if err != nil {
		return []Reward{}, err
	}

	return rsr.Data, nil
}

func (campaign *Campaign) GetCampaignPolls() (polls []Poll, err error) {
	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/polls"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return []Poll{}, err
	}

	resp, err := campaign.tiltifyClient.DoRequest(req)
	defer resp.Body.Close()

	if err != nil {
		return []Poll{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Poll{}, err
	}

	var psr PollsResponse
	err = json.Unmarshal(b, &psr)
	if err != nil {
		return []Poll{}, err
	}

	return psr.Data, nil

}

func (campaign *Campaign) GetCampaignChallenges() (challenges []Challenge, err error) {
	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/challenges"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return []Challenge{}, err
	}

	resp, err := campaign.tiltifyClient.DoRequest(req)
	defer resp.Body.Close()

	if err != nil {
		return []Challenge{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Challenge{}, err
	}

	var csr ChallengesResponse
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return []Challenge{}, err
	}

	return csr.Data, nil
}

func (campaign *Campaign) GetCampaignSchedule() (scheds []Schedule, err error) {
	// GET /campaigns/:id/donations

	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/schedule"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return []Schedule{}, err
	}

	resp, err := campaign.tiltifyClient.DoRequest(req)
	defer resp.Body.Close()

	if err != nil {
		return []Schedule{}, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Schedule{}, err
	}

	var sr SchedulesResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		return []Schedule{}, err
	}

	return sr.Data, nil
}

// TODO check the proper url when i get back online
func (campaign *Campaign) GetCampaignsSupportingCampaigns() (campaigns []Campaign, err error) {
	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/supportingcampaigns"

	req, err := campaign.tiltifyClient.MakeRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	// TODO should i defer resp.Body after the err check in all the places?
	resp, err := campaign.tiltifyClient.DoRequest(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var csr CampaignsResponse
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return nil, err
	}

	return csr.Data, nil
}
