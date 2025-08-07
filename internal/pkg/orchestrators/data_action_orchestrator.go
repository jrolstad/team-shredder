package orchestrators

import (
	"errors"
	"github.com/jrolstad/team-shredder/internal/pkg/core"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/processors"
	"github.com/jrolstad/team-shredder/internal/pkg/repositories"
)

func ExecuteDataActions(configurationRepository repositories.DataActionConfigurationRepository, processorFactory processors.DataActionProcessorFactory) []*models.DataActionResult {
	organizations, err := configurationRepository.GetOrganizations()
	if err != nil {
		return []*models.DataActionResult{createErrorResult(err)}
	}

	results := make(map[string][]*models.DataActionResult)
	processingErrors := make(map[string]error)
	for _, orgId := range organizations {
		result, processingError := processOrganization(configurationRepository, processorFactory, orgId)

		results[orgId] = result
		processingErrors[orgId] = processingError
	}

	return flattenResults(results)
}

func processOrganization(configurationRepository repositories.DataActionConfigurationRepository, processorFactory processors.DataActionProcessorFactory, organizationId string) ([]*models.DataActionResult, error) {
	configuredTargets, err := configurationRepository.Get(organizationId)
	if err != nil {
		return []*models.DataActionResult{createErrorResult(err)}, err
	}

	results := make(map[string]*models.DataActionResult)
	processingErrors := make(map[string]error)
	for _, configuredTarget := range configuredTargets {
		result, processingError := processTarget(processorFactory, configuredTarget)

		results[configuredTarget.Id] = result
		processingErrors[configuredTarget.Id] = processingError
	}

	return flattenResult(results), errors.Join(core.FlattenErrors(processingErrors)...)
}

func processTarget(processorFactory processors.DataActionProcessorFactory, actionConfiguration *models.DataActionConfiguration) (*models.DataActionResult, error) {
	actionProcessor, err := processorFactory.GetProcessor(actionConfiguration)
	if err != nil {
		return createErrorResult(err), err
	}

	return actionProcessor.Process(actionConfiguration)
}

func createErrorResult(err error) *models.DataActionResult {
	return &models.DataActionResult{
		FailureCount: 1,
		Failures:     []error{err},
	}
}
