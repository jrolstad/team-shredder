package models

import "time"

type DataActionResult struct {
	OrganizationId      string
	StartedAt           time.Time
	EndedAt             time.Time
	AffectedObjectCount int
	FailureCount        int
	Failures            []error
}
