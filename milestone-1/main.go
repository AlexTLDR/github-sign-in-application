package main

import "golang.org/x/oauth2"

// type Config struct {
// 	ClientID     string
// 	ClientSecret string
// 	Endpoint     Endpoint
// 	RedirectURL  string
// 	Scopes       []string
// }

var oauthConf = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	Scopes:       []string{"SCOPE1", "SCOPE2"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/o/oauth2/auth",
		TokenURL: "https://provider.com/o/oauth2/token",
	},
}

func main() {

}
