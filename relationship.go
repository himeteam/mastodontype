package mastodontype

// Relationship Represents the relationship between accounts, such as following / blocking / muting / etc.
// @doc https://docs.joinmastodon.org/entities/relationship/
type Relationship struct {
	// The account id.
	ID string `json:"id"`
	// Are you following this user?
	Following bool `json:"following"`
	// Do you have a pending follow request for this user?
	Requested bool `json:"requested"`
	// Are you featuring this user on your profile?
	Endorsed bool `json:"endorsed"`
	// Are you followed by this user?
	FollowedBy bool `json:"followed_by"`
	// Are you muting this user?
	Muting bool `json:"muting"`
	// Are you muting notifications from this user?
	MutingNotifications bool `json:"muting_notifications"`
	// Are you receiving this user's boosts in your home timeline?
	ShowingReblogs bool `json:"showing_reblogs"`
	// Have you enabled notifications for this user?
	Notifying bool `json:"notifying"`
	//  Are you blocking this user?
	Blocking bool `json:"blocking"`
	// Are you blocking this user's domain?
	DomainBlocking bool `json:"domain_blocking"`
	// Is this user blocking you?
	BlockedBy bool `json:"blocked_by"`
	// This user's profile bio
	Note string `json:"note"`
}
