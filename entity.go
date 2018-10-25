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

type CampaignData struct {
	Id                     int         `json:"id"`
	Name                   string      `json:"name"`
	Slug                   string      `json:"slug"`
	URL                    string      `json:"url"`
	StartsAt               int64       `json:"startsAt"`
	EndsAt                 int64       `json:"endsAt"`
	Description            string      `json:"description"`
	Avatar                 *Image      `json:"avatar"`
	CauseId                int         `json:"causeId"`
	FundraisingEventId     int         `json:"fundraisingEventId"`
	FundraiserGoalAmount   float32     `json:"fundraiserGoalAmount"`
	OriginalGoalAmount     float32     `json:"originalGoalAmount"`
	AmountRaised           float32     `json:"amountRaised"`
	SupportingAmountRaised float32     `json:"supportingAmountRaised"`
	TotalAmountRaised      float32     `json:"totalAmountRaised"`
	Supportable            bool        `json:"supportable"`
	Status                 string      `json:"status"`
	User                   *UserData   `json:"user"`
	Team                   *Team       `json:"team"`
	Livestream             *Livestream `json:"livestream"`
}

type Campaign struct {
	Meta *Meta         `json:"meta"`
	Data *CampaignData `json:"data"`

	tc *TiltifyClient
}

type Links struct {
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Campaigns struct {
	Meta  *Meta          `json:"meta"`
	Data  []CampaignData `json:"data"`
	Links *Links         `json:"links"`
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

type ChallengesResponse struct {
	Meta  *Meta       `json:"meta"`
	Data  []Challenge `json:"data"`
	Links *Links      `json:"links"'`
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

type PollData struct {
	Id         int
	Name       string
	Active     bool
	CampaignId int
	CreatedAt  int64
	UpdatedAt  int64
	Options    []Option
}

type Polls struct {
	Meta  *Meta      `json:"meta"`
	Data  []PollData `json:"data"`
	Links *Links     `json:"links"`
}

type RewardData struct {
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
	Image                   *Image
	Active                  bool
	StartsAt                int64
	CreatedAt               int64
	UpdatedAt               int64
}

type Rewards struct {
	Meta  *Meta        `json:"meta"`
	Data  []RewardData `json:"data"`
	Links *Links       `json:"links"`
}

type ScheduleData struct {
	Id          int
	Name        string
	Description string
	StartsAt    int64
}

type Schedules struct {
	Meta  *Meta          `json:"meta""`
	Data  []ScheduleData `json:"data"`
	Links *Links         `json:"data"`
}

type Team struct {
	Id       int
	Username string
	Slug     string
	URL      string
	Avatar   *Image
}

type Teams struct {
	Meta *Meta  `json:"meta""`
	Data []Team `json:"data"`
}

type DonationData struct {
	Id          int     `json:"id"`
	Amount      float32 `json:"amount"`
	Name        string  `json:"name"`
	Comment     string  `json:"comment"`
	CompletedAt int64   `json:"completedAt"`
	RewardId    int     `json:"rewardId"`
}

type Donation struct {
	Meta *Meta         `json:"meta"`
	Data *DonationData `json:"data"`
}

type Donations struct {
	Meta  *Meta          `json:"meta"`
	Data  []DonationData `json:"data"`
	Links *Links         `json:"links"`
}

type UserData struct {
	Id                int     `json:"id"`
	Username          string  `json:"username"`
	Slug              string  `json:"slug"`
	URL               string  `json:"url"`
	Avatar            *Image  `json:"avatar"` // image object proper?
	About             string  `json:"about"`
	TotalAmountRaised float32 `json:"totalAmountRaised"`
	Social            *Social `json:"social"`
}

type User struct {
	Meta   *Meta     `json:"meta"`
	Data   *UserData `json:"data"`
	Status string    `json:"status"`

	tc *TiltifyClient
}

type Users struct {
	Meta   *Meta      `json:"meta"`
	Data   []UserData `json:"data"`
	Status string     `json:"status"`
}

type Meta struct {
	Status int `json:"status"`
}
