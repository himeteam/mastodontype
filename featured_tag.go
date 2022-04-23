package mastodontype

// FeaturedTag Represents a hashtag that is featured on a profile.
// @doc https://docs.joinmastodon.org/entities/featuredtag/
type FeaturedTag struct {
	// The internal ID of the featured tag in the database.
	ID string `json:"id"`
	// The name of the hashtag being featured.
	Name string `json:"name"`
	// A link to all statuses by a user that contain this hashtag.
	URL string `json:"url"`
	// The number of authored statuses containing this hashtag.
	StatusesCount int `json:"statuses_count"`
	// The timestamp of the last authored status containing this hashtag.
	// ISO 8601 Datetime
	LastStatusAt string `json:"last_status_at"`
}
