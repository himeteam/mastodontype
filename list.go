package mastodontype

const (
	// RepliesPolicyFollowed Show replies to any followed user
	RepliesPolicyFollowed = "followed"
	// RepliesPolicyList Show replies to members of the list
	RepliesPolicyList = "list"
	// RepliesPolicyNone Show replies to no one
	RepliesPolicyNone = "none"
)

// List Represents a list of some users that the authenticated user follows.
// @doc https://docs.joinmastodon.org/entities/list/
type List struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	RepliesPolicy string `json:"replies_policy"`
}
