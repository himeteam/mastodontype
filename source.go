package mastodontype

// Source Represents display or publishing preferences of user's own account. Returned as an additional entity when verifying and updated credentials, as an attribute of Account.
// @doc https://docs.joinmastodon.org/entities/source/
type Source struct {
	// Profile bio.
	Note string `json:"note"`
	// Metadata about the account.
	Fields []Field `json:"fields"`
	// The default post privacy to be used for new statuses.
	Privacy string `json:"privacy"`
	// Whether new statuses should be marked sensitive by default.
	Sensitive bool `json:"sensitive"`
	// The default posting language for new statuses.
	// ISO 639-1 language two-letter code
	Language string `json:"language"`
	// The number of pending follow requests.
	FollowRequestsCount int `json:"follow_requests_count"`
}
