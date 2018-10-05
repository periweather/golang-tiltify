package tiltify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TODO according to the API, these don't just return the json object, but other stuff (meta) wrapped around it

// TODO error checking
func (campaign *Campaign) GetCampaignDonations() []Donation {
	// GET /campaigns/:id/donations

	url := baseURL + "/campaigns/" + strconv.Itoa(campaign.Id) + "/donations"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []Donation{}
	}
	req.Header.Add("Authorization", "Bearer 924f576d2ef50a6c252957a008e5bfbc2232f61c3f2354abfe33af19dc1546f3")

	client := http.DefaultClient

	resp, err := client.Do(req)

	//resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return []Donation{}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Donation{}
	}

	var dsr DonationsResponse
	err = json.Unmarshal(b, &dsr)
	if err != nil {
		return []Donation{}
	}

	fmt.Println(string(b))

	return dsr.Data
}

func GetCampaignRewards(campaignId int) []Reward {
	return nil
}

func GetCampaignPolls(campaignId int) []Poll {
	return nil
}

func GetCampaignChallenges(campaignId int) []Challenge {
	return nil
}

func GetCampaignSchedule(campaignId int) Schedule {
	return Schedule{}
}

func GetCampaignsSupportingCampaigns(campaignId int) []Campaign {
	return nil
}
