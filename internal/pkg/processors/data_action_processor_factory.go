package processors

import (
	"errors"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/services"
	"strings"
)

type DataActionProcessorFactory struct {
	RegisteredProcessors map[string]DataActionProcessor
	SecretService        services.SecretService
}

func NewDataActionProcessorFactory(secretService services.SecretService) DataActionProcessorFactory {
	instance := DataActionProcessorFactory{
		RegisteredProcessors: make(map[string]DataActionProcessor),
	}
	instance.RegisteredProcessors["jira"] = &JiraActionProcessor{
		SecretService: secretService,
	}
	instance.RegisteredProcessors["confluence"] = &ConfluenceActionProcessor{
		SecretService: secretService,
	}

	return instance
}

func (f *DataActionProcessorFactory) GetProcessor(config *models.DataActionConfiguration) (DataActionProcessor, error) {

	processor := f.RegisteredProcessors[strings.ToLower(config.AppType)]
	if processor == nil {
		return nil, errors.New("processor not supported")
	}

	return processor, nil
}
