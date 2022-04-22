package mastodontype

const (
	// StatusVisibilityPublic Visible to everyone, shown in public timelines.
	StatusVisibilityPublic = "public"
	// StatusVisibilityUnlisted Visible to public, but not included in public timelines.
	StatusVisibilityUnlisted = "unlisted"
	// StatusVisibilityPrivate Visible to followers only, and to any mentioned users.
	StatusVisibilityPrivate = "private"
	// StatusVisibilityDirect Visible only to mentioned users.
	StatusVisibilityDirect = "direct"
)

// Status Represents a status posted by an account.
// @doc https://docs.joinmastodon.org/entities/status/
type Status struct {
	// ID of the status in the database.
	// cast from an integer but not guaranteed to be a number
	ID string `json:"id"`
	//  URI of the status used for federation.
	URI string `json:"uri"`
	// CreatedAt The date when this status was created.
	// ISO 8601 Datetime
	CreatedAt string `json:"created_at"`
	// Account The account that authored this status.
	Account Account `json:"account"`
	// Content  HTML-encoded status content.
	Content string `json:"content"`
	// Visibility of this status.
	Visibility string `json:"visibility"`
	// Sensitive Is this status marked as sensitive content?
	Sensitive bool `json:"sensitive"`
	// SpoilerText Subject or summary line, below which status content is collapsed until expanded.
	SpoilerText string `json:"spoiler_text"`
	// MediaAttachments Media that is attached to this status.
	MediaAttachments []Attachment `json:"media_attachments"`
	// Application The application used to post this status.
	Application Application `json:"application"`
	// Mentions of users within the status content.
	Mentions []Mention `json:"mentions"`
	// Tags Hashtags used within the status content.
	Tags []Tag `json:"tags"`
	// Emojis Custom emoji to be used when rendering status content.
	Emojis []Emoji `json:"emojis"`
	// ReblogsCount How many boosts this status has received.
	ReblogsCount int `json:"reblogs_count"`
	// FavouritesCount How many favourites this status has received.
	FavouritesCount int `json:"favourites_count"`
	// RepliesCount How many replies this status has received.
	RepliesCount int `json:"replies_count"`
	// URL A link to the status's HTML representation.
	URL *string `json:"url"`
	// InReplyToID ID of the status being replied.
	InReplyToID *string `json:"in_reply_to_id"`
	// InReplyToAccountID ID of the account being replied to.
	InReplyToAccountID *string `json:"in_reply_to_account_id"`
	// Reblog The status being reblogged.
	Reblog *Status `json:"reblog"`
	// Poll The poll attached to the status.
	Poll *Poll `json:"poll"`
	// Card Preview card for links included within status content.
	Card *Card `json:"card"`
	// Language Primary language of this status.
	Language *string `json:"language"`
	// Text Plain-text source of a status. Returned instead of content when status is deleted, so the user may redraft from the source text without the client having to reverse-engineer the original text from the HTML content.
	Text *string `json:"text"`
	// Favorited Have you favourited this status?
	Favorited *bool `json:"favorited;omitempty"`
	// Reblogged Have you boosted this status?
	Reblogged *bool `json:"reblogged;omitempty"`
	// Muted Have you muted notifications for this status's conversation?
	Muted *bool `json:"muted;omitempty"`
	// Bookmarked Have you bookmarked this status?
	Bookmarked *bool `json:"bookmarked;omitempty"`
	// Pinned Have you pinned this status? Only appears if the status is pinnable.
	Pinned *bool `json:"pinned;omitempty"`
}
