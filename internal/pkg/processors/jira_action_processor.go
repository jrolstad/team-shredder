package processors

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"strings"
	"time"
)

type JiraActionProcessor struct {
}

func (p *JiraActionProcessor) Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error) {

	jiraClient, _ := jira.NewClient(nil, p.createBaseUrl(toProcess))
	issuesToActOn, err := p.queryIssues(toProcess, jiraClient)
	if err != nil {
		return createErrorResult(err), err
	}

	if strings.EqualFold(toProcess.Action, "delete") {
		return p.deleteIssues(issuesToActOn, toProcess, jiraClient)
	} else {
		err = errors.New("unsupported action: " + toProcess.Action)
		return createErrorResult(err), err
	}
}

func (p *JiraActionProcessor) createBaseUrl(toProcess *models.DataActionConfiguration) string {
	return "https://" + toProcess.Site
}

func (p *JiraActionProcessor) queryIssues(toProcess *models.DataActionConfiguration, client *jira.Client) ([]string, error) {
	options := &jira.SearchOptions{}
	searchResults, response, err := client.Issue.Search(toProcess.Query, options)
	if err != nil {
		return make([]string, 0), err
	}
	if response.StatusCode != 200 {
		return make([]string, 0), errors.New(response.Status)
	}

	issues := make([]string, len(searchResults))
	for i, issue := range searchResults {
		issues[i] = issue.Key
	}

	return issues, err
}

func (p *JiraActionProcessor) deleteIssues(toDelete []string, toProcess *models.DataActionConfiguration, client *jira.Client) (*models.DataActionResult, error) {
	for _, issue := range toDelete {
		fmt.Printf("Deleting %s", issue)
	}

	return &models.DataActionResult{
		OrganizationId:      toProcess.OrganizationId,
		Site:                toProcess.Site,
		AppType:             toProcess.AppType,
		Action:              toProcess.Action,
		StartedAt:           time.Now(),
		EndedAt:             time.Now(),
		AffectedObjectCount: len(toDelete),
		FailureCount:        0,
		Failures:            make([]error, 0),
	}, nil
}
