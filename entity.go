package tiltify

type Image struct {
	Src    string
	Alt    string
	Width  int
	Height int
}

type Social struct {
	Facebook string
	Twitch   string
	Twitter  string
	Website  string
	Youtube  string
}

type Livestream struct {
	Type    string
	Channel string
}

type Option struct {
	Id                int
	PollId            int
	Name              string
	TotalAmountRaised float32 // int?
	CreatedAt         int64   // date
	UpdatedAt         int64
}

type Colors struct {
	Background string
	Highlight  string
}

type Settings struct {
	Colors          *Colors
	HeaderIntro     string
	HeaderTitle     string
	FooterCopyright string
	FindOutMoreLink string
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	Region       string
	PostalCode   int
	Country      string
}

type Campaign struct {
	Id                     int
	Name                   string
	Slug                   string
	URL                    string
	StartsAt               int64
	EndsAt                 int64
	Description            string
	Avatar                 Image
	CauseId                int
	FundraisingEventId     int
	FundraiserGoalAmount   float32
	OriginalGoalAmount     float32
	AmountRaised           float32
	SupportingAmountRaised float32
	TotalAmountRaised      float32
	Supportable            bool
	Status                 string
	User                   *User
	Team                   *Team
	Livestream             *Livestream
}

type Cause struct {
	Id                 int
	Name               string
	LegalName          string
	Slug               string
	Currency           string
	About              string
	Video              string
	Image              Image
	Avatar             Image
	Logo               *Image
	Banner             *Image
	ContactEmail       string
	PaypalEmail        string
	PaypalCurrencyCode string
	Social             *Social
	Settings           *Settings
	Status             string
	StripeConnected    bool
	MailchimpConnected bool
	Address            *Address
}

type Challenge struct {
	Id                int
	Name              string
	Amount            float32 // currency
	TotalAmountRaised float32
	Active            bool
	ActivatesOn       int64 // date
	CampaignId        int
	EndsAt            int64
	CreatedAt         int64
	UpdatedAt         int64
}

type FundraisingEvent struct {
	Id           int
	Name         string
	Slug         string
	URL          string
	Description  string
	Video        *Image // do we want to rename the Image struct?
	Image        *Image
	Avatar       *Image
	Logo         *Image
	Banner       *Image
	BannerTitle  string
	BannerIntro  string
	Currency     string
	Goal         float32
	AmountRaised float32
	StartsOn     string // ISO formatted date instead of the epoc time for others?
	EndsOn       string
	CauseId      int
}

type Poll struct {
	Id         int
	Name       string
	Active     bool
	CampaignId int
	CreatedAt  int64
	UpdatedAt  int64
	Options    []Option
}

type Reward struct {
	Id                      int
	Name                    string
	Description             string
	Amount                  int
	Kind                    string
	Quantity                int
	Remaining               int
	FairMarketValue         int
	Currency                string
	ShippingAddressRequired bool
	ShippingNote            string
	Image                   Image
	Active                  bool
	StartsAt                int64
	CreatedAt               int64
	UpdatedAt               int64
}

type Schedule struct {
	Id          int
	Name        string
	Description string
	StartsAt    int64
}

type Team struct {
	Id       int
	Username string
	Slug     string
	URL      string
	Avatar   *Image
}

type Donation struct {
	Id           int     `json:"id"`
	Amount       float32 `json:"amount"`
	Name         string  `json:"name"`
	DonorComment string  `json:"donorComment"`
	CreatedAt    int64   `json:"createdAt"`
}

type User struct {
	Id                int     `json:"id"`
	Username          string  `json:"username"`
	Slug              string  `json:"slug"`
	URL               string  `json:"url"`
	Avatar            *Image  `json:"avatar"` // image object proper?
	About             string  `json:"about"`
	TotalAmountRaised float32 `json:"totalAmountRaised"`
	Social            *Social `json:"social"`
}

type UserResponse struct {
	Meta *Meta `json:"meta"`
	Data *User `json:"data"`
}

type Meta struct {
	Status int `json:"status"`
}
