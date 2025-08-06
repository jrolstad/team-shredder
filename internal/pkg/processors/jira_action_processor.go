package processors

import (
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"time"
)

type JiraActionProcessor struct {
}

func (p *JiraActionProcessor) Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error) {
	return &models.DataActionResult{
		OrganizationId:      toProcess.OrganizationId,
		Site:                toProcess.Site,
		AppType:             toProcess.AppType,
		Action:              toProcess.Action,
		StartedAt:           time.Now(),
		EndedAt:             time.Now(),
		AffectedObjectCount: 1,
		FailureCount:        0,
		Failures:            make([]error, 0),
	}, nil
}
