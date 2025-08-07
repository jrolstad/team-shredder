package processors

import (
	"context"
	"errors"
	"fmt"
	v3 "github.com/ctreminiom/go-atlassian/v2/jira/v3"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/services"
	"strings"
	"time"
)

type JiraActionProcessor struct {
	SecretService services.SecretService
}

func (p *JiraActionProcessor) Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error) {

	jiraClient, err := p.createJiraClient(toProcess)
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

func (p *JiraActionProcessor) createJiraClient(toProcess *models.DataActionConfiguration) (*v3.Client, error) {
	instance, err := v3.New(nil, toProcess.Site)
	if err != nil {
		return nil, err
	}

	userName, err := p.SecretService.GetValue(models.Secret_AtlassianUserNameKey)
	if err != nil {
		return nil, err
	}

	password, err := p.SecretService.GetValue(models.Secret_AtlassianApiKey)
	if err != nil {
		return nil, err
	}

	instance.Auth.SetBasicAuth(userName, password)
	return instance, nil
}

func (p *JiraActionProcessor) queryIssues(toProcess *models.DataActionConfiguration, client *v3.Client) ([]string, error) {

	searchResults, _, err := client.Issue.Search.SearchJQL(context.Background(),
		toProcess.Query,
		[]string{},
		[]string{},
		1000,
		"")
	if err != nil {
		return nil, err
	}

	issues := make([]string, len(searchResults.Issues))
	for i, issue := range searchResults.Issues {
		issues[i] = issue.Key
	}

	return issues, err
}

func (p *JiraActionProcessor) deleteIssues(toDelete []string, toProcess *models.DataActionConfiguration, client *v3.Client) (*models.DataActionResult, error) {
	for _, issue := range toDelete {
		//client.Issue.Delete(context.Background(), issue, false)
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
