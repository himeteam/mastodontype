package mastodontype

// Account Represents a user of Mastodon and their associated profile.
// @doc https://docs.joinmastodon.org/entities/account/
type Account struct {
	// ID The account idheader
	ID string `json:"id"`
	// UserName The username of the account, not including domain.
	UserName string `json:"username"`
	// Acct The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	Acct string `json:"acct"`
	// URL The location of the user's profile page.
	URL string `json:"url"`
	// DisplayName The profile's display name.
	DisplayName string `json:"display_name"`
	// Note The profile's bio / description(HTML).
	Note string `json:"note"`
	// Avatar An image icon that is shown next to statuses and in the profile(URL).
	Avatar string `json:"avatar"`
	// AvatarStatic A static version of the Avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	AvatarStatic string `json:"avatar_static"`
	// Header An image banner that is shown above the profile and in profile cards(URL).
	Header string `json:"header"`
	// HeaderStatic  A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF(URL).
	HeaderStatic string `json:"header_static"`
	// Locked Whether the account manually approves follow requests.
	Locked bool `json:"locked"`
	// Emojis Custom emoji entities to be used when rendering the profile. If none, an empty array will be returned.
	Emojis []Emoji `json:"emojis"`
	// Discoverable Whether the account has opted into discovery features such as the profile directory.
	Discoverable bool `json:"discoverable"`
	// CreatedAt When the account was created.
	CreatedAt string `json:"created_at"`
	// LastStatusAt When the most recent status was posted.
	LastStatusAt string `json:"last_status_at"`
	// StatusesCount How many statuses are attached to this account.
	StatusesCount int `json:"statuses_count"`
	// FollowersCount The reported followers of this profile.
	FollowersCount int `json:"followers_count"`
	// FollowingCount The reported follows of this profile.
	FollowingCount int `json:"following_count"`
	// Moved Indicates that the profile is currently inactive and that its user has moved to a new account.
	Moved *Account `json:"moved"`
	// Fields  Additional metadata attached to a profile as name-value pairs.
	Fields []Field `json:"fields"`
	// Bot A presentational flag. Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Bot bool `json:"bot"`
	// Suspended  An extra entity returned when an account is suspended.
	Suspended bool `json:"suspended"`
	// MuteExpiresAt When a timed mute will expire, if applicable.
	MuteExpiresAt string `json:"mute_expires_at"`
}
