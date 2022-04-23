package mastodontype

// AdminReport Admin::Report Admin-level information about a filed report.
// @doc https://docs.joinmastodon.org/entities/admin-report/
type AdminReport struct {
	// The ID of the report in the database.
	ID string `json:"id"`
	// The action taken to resolve this report.
	ActionTaken bool `json:"action_taken"`
	// An optional reason for reporting.
	Comment string `json:"comment"`
	// The time the report was filed.
	CreatedAt string `json:"created_at"`
	// The time of last action on this report.
	UpdatedAt string `json:"updated_at"`
	// The account which filed the report.
	Account Account `json:"account"`
	// The account being reported.
	TargetAccount Account `json:"target_account"`
	// The account of the moderator assigned to this report.
	AssignedAccount Account `json:"assigned_account"`
	// The action taken by the moderator who handled the report.
	ActionTakenByAccount string `json:"action_taken_by_account"`
	// Statuses attached to the report, for context.
	Statues []Status `json:"statuses"`
}
