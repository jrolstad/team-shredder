package services

type SecretService interface {
	GetValue(name string) (string, error)
}

func NewSecretService() SecretService {
	return &EnvironmentSecretService{}
}
