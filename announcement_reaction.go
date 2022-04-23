package mastodontype

// AnnouncementReaction  Represents an emoji reaction to an Announcement.
// @doc https://docs.joinmastodon.org/entities/announcementreaction/
type AnnouncementReaction struct {
	// Name The emoji used for the reaction. Either a unicode emoji, or a custom emoji's shortcode.
	Name string `json:"name"`
	// Count The total number of users who have added this reaction.
	Count int `json:"count"`
	// Me Whether the authorized user has added this reaction to the announcement.
	Me bool `json:"me"`
	// URL A link to the custom emoji.
	URL *string `json:"url,omitempty"`
	// StaticURL A link to a non-animated version of the custom emoji.
	StaticURL *string `json:"static_url,omitempty"`
}
