package mastodontype

const (
	// FilterContextHome  home timeline and lists
	FilterContextHome = "home"
	// FilterContextPublic  public timelines
	FilterContextPublic = "public"
	// FilterContextThread expanded thread of a detailed status
	FilterContextThread = "thread"
	// FilterContextNotification notifications timeline
	FilterContextNotification = "notification"
)

// Filter Represents a user-defined filter for determining which statuses should not be shown to the user.
// @doc https://docs.joinmastodon.org/entities/filter/
type Filter struct {
	// The ID of the filter in the database.
	ID string `json:"id"`
	// The text to be filtered.
	Phrase string `json:"phrase"`
	// The contexts in which the filter should be applied.
	Context []string `json:"context"`
	// When the filter should no longer be applied
	ExpiredAt *string `json:"expired_at"`
	// Should matching entities in home and notifications be dropped by the server?
	Irreversible bool `json:"irreversible"`
	// Should the filter consider word boundaries?
	WholeWord bool `json:"whole_word"`
}
