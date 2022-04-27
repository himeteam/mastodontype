package mastodontype

// Results Represents the results of a search.
// @doc https://docs.joinmastodon.org/entities/results/
type Results struct {
	// Accounts which match the given query
	Accounts []Account `json:"accounts"`
	// Statuses which match the given query
	Statuses []Status `json:"statuses"`
	// Hashtags which match the given query
	Hashtags []Tag `json:"hashtags"`
}
