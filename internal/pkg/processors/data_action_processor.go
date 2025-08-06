package processors

import "github.com/jrolstad/team-shredder/internal/pkg/models"

type DataActionProcessor interface {
	Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error)
}
