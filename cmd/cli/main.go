package main

import (
	"fmt"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/orchestrators"
	"github.com/jrolstad/team-shredder/internal/pkg/processors"
	"github.com/jrolstad/team-shredder/internal/pkg/repositories"
)

func main() {
	configurationRepository := repositories.NewDataActionConfigurationRepository()
	processorFactory := processors.NewDataActionProcessorFactory()
	result := orchestrators.ExecuteDataActions(configurationRepository, processorFactory)
	showResult(result)
}

func showResult(result []*models.DataActionResult) {
	if len(result) == 0 {
		fmt.Println("No results")
		return
	}

	for _, item := range result {
		fmt.Println("----------------------")
		fmt.Printf("Org Id: %v\n", item.OrganizationId)
		fmt.Printf("  Site: %v\n", item.Site)
		fmt.Printf("  App Type: %v\n", item.AppType)
		fmt.Printf("  Action: %v\n", item.Action)
		fmt.Printf("    %v => %v\n", item.StartedAt, item.EndedAt)
		fmt.Printf("    Affected Items: %v\n", item.AffectedObjectCount)
		fmt.Printf("    Failures: %v\n", item.FailureCount)
		if len(item.Failures) > 0 {
			fmt.Println(item.Failures)
		}
	}
}
