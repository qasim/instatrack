package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURI  = os.Getenv("REDIRECT_URI")
	upgrader     = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		// We only want this to match the true index
		return
	}

	url := fmt.Sprintf("https://api.instagram.com/oauth/authorize/?client_id=%s&redirect_uri=%s&response_type=code",
		clientID, redirectURI)
	http.Redirect(w, r, url, 301)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()

	if query.Get("error") != "" {
		// User refused to be insta-cool
		http.Redirect(w, r, "/", 301)
	}

	resp, err := http.PostForm("https://api.instagram.com/oauth/access_token",
		url.Values{
			"client_id":     {clientID},
			"client_secret": {clientSecret},
			"grant_type":    {"authorization_code"},
			"redirect_uri":  {redirectURI},
			"code":          {query.Get("code")}})
	if err != nil {
		http.Redirect(w, r, "/error", 301)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var o OAuthResponse
	err = decoder.Decode(&o)
	if err != nil {
		http.Redirect(w, r, "/error", 301)
	}

	expiryDate := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "access_token",
		Value:    o.AccessToken,
		Expires:  expiryDate,
		HttpOnly: true,
		MaxAge:   50000,
		Path:     "/"}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/track", 301)
}

func handleSubscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, r.URL.Query().Get("hub.challenge"))
}

func handleSocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		return
	case "GET":
		return
	default:
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/auth/", handleAuth)
	http.Handle("/track/", http.StripPrefix("/track/", http.FileServer(http.Dir("./www"))))

	http.HandleFunc("/socket", handleSocket)

	http.HandleFunc("/subscriptions", handleSubscriptions)

	http.HandleFunc("/", handleIndex)

	log.Println("Server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
