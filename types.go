package main

// User object from Instagram
type User struct {
	ID             string
	Username       string
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string
	Website        string
	Counts         *UserCounts
}

// UserCounts object from Instagram
type UserCounts struct {
	Media      int64
	Follows    int64
	FollowedBy int64 `json:"followed_by"`
}

// UserPosition object from Instagram
type UserPosition struct {
	User     *User
	Position *Position
}

// Position object from Instagram
type Position struct {
	X float64
	Y float64
}

// OAuthResponse object from Instagram
type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	User        *User
}
