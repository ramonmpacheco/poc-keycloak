package main

import (
	"context"
	"encoding/json"
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
	http.HandleFunc("/auth/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "invalid state", http.StatusBadRequest)
			return
		}
		// getting the authorization token -> type bearer, it's not the same as authentication token
		// you are authorized to authenticate
		token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "change token failed", http.StatusInternalServerError)
			return
		}
		resp := struct {
			AccessToken *oauth2.Token
		}{token}
		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
