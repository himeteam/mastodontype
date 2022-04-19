package mastodontype

// AdminAccount Admin::Account
// Admin-level information about a given account.
// @doc https://docs.joinmastodon.org/entities/admin-account/
type AdminAccount struct {
	// ID The ID of the account in the database.
	ID string `json:"id"`
	// UserName The username of the account.
	UserName string `json:"username"`
	// Domain The domain of the account.
	Domain string `json:"domain"`
	// CreatedAt When the account was first discovered.
	CreatedAt string `json:"created_at"`
	// Email The email address associated with the account.
	Email string `json:"email"`
	// IP  The IP address last used to login to this account.
	IP string `json:"ip"`
	// Locale The locale of the account.
	// ISO 639 Part 1 two-letter language code
	Locale string `json:"locale"`
	// InviteRequest  Invite request text ???
	InviteRequest string `json:"invite_request"`
	// Role The current role of the account.
	Role string `json:"role"`
	// Confirmed Whether the account has confirmed their email address.
	Confirmed bool `json:"confirmed"`
	// Approved Whether the account is currently approved.
	Approved bool `json:"approved"`
	//Disabled  Whether the account is currently disabled.
	Disabled bool `json:"disabled"`
	// Silenced Whether the account is currently silenced
	Silenced bool `json:"silenced"`
	// Suspended Whether the account is currently suspended
	Suspended bool `json:"suspended"`
	// Account  User-level information about the account.
	Account Account `json:"account"`
	// CreatedByApplicationId The ID of the application that created this account.
	CreatedByApplicationId *string `json:"created_by_application_id"`
	// InvitedByAccountId  The ID of the account that invited this user
	InvitedByAccountId *string `json:"invited_by_account_id"`
}
