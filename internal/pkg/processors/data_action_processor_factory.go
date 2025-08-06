package processors

import (
	"errors"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"strings"
)

type DataActionProcessorFactory struct {
	RegisteredProcessors map[string]DataActionProcessor
}

func NewDataActionProcessorFactory() DataActionProcessorFactory {
	instance := DataActionProcessorFactory{
		RegisteredProcessors: make(map[string]DataActionProcessor),
	}
	instance.RegisteredProcessors["jira"] = &JiraActionProcessor{}
	instance.RegisteredProcessors["confluence"] = &ConfluenceActionProcessor{}

	return instance
}

func (f *DataActionProcessorFactory) GetProcessor(config *models.DataActionConfiguration) (DataActionProcessor, error) {

	processor := f.RegisteredProcessors[strings.ToLower(config.AppType)]
	if processor == nil {
		return nil, errors.New("processor not supported")
	}

	return processor, nil
}
