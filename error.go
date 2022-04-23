package mastodontype

// Error Represents an error message.
// @doc https://docs.joinmastodon.org/entities/error/
type Error struct {
	// The error message.
	Error string `json:"error"`
	// A longer description of the error, mainly provided with the OAuth API.
	ErrorDescription string `json:"error_description"`
}
