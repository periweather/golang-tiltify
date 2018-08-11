package tiltify

type Image struct {
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
	Avatar Image // image object proper?
	About string
	TotalAmountRaised float32
	Social Social
}

type Team struct {
	Id int
	Username string
	Slug string
	URL string
	Avatar Image
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
	Avatar Image
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

type Challenge struct{
	Id int
	Name string
	Amount float32 // currency
	TotalAmountRaised float32
	Active bool
	ActivatesOn int64 // date
	CampaignId int
	EndsAt int64
	CreatedAt int64
	UpdatedAt int64
}

type Option struct{
	Id int
	PollId int
	Name string
	TotalAmountRaised float32 // int?
	CreatedAt int64 // date
	UpdatedAt int64
}


type Poll struct {
	Id int
	Name string
	Active bool
	CampaignId int
	CreatedAt int64
	UpdatedAt int64
	Options []Option
}

type Rewards struct {
	Id int
	Name string
	Description string
	Amount int
	Kind string
	Quantity int
	Remaining int
	FairMarketValue int
	Currency string
	ShippingAddressRequired bool
	ShippingNote string
	Image Image
	Active bool
	StartsAt int64
	CreatedAt int64
	UpdatedAt int64
}



