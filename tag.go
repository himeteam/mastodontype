package mastodontype

// Tag Represents a hashtag used within the content of a status.
// @doc https://docs.joinmastodon.org/entities/tag/
type Tag struct {
	// Name The value of the hashtag after the # sign.
	Name string `json:"name"`
	// URL A link to the hashtag on the instance.
	URL     string `json:"url"`
	History []History
}
