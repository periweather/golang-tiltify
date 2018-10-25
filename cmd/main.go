package main

import (
	"fmt"
	"net/http"
	"tiltify"
)

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}

// peri id 34660
// mavi id 34652
// displaced slug/id play-for-the-displaced/14861

func main() {

	//url := "https://tiltify.com/api/v3"

	client := http.DefaultClient
	tc := tiltify.TiltifyClient{Client: client, AuthorizationToken: "924f576d2ef50a6c252957a008e5bfbc2232f61c3f2354abfe33af19dc1546f3"}

	user, err := tc.GetUser(34660)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Data)
	fmt.Println(user.GetOwnedTeams())

	campaign, err := tc.GetCampaign(14861)
	if err != nil {
		panic(err)
	}
	//fmt.Println(campaign.Data.User.GetTeams())

}
