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

// NewSubscription object from Instagram
type NewSubscription struct {
	Meta *Meta
	Data []*NewSubscriptionData
}

// Meta object from Instagram
type Meta struct {
	Code int64
}

// NewSubscriptionData object from Instagram
type NewSubscriptionData struct {
	ID          string
	Type        string
	Object      string
	ObjectID    string `json:"object_id"`
	Aspect      string
	CallbackURL string `json:"callback_url"`
}
