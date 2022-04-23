package mastodontype

// IdentityProof Represents a proof from an external identity provider.
// @doc https://docs.joinmastodon.org/entities/identityproof/
type IdentityProof struct {
	// The name of the identity provider.
	Provider string `json:"provider"`
	// The account owner's username on the identity provider's service.
	ProviderUsername string `json:"provider_username"`
	// The account owner's profile URL on the identity provider.
	ProfileURL string `json:"profile_url"`
	// A link to a statement of identity proof, hosted by the identity provider.
	ProofURL string `json:"proof_url"`
	// When the identity proof was last updated.
	// ISO 8601 Datetime
	UpdatedAt string `json:"updated_at"`
}
