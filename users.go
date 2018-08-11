package tiltify

func GetUser() User {
	return User{}
}

func GetUsers() []User {
	return nil
}

func GetUserById(userId int) User {
	return User{}
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
