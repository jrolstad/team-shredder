package orchestrators

import "github.com/jrolstad/team-shredder/internal/pkg/models"

func flattenResults(toMap map[string][]*models.DataActionResult) []*models.DataActionResult {
	result := make([]*models.DataActionResult, 0)

	for _, value := range toMap {
		result = append(result, value...)
	}

	return result
}

func flattenResult(toMap map[string]*models.DataActionResult) []*models.DataActionResult {
	result := make([]*models.DataActionResult, 0)

	for _, value := range toMap {
		result = append(result, value)
	}

	return result
}
