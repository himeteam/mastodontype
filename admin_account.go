package mastodontype

// AdminAccount Admin::Account
// Admin-level information about a given account.
// @doc https://docs.joinmastodon.org/entities/admin-account/
type AdminAccount struct {
	// The ID of the account in the database.
	ID string `json:"id"`
	// The username of the account.
	UserName string `json:"username"`
	// The domain of the account.
	Domain string `json:"domain"`
	// When the account was first discovered.
	CreatedAt string `json:"created_at"`
	// The email address associated with the account.
	Email string `json:"email"`
	// The IP address last used to login to this account.
	IP string `json:"ip"`
	// The locale of the account.
	// ISO 639 Part 1 two-letter language code
	Locale string `json:"locale"`
	// Invite request text ???
	InviteRequest string `json:"invite_request"`
	// The current role of the account.
	Role string `json:"role"`
	// Whether the account has confirmed their email address.
	Confirmed bool `json:"confirmed"`
	// Whether the account is currently approved.
	Approved bool `json:"approved"`
	// Whether the account is currently disabled.
	Disabled bool `json:"disabled"`
	// Whether the account is currently silenced
	Silenced bool `json:"silenced"`
	// Whether the account is currently suspended
	Suspended bool `json:"suspended"`
	//  User-level information about the account.
	Account Account `json:"account"`
	// The ID of the application that created this account.
	CreatedByApplicationId *string `json:"created_by_application_id"`
	// The ID of the account that invited this user
	InvitedByAccountId *string `json:"invited_by_account_id"`
}
