package processors

import (
	"context"
	"errors"
	"fmt"
	"github.com/ctreminiom/go-atlassian/v2/confluence"
	models2 "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/jrolstad/team-shredder/internal/pkg/core"
	"github.com/jrolstad/team-shredder/internal/pkg/models"
	"github.com/jrolstad/team-shredder/internal/pkg/services"
	"strconv"
	"strings"
	"time"
)

type ConfluenceActionProcessor struct {
	SecretService services.SecretService
}

func (p *ConfluenceActionProcessor) Process(toProcess *models.DataActionConfiguration) (*models.DataActionResult, error) {
	client, err := p.createConfluenceClient(toProcess)
	if err != nil {
		return createErrorResult(err), err
	}
	contentToActOn, err := p.queryContent(toProcess, client)
	if err != nil {
		return createErrorResult(err), err
	}

	if strings.EqualFold(toProcess.Action, "delete") {
		return p.deleteContent(contentToActOn, toProcess, client)
	} else if strings.EqualFold(toProcess.Action, "archive") {
		return p.archiveContent(contentToActOn, toProcess, client)
	} else {
		err = errors.New("unsupported action: " + toProcess.Action)
		return createErrorResult(err), err
	}

}

func (p *ConfluenceActionProcessor) createConfluenceClient(toProcess *models.DataActionConfiguration) (*confluence.Client, error) {
	instance, err := confluence.New(nil, toProcess.Site)
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

func (p *ConfluenceActionProcessor) queryContent(toProcess *models.DataActionConfiguration, client *confluence.Client) ([]string, error) {
	//TODO: Implement paging

	searchResults, response, err := client.Search.Content(context.Background(),
		toProcess.Query,
		&models2.SearchContentOptions{},
	)
	if response != nil && response.StatusCode != 200 {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	items := make([]string, len(searchResults.Results))
	for i, item := range searchResults.Results {
		items[i] = item.Content.ID
	}
	return items, err
}

func (p *ConfluenceActionProcessor) deleteContent(toDelete []string, toProcess *models.DataActionConfiguration, client *confluence.Client) (*models.DataActionResult, error) {
	deleteErrors := map[string]error{}

	for _, item := range toDelete {
		_, err := client.Content.Delete(context.Background(), item, "")
		if err != nil {
			deleteErrors[item] = err
		} else {
			fmt.Printf("Deleting %s\n", item)
		}
	}

	return &models.DataActionResult{
		OrganizationId:      toProcess.OrganizationId,
		Site:                toProcess.Site,
		AppType:             toProcess.AppType,
		Action:              toProcess.Action,
		StartedAt:           time.Now(),
		EndedAt:             time.Now(),
		AffectedObjectCount: len(toDelete),
		FailureCount:        len(deleteErrors),
		Failures:            core.FlattenErrors(deleteErrors),
	}, nil
}

func (p *ConfluenceActionProcessor) archiveContent(toArchive []string, toProcess *models.DataActionConfiguration, client *confluence.Client) (*models.DataActionResult, error) {
	archiveErrors := map[string]error{}

	for _, item := range toArchive {
		contentId, err := strconv.Atoi(item)
		_, _, err = client.Content.Archive(context.Background(), &models2.ContentArchivePayloadScheme{
			Pages: []*models2.ContentArchiveIDPayloadScheme{
				{
					ID: contentId,
				},
			},
		})
		if err != nil {
			archiveErrors[item] = err
		} else {
			fmt.Printf("Archving %s\n", item)
		}
	}

	return &models.DataActionResult{
		OrganizationId:      toProcess.OrganizationId,
		Site:                toProcess.Site,
		AppType:             toProcess.AppType,
		Action:              toProcess.Action,
		StartedAt:           time.Now(),
		EndedAt:             time.Now(),
		AffectedObjectCount: len(toArchive),
		FailureCount:        len(archiveErrors),
		Failures:            core.FlattenErrors(archiveErrors),
	}, nil
}
