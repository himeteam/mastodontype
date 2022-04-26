package mastodontype

// Marker Represents the last read position within a user's timelines.
// @doc https://docs.joinmastodon.org/entities/marker/
type Marker struct {
	// Information about the user's position in the home timeline.
	Home TimelineInfo `json:"home"`
	//  Information about the user's position in their notifications.
	Notifications TimelineInfo `json:"notifications"`
}

type TimelineInfo struct {
	// The ID of the most recently viewed entity.
	LastReadID string `json:"last_read_id"`
	// Used for locking to prevent write conflicts.
	Version int `json:"version"`
	// The timestamp of when the marker was set.
	UpdatedAt string `json:"updated_at"`
}
