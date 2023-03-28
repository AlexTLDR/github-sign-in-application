package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// type Config struct {
// 	ClientID     string
// 	ClientSecret string
// 	Endpoint     Endpoint
// 	RedirectURL  string
// 	Scopes       []string
// }

var oauthConf = &oauth2.Config{
	ClientID:     goDotEnvVariable("client ID"),
	ClientSecret: goDotEnvVariable("secret"),
	Scopes:       []string{"repo", "user"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/o/oauth2/auth",
		TokenURL: "https://provider.com/o/oauth2/token",
	},
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/callback", githubCallbackHandler)

}
