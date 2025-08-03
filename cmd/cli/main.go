package main

import (
	"fmt"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/orchestrators"
)

func main() {
	result := orchestrators.ExecuteDataActions()
	showResult(result)
}

func showResult(result []*models.DataActionResult) {
	if len(result) == 0 {
		fmt.Println("No results")
		return
	}

	for _, item := range result {
		fmt.Println("----------------------")
		fmt.Printf("Org Id %v\n", item.OrganizationId)
		fmt.Printf("    %v => %v", item.StartedAt, item.EndedAt)
		fmt.Printf("    Affected Items: %v", item.AffectedObjectCount)
		fmt.Printf("    Failures: %v", item.FailureCount)
		if len(item.Failures) > 0 {
			fmt.Println(item.Failures)
		}
	}
}
