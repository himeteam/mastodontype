package mastodontype

// Poll Represents a poll attached to a status.
// @doc https://docs.joinmastodon.org/entities/poll/
type Poll struct {
	// The ID of the poll in the database.
	ID string `json:"id"`
	// When the poll ends.
	// String (ISO 8601 Datetime), or null if the poll does not end
	ExpiresAt *string `json:"expires_at"`
	// Is the poll currently expired?
	Expired bool `json:"expired"`
	// Does the poll allow multiple-choice answers?
	Multiple bool `json:"multiple"`
	// How many votes have been received.
	VotesCount int `json:"votes_count"`
	// How many unique accounts have voted on a multiple-choice poll.
	// Number, or null if multiple is false.
	VotersCount *int `json:"voters_count"`
	// When called with a user token, has the authorized user voted?
	//  Boolean, or null if no current user
	Voted *bool `json:"voted"`
	// When called with a user token, which options has the authorized user chosen? Contains an array of index values for options.
	// Array of Number, or null if no current user
	OwnVotes *[]int `json:"own_votes"`
	//  Possible answers for the poll.
	Options []PollOption `json:"options"`
	// Custom emoji to be used for rendering poll options.
	Emojis []Emoji `json:"emojis"`
	// Hide vote counts until the poll ends?
	HideTotals bool `json:"hide_totals"`
}

// PollOption Possible answers for the poll.
type PollOption struct {
	//  The text value of the poll option
	Title string `json:"title"`
	// The number of received votes for this option. Number, or null if results are not published yet.
	VotesCount int `json:"votes_count"`
}
