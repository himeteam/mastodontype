package mastodontype

// History Represents daily usage history of a hashtag.
// @doc https://docs.joinmastodon.org/entities/history/
type History struct {
	// UNIX timestamp on midnight of the given day.
	// UNIX timestamp
	Day string `json:"day"`
	// the counted usage of the tag within that day.
	// cast from an integer
	Uses string `json:"uses"`
	// the total of accounts using the tag within that day.
	// cast from an integer
	Accounts string `json:"accounts"`
}
