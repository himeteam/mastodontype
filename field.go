package mastodontype

//Field Represents a profile field as a name-value pair with optional verification.
// @doc https://docs.joinmastodon.org/entities/field/
type Field struct {
	// The key of a given field's key-value pair.
	Name string `json:"name"`
	// The value associated with the name key.
	// HTML
	Value string `json:"value"`
	// Timestamp of when the server verified a URL value for a rel="me” link.
	// String (ISO 8601 Datetime) if value is a verified URL. Otherwise, null
	VerifiedAt *string `json:"verified_at"`
}
