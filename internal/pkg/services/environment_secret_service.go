package services

import "os"

type EnvironmentSecretService struct {
}

func (s *EnvironmentSecretService) GetValue(name string) (string, error) {
	return os.Getenv(name), nil
}
