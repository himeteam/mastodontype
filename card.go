package mastodontype

const (
	// CardTypeLink Link OEmbed
	CardTypeLink = "link"
	// CardTypePhoto Photo OEmbed
	CardTypePhoto = "photo"
	// CardTypeVideo Video OEmbed
	CardTypeVideo = "video"
	// CardTypeRich iframe OEmbed. Not currently accepted, so won't show up in practice.
	CardTypeRich = "rich"
)

// Card Represents a rich preview card that is generated using OpenGraph tags from a URL.
// @doc https://docs.joinmastodon.org/entities/card/
type Card struct {
	// Location of linked resource.
	URL string `json:"url"`
	// of linked resource.
	Title string `json:"title"`
	// of preview.
	Description string `json:"description"`
	// of the preview card.
	Type string `json:"type"`
	// of the original resource.
	AuthorName string `json:"author_name"`
	// A link to the author of the original resource.
	AuthorURL string `json:"author_url"`
	// The provider of the original resource.
	ProviderName string `json:"provider_name"`
	// A link to the provider of the original resource.
	ProviderURL string `json:"provider_url"`
	// to be used for generating the preview card.
	HTML string `json:"html"`
	// of preview, in pixels.
	Width int `json:"width"`
	// of preview, in pixels.
	Height int `json:"height"`
	// Preview thumbnail.
	Image string `json:"image"`
	// Used for photo embeds, instead of custom html.
	EmbedURL string `json:"embed_url"`
	// BlurHash  A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	BlurHash string `json:"blurhash"`
}
