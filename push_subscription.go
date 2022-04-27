package mastodontype

// PushSubscription Represents a subscription to the push streaming server.
// @doc https://docs.joinmastodon.org/entities/pushsubscription/
type PushSubscription struct {
	// The id of the push subscription in the database.
	ID string `json:"id"`
	// Where push alerts will be sent to.
	// URL
	Endpoint string `json:"endpoint"`
	// The streaming server's VAPID key.
	ServerKey string `json:"server_key"`
	// Which alerts should be delivered to the endpoint
	Alerts map[string]any `json:"alerts"`
}
