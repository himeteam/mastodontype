package mastodontype

// Mention Represents a mention of a user within the content of a status.
// @doc https://docs.joinmastodon.org/entities/mention/
type Mention struct {
	// ID The account id of the mentioned user.
	// (cast from an integer, but not guaranteed to be a number)
	ID string `json:"id"`
	// UserName The username of the mentioned user.
	UserName string `json:"username"`
	// Acct The webfinger acct: URI of the mentioned user. Equivalent to username for local users, or username@domain for remote users.
	Acct string `json:"acct"`
	// URL The location of the mentioned user's profile.
	URL string `json:"url"`
}
