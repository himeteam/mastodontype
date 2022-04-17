package mastodontype

//Field Represents a profile field as a name-value pair with optional verification.
// @doc https://docs.joinmastodon.org/entities/field/
type Field struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	VerifiedAt string `json:"verified_at"`
}
