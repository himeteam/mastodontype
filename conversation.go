package mastodontype

// Conversation Represents a conversation with "direct message" visibility.
// @doc https://docs.joinmastodon.org/entities/conversation/
type Conversation struct {
	// Local database ID of the conversation.
	ID string `json:"id"`
	// Participants in the conversation.
	Accounts []Account `json:"accounts"`
	// Is the conversation currently marked as unread?
	Unread bool `json:"unread"`
	// The last status in the conversation, to be used for optional display.
	LastStatus *Status `json:"last_status"`
}
