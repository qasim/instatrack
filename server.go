package main

import (
	"encoding/json"
	"fmt"
	"io"
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
)

func handleInstagram(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
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
		return
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
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var o OAuthResponse
	err = decoder.Decode(&o)
	if err != nil {
		http.Redirect(w, r, "/error", 301)
		return
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

func handleMedia(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method),
			http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("access_token")

	if err != nil {
		http.Redirect(w, r, "/", 301)
		return
	}

	query := r.URL.Query()
	tag := query.Get("tag")
	if len(tag) == 0 {
		http.Redirect(w, r, "/track", 301)
		return
	}

	minTagID := query.Get("min_tag_id")

	url := fmt.Sprintf("https://api.instagram.com/v1/tags/%s/media/recent?access_token=%s&min_tag_id=%s",
		tag, cookie.Value, minTagID)
	resp, err := http.Get(url)
	if err != nil {
		http.Redirect(w, r, "/error", 301)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	io.Copy(w, resp.Body)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/instagram", handleInstagram)
	http.HandleFunc("/auth/", handleAuth)
	http.HandleFunc("/media", handleMedia)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Server running at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
