package tiltify

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// TODO according to the API, these don't just return the json object, but other stuff (meta) wrapped around it

// TODO error checking
func (campaign *Campaign) GetCampaignDonations() (donations []DonationData, err error) {
	// GET /campaigns/:id/donations

	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/donations")
	if err != nil {
		return nil, err
	}

	var dsr Donations
	err = json.Unmarshal(b, &dsr)
	if err != nil {
		return nil, err
	}

	fmt.Println(dsr.Links)

	return dsr.Data, nil
}

func (campaign *Campaign) GetCampaignRewards(campaignId int) (rewards []RewardData, err error) {
	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/rewards")
	if err != nil {
		return nil, err
	}

	var rsr Rewards
	err = json.Unmarshal(b, &rsr)
	if err != nil {
		return nil, err
	}

	return rsr.Data, nil
}

func (campaign *Campaign) GetCampaignPolls() (polls []PollData, err error) {
	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/polls")
	if err != nil {
		return nil, err
	}

	var psr Polls
	err = json.Unmarshal(b, &psr)
	if err != nil {
		return nil, err
	}

	return psr.Data, nil

}

func (campaign *Campaign) GetCampaignChallenges() (challenges []Challenge, err error) {
	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/challenges")
	if err != nil {
		return nil, err
	}

	var csr ChallengesResponse
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return nil, err
	}

	return csr.Data, nil
}

func (campaign *Campaign) GetCampaignSchedule() (scheds *Schedules, err error) {
	// GET /campaigns/:id/donations
	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/schedule")
	if err != nil {
		return nil, err
	}

	var sr Schedules
	err = json.Unmarshal(b, &sr)
	if err != nil {
		return nil, err
	}

	return &sr, nil
}

// TODO check the proper url when i get back online
func (campaign *Campaign) GetCampaignsSupportingCampaigns() (campaigns *Campaigns, err error) {
	b, err := campaign.tc.makeRequest("/campaigns/" + strconv.Itoa(campaign.Data.Id) + "/supporting-campaigns")
	if err != nil {
		return nil, err
	}

	var csr Campaigns
	err = json.Unmarshal(b, &csr)
	if err != nil {
		return nil, err
	}

	return &csr, nil
}
