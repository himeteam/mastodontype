package mastodontype

// AdminReport Admin::Report Admin-level information about a filed report.
// @doc https://docs.joinmastodon.org/entities/admin-report/
type AdminReport struct {
	// ID The ID of the report in the database.
	ID string `json:"id"`
	// ActionTaken The action taken to resolve this report.
	ActionTaken bool `json:"action_taken"`
	// Comment  An optional reason for reporting.
	Comment string `json:"comment"`
	// CreatedAt The time the report was filed.
	CreatedAt string `json:"created_at"`
	// UpdatedAt The time of last action on this report.
	UpdatedAt string `json:"updated_at"`
	// Account The account which filed the report.
	Account Account `json:"account"`
	// TargetAccount The account being reported.
	TargetAccount Account `json:"target_account"`
	// AssignedAccount The account of the moderator assigned to this report.
	AssignedAccount Account `json:"assigned_account"`
	// ActionTakenByAccount The action taken by the moderator who handled the report.
	ActionTakenByAccount string   `json:"action_taken_by_account"`
	Statues              []Status `json:"statuses"`
}
