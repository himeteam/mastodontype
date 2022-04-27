package mastodontype

// Token Represents an OAuth token used for authenticating with the API and performing actions.
// @doc https://docs.joinmastodon.org/entities/token/
type Token struct {
	// An OAuth token to be used for authorization.
	AccessToken string `json:"access_token"`
	// The OAuth token type. Mastodon uses Bearer tokens.
	TokenType string `json:"token_type"`
	// The OAuth scopes granted by this token, space-separated.
	Scope string `json:"scope"`
	// When the token was generated.
	// UNIX Timestamp
	CreatedAt int64 `json:"created_at"`
}
