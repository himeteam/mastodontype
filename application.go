package mastodontype

// Application Represents an application that interfaces with the REST API to access accounts or post statuses.
// @doc https://docs.joinmastodon.org/entities/application/
type Application struct {
	// The name of your application.
	Name string `json:"name"`
	//  The website associated with your application.
	Website *string `json:"website"`
	//  Used for Push Streaming API. Returned with POST /api/v1/apps. Equivalent to PushSubscription#server_key
	VapidKey *string `json:"vapid_key"`
	// Client ID key, to be used for obtaining OAuth tokens
	ClientID string `json:"client_id"`
	// Client secret key, to be used for obtaining OAuth tokens
	ClientSecret string `json:"client_secret"`
}
