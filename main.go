package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofrs/uuid"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func oauth2CallbackHandler(config *oauth2.Config, authState string, c chan<- *oauth2.Token) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		if state != authState {
			http.Error(w, "Auth state does not match.", http.StatusUnauthorized)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Missing code param.", http.StatusBadRequest)
			return
		}

		t, err := config.Exchange(context.Background(), code)
		if err != nil {
			fmt.Fprintf(w, "Unable to exchange code for token: %v\n", err)
			return
		}

		fmt.Fprintf(w, "Access token: %s\n", t.AccessToken)
		fmt.Fprintf(w, "Refresh token: %s\n", t.RefreshToken)
		fmt.Fprintln(w, "You can close this window now.")

		c <- t
	})
}

func printUsage() {
	fmt.Println("usage: google-access-token <client_id> <client_secret>")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "incorrect arguments")
		printUsage()
		os.Exit(1)
	}

	port := "9889"
	redirectURL := fmt.Sprintf("http://localhost:%s/", port)

	config := &oauth2.Config{
		ClientID:     os.Args[1],
		ClientSecret: os.Args[2],
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes: []string{
			drive.DriveScope,
		},
	}

	id, _ := uuid.NewV4()
	authState := id.String()
	u := config.AuthCodeURL(authState, oauth2.AccessTypeOffline)

	c := make(chan *oauth2.Token)
	go func() {
		http.ListenAndServe(":"+port, oauth2CallbackHandler(config, authState, c))
	}()

	err := browser.OpenURL(u)
	if err != nil {
		fmt.Printf("Please access the following URL on your browser: %s\n", u)
	}

	fmt.Println("Waiting for OAuth2 calback...")

	t := <-c

	fmt.Printf("Access token: %s\n", t.AccessToken)
	fmt.Printf("Refresh token: %s\n", t.RefreshToken)
}
