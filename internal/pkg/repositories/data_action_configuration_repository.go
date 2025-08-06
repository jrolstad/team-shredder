package repositories

import "github.com/jrolstad/team-shredder/internal/pkg/models"

type DataActionConfigurationRepository interface {
	GetOrganizations() ([]string, error)
	Get(organizationId string) ([]*models.DataActionConfiguration, error)
}

func NewDataActionConfigurationRepository() DataActionConfigurationRepository {
	return &InMemoryDataActionConfigurationRepository{}
}
