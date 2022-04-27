package mastodontype

// ScheduledStatus Represents a status that will be published at a future scheduled date.
// @doc https://docs.joinmastodon.org/entities/scheduledstatus/
type ScheduledStatus struct {
	// ID of the scheduled status in the database.
	ID string `json:"id"`
	// ID of the status in the database.
	// ISO 8601 Datetime
	ScheduledAt string `json:"scheduled_at"`
	params      map[string]any
}
