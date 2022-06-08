package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

type IdToken struct {
	IdToken string `json:"idToken"`
}

func HandleIdToken(client *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		var idToken IdToken
		if err := json.NewDecoder(r.Body).Decode(&idToken); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		const expiresIn = time.Hour * 24 * 7
		sessionCookie, err := client.SessionCookie(ctx, idToken.IdToken, expiresIn)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionCookie,
			MaxAge:   int(expiresIn.Seconds()),
			HttpOnly: true,
			Path:     "/",
		})
	}
}

func HandleWhoami(client *auth.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := client.VerifySessionCookie(ctx, sessionCookie.Value)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, _ := client.GetUser(ctx, token.UID)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("You are %v", user.DisplayName)))
	}
}

func main() {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		log.Fatal(err)
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/id-token", HandleIdToken(authClient))
	mux.HandleFunc("/whoami", HandleWhoami(authClient))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	})

	log.Println("Listening on http://localhost:8080...")
	if err := http.ListenAndServe(":8080", corsHandler.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
