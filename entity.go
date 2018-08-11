package tiltify

type Avatar struct {
	Src string
	Alt string
	Width int
	Height int
}

type Social struct {
	Facebook string
	Twitch string
	Twitter string
	Website string
	Youtube string
}

type User struct {
	Id int
	Username string
	Slug string
	URL string
	Avatar Avatar
	About string
	TotalAmountRaised float32
	Social Social
}

type Team struct {
	Id int
	Username string
	Slug string
	URL string
	Avatar Avatar
}

type Livestream struct {
	Type string
	Channel string
}

type Campaign struct {
	Id int
	Name string
	Slug string
	URL string
	StartsAt int64
	EndsAt int64
	Description string
	Avatar Avatar
	CauseId int
	FundraisingEventId int
	FundraiserGoalAmount float32
	OriginalGoalAmount float32
	AmountRaised float32
	SupportingAmountRaised float32
	TotalAmountRaised float32
	Supportable bool
	Status string
	User User
	Team Team
	Livestream Livestream
}



