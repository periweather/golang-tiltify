package tiltify

type Avatar struct {
	Src string
	Alt string
	Width int
	Height int
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
