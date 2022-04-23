package mastodontype

// Tag Represents a hashtag used within the content of a status.
// @doc https://docs.joinmastodon.org/entities/tag/
type Tag struct {
	// The value of the hashtag after the # sign.
	Name string `json:"name"`
	// A link to the hashtag on the instance.
	URL string `json:"url"`
	// Usage statistics for given days.
	History []History
}
