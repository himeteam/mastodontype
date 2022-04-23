package mastodontype

// Account Represents a user of Mastodon and their associated profile.
// @doc https://docs.joinmastodon.org/entities/account/
type Account struct {
	// The account idheader
	ID string `json:"id"`
	// The username of the account, not including domain.
	UserName string `json:"username"`
	// The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	Acct string `json:"acct"`
	// The location of the user's profile page.
	URL string `json:"url"`
	// The profile's display name.
	DisplayName string `json:"display_name"`
	// The profile's bio / description(HTML).
	Note string `json:"note"`
	// An image icon that is shown next to statuses and in the profile(URL).
	Avatar string `json:"avatar"`
	// A static version of the Avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	AvatarStatic string `json:"avatar_static"`
	// An image banner that is shown above the profile and in profile cards(URL).
	Header string `json:"header"`
	// A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF(URL).
	HeaderStatic string `json:"header_static"`
	// Whether the account manually approves follow requests.
	Locked bool `json:"locked"`
	// Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	Emojis []Emoji `json:"emojis"`
	// Whether the account has opted into discovery features such as the profile directory.
	Discoverable bool `json:"discoverable"`
	// When the account was created.
	CreatedAt string `json:"created_at"`
	// When the most recent status was posted.
	LastStatusAt string `json:"last_status_at"`
	// How many statuses are attached to this account.
	StatusesCount int `json:"statuses_count"`
	// The reported followers of this profile.
	FollowersCount int `json:"followers_count"`
	// FollowingCount The reported follows of this profile.
	FollowingCount int `json:"following_count"`
	// Indicates that the profile is currently inactive and that its user has moved to a new account.
	Moved *Account `json:"moved"`
	//  Additional metadata attached to a profile as name-value pairs.
	Fields []Field `json:"fields"`
	// A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Bot bool `json:"bot"`
	//  An extra entity returned when an account is suspended.
	Suspended bool `json:"suspended"`
	// When a timed mute will expire, if applicable.
	MuteExpiresAt string `json:"mute_expires_at"`
}
