package repositories

import (
	"github.com/jrolstad/team-shredder/internal/pkg/models"
)

type InMemoryDataActionConfigurationRepository struct {
}

func (r InMemoryDataActionConfigurationRepository) GetOrganizations() ([]string, error) {
	return []string{"cf35573a-88ed-4070-a8fa-edbb5d42bb55"}, nil
}

func (r InMemoryDataActionConfigurationRepository) Get(organizationId string) ([]*models.DataActionConfiguration, error) {
	return []*models.DataActionConfiguration{
		{
			Id:             "1",
			OrganizationId: organizationId,
			AppType:        "confluence",
			Action:         "purgeTrash",
			Site:           "jrolstad-sandbox-1",
			Query:          "lastModified < now(\"-5d\") AND type = page",
		},
		{
			Id:             "2",
			OrganizationId: organizationId,
			AppType:        "jira",
			Action:         "delete",
			Site:           "jrolstad-sandbox-1",
			Query:          "updated <= \"-5d\"",
		},
	}, nil
}
