package mastodontype

// Context Represents the tree around a given status. Used for reconstructing threads of statuses.
// @doc https://docs.joinmastodon.org/entities/context/
type Context struct {
	// Parents in the thread.
	Ancestors []Status `json:"ancestors"`
	// Children in the thread.
	Descendants []Status `json:"descendants"`
}
