package mastodontype

const (
	// PostingVisibilityTypePublic Public post
	PostingVisibilityTypePublic = "public"
	// PostingVisibilityTypeUnlisted Unlisted post
	PostingVisibilityTypeUnlisted = "unlisted"
	// PostingVisibilityTypePrivate Followers-only post
	PostingVisibilityTypePrivate = "private"
	// PostingVisibilityTypeDirect Direct post
	PostingVisibilityTypeDirect = "direct"
	// ReadingExpandMediaTypeDefault Hide media marked as sensitive
	ReadingExpandMediaTypeDefault = "default"
	// ReadingExpandMediaTypeShowAll Always show all media by default, regardless of sensitivity
	ReadingExpandMediaTypeShowAll = "show_all"
	// ReadingExpandMediaTypeHideAll Always hide all media by default, regardless of sensitivity
	ReadingExpandMediaTypeHideAll = "hide_all"
)

type Preferences struct {
	// Default visibility for new posts. Equivalent to Source#privacy
	PostingDefaultVisibility string `json:"posting:default:visibility"`
	// Default sensitivity flag for new posts. Equivalent to Source#sensitive
	PostingDefaultSensitive bool `json:"posting:default:sensitive"`
	// default language for new posts. Equivalent to Source#language
	PostingDefaultLanguage *string `json:"posting:default:language"`
	// Whether media attachments should be automatically displayed or blurred/hidden.
	ReadingExpandMedia string `json:"reading:expand:media"`
	// Whether CWs should be expanded by default.
	ReadingExpandSpoilers bool `json:"reading:expand:spoilers"`
}
