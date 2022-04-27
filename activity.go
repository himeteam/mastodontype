package mastodontype

// Activity Represents a weekly bucket of instance activity.
// @doc https://docs.joinmastodon.org/entities/activity/
type Activity struct {
	// Midnight at the first day of the week.
	Week string `json:"week"`
	// created since the week began.
	Statuses string `json:"statuses"`
	// User logins since the week began.
	Logins string `json:"logins"`
	// User registrations since the week began.
	Registrations string `json:"registrations"`
}
