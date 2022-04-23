package mastodontype

// Instance Represents the software instance of Mastodon running on this domain.
// @doc https://docs.joinmastodon.org/entities/instance/
type Instance struct {
	// The domain name of the instance.
	URI string `json:"uri"`
	// The title of the website.
	Title string `json:"title"`
	// Admin-defined description of the Mastodon site.
	Description string `json:"description"`
	// A shorter description defined by the admin.
	ShortDescription string `json:"short_description"`
	// An email that may be contacted for any inquiries.
	Email string `json:"email"`
	// The version of Mastodon installed on the instance.
	Version string `json:"version"`
	// Primary languages of the website and its staff.
	// ISO 639 Part 1-5 language codes
	Languages []string `json:"language"`
	// Registrations Whether registrations are enabled.
	Registrations bool `json:"registrations"`
	// ApprovalRequired Whether registrations require moderator approval.
	ApprovalRequired bool `json:"approval_required"`
	// InvitesEnabled Whether invites are enabled.
	InvitesEnabled bool `json:"invites_enabled"`
	//  URLs of interest for clients apps.
	URLs InstanceURLs `json:"urls"`
	// Statistics about how much information the instance contains.
	Stats InstanceStats `json:"stats"`
	// Banner image for the website.
	Thumbnail *string `json:"thumbnail"`
	// A user that can be contacted, as an alternative to email.
	ContactAccount *Account `json:"contact_account"`
}

// InstanceURLs URLs of interest for clients apps.
type InstanceURLs struct {
	// Websockets address for push streaming. String (URL).
	StreamingAPI string `json:"streaming_api"`
}

// InstanceStats Statistics about how much information the instance contains.
type InstanceStats struct {
	// Users registered on this instance. Number.
	UserCount int `json:"user_count"`
	// Statuses authored by users on instance. Number.
	StatusCount int `json:"status_count"`
	// Domains federated with this instance. Number.
	DomainCount int `json:"domain_count"`
}
