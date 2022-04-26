package mastodontype

const (
	// NotificationTypeFollow Someone followed you
	NotificationTypeFollow = "follow"
	// NotificationTypeFollowRequest Someone requested to follow you
	NotificationTypeFollowRequest = "follow_request"
	// NotificationTypeMention Someone mentioned you in their status
	NotificationTypeMention = "mention"
	// NotificationTypeReblog Someone boosted one of your statuses
	NotificationTypeReblog = "reblog"
	// NotificationTypeFavourite Someone favourited one of your statuses
	NotificationTypeFavourite = "favourite"
	// NotificationTypePoll A poll you have voted in or created has ended
	NotificationTypePoll = "poll"
	// NotificationTypeStatus Someone you enabled notifications for has posted a status
	NotificationTypeStatus = "status"
)

// Notification Represents a notification of an event relevant to the user.
type Notification struct {
	// The id of the notification in the database.
	ID string `json:"id"`
	// The type of event that resulted in the notification.
	Type string `json:"type"`
	// The timestamp of the notification.
	// ISO 8601 Datetime
	CreatedAt string `json:"created_at"`
	// The account that performed the action that generated the notification.
	Account Account `json:"account"`
	// Status that was the object of the notification, e.g. in mentions, reblogs, favourites, or polls.
	Status *Status `json:"status"`
}
