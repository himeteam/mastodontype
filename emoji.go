package mastodontype

// Emoji Represents a custom emoji.
// @doc https://docs.joinmastodon.org/entities/emoji/
type Emoji struct {
	Shortcode string `json:"shortcode"`
	// URL A link to the custom emoji.
	URL string `json:"url"`
	// StaticURL A link to a static copy of the custom emoji.
	StaticURL string `json:"static_url"`
	// VisibleInPicker  Whether this Emoji should be visible in the picker or unlisted.
	VisibleInPicker bool `json:"visible_in_picker"`
	// Category Used for sorting custom emoji in the picker.
	Category string `json:"category"`
}
