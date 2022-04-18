package mastodontype

import "strconv"

// Activity Represents a weekly bucket of instance activity.
// @doc https://docs.joinmastodon.org/entities/activity/
type Activity struct {
	// Week Midnight at the first day of the week.
	Week string `json:"week"`
	// Statuses created since the week began.
	Statuses string `json:"statuses"`
	// Logins User logins since the week began.
	Logins string `json:"logins"`
	// Registrations User registrations since the week began.
	Registrations string `json:"registrations"`
}

// GetWeek Get the week
func (a Activity) GetWeek() int {
	if a.Week == "" {
		return 0
	}

	week, err := strconv.Atoi(a.Week)
	if err != nil {
		return 0
	}

	return week
}

// SetWeek Set the week.
func (a *Activity) SetWeek(week int) {
	a.Week = strconv.Itoa(week)
}

// GetStatuses Get the statuses.
func (a Activity) GetStatuses() int {
	if a.Statuses == "" {
		return 0
	}

	statuses, err := strconv.Atoi(a.Statuses)
	if err != nil {
		return 0
	}

	return statuses
}

// SetStatuses Set the statuses.
func (a *Activity) SetStatuses(statuses int) {
	a.Statuses = strconv.Itoa(statuses)
}

func (a Activity) GetLogins() int {
	if a.Logins == "" {
		return 0
	}

	logins, err := strconv.Atoi(a.Logins)
	if err != nil {
		return 0
	}

	return logins
}

// SetLogins Set the logins.
func (a *Activity) SetLogins(logins int) {
	a.Logins = strconv.Itoa(logins)
}

// GetRegistrations Get the registrations.
func (a Activity) GetRegistrations() int {
	if a.Registrations == "" {
		return 0
	}

	registrations, err := strconv.Atoi(a.Registrations)
	if err != nil {
		return 0
	}

	return registrations
}

// SetRegistrations Set the registrations.
func (a *Activity) SetRegistrations(registrations int) {
	a.Registrations = strconv.Itoa(registrations)
}
