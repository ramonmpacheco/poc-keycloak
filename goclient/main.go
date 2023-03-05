package main

import (
	"context"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientId     = "myclient"
	clientSecret = "wtAnzjcMr1PTUe120SBPwgkf9y4BfhJd"
	redirectURL  = "http://localhost:8081/auth/callback"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/realms/myrealm")
	if err != nil {
		log.Fatal(err)
	}
	config := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes: []string{
			oidc.ScopeOpenID, "profile", "email", "roles",
		},
	}
	state := "123"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
