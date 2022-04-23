package mastodontype

// Announcement Represents an announcement set by an administrator.
// @doc https://docs.joinmastodon.org/entities/announcement/
type Announcement struct {
	// ID The announcement id.
	ID string `json:"id"`
	// Text The content of the announcement.
	Text string `json:"text"`
	// Published  Whether the announcement is currently active.
	Published bool `json:"published"`
	// AllDay Whether the announcement has a start/end time.
	AllDay bool `json:"all_day"`
	// CreatedAt When the announcement was created.
	// ISO 8601 Datetime
	CreatedAt string `json:"created_at"`
	// UpdatedAt When the announcement was last updated.
	// ISO 8601 Datetime
	UpdatedAt string `json:"updated_at"`
	// Read Whether the announcement has been read by the user.
	Read bool `json:"read"`
	// Reactions  Emoji reactions attached to the announcement.
	Reactions []AnnouncementReaction `json:"reactions"`
	// ScheduledAt When the future announcement was scheduled.
	ScheduledAt *string `json:"scheduled_at,omitempty"`
	// StartsAt When the future announcement will start.
	StartsAt *string `json:"starts_at,omitempty"`
	// EndsAt When the future announcement will end.
	EndsAt *string `json:"ends_at,omitempty"`
}
