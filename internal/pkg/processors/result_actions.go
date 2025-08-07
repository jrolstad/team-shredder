package processors

import "github.com/jrolstad/team-shredder/internal/pkg/models"

func createErrorResult(err error) *models.DataActionResult {
	return &models.DataActionResult{
		FailureCount: 1,
		Failures:     []error{err},
	}
}
