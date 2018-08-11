package tiltify

// TODO according to the API, these don't just return the json object, but other stuff (meta) wrapped around it

func GetCampaignById(campaignId int) Campaign {
	return Campaign{}
}

func GetCampaignDonations(campaignId int) []Donation {
	return nil
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
