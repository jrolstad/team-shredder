package processors

import (
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/services"
	"time"
)

type ConfluenceActionProcessor struct {
	SecretService services.SecretService
}

func (p *ConfluenceActionProcessor) Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error) {
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
