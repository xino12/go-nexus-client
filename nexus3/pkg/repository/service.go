package repository

import (
	"fmt"
	"net/http"

	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
)

const (
	repositoryAPIEndpoint = client.BasePath + "v1/repositories"
)

type RepositoryService struct {
	client *client.Client

	// API Services
	Apt       *RepositoryAptService
	Bower     *RepositoryBowerService
	Cocoapods *RepositoryCocoapodsService
	Conan     *RepositoryConanService
	Conda     *RepositoryCondaService
	Legacy    *RepositoryLegacyService
}

func NewRepositoryService(c *client.Client) *RepositoryService {
	return &RepositoryService{
		client: c,

		Apt:       NewRepositoryAptService(c),
		Bower:     NewRepositoryBowerService(c),
		Cocoapods: NewRepositoryCocoapodsService(c),
		Conan:     NewRepositoryConanService(c),
		Conda:     NewRepositoryCondaService(c),
		Legacy:    NewRepositoryLegacyService(c),
	}
}

func deleteRepository(client *client.Client, id string) error {
	body, resp, err := client.Delete(fmt.Sprintf("%s/%s", repositoryAPIEndpoint, id))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("could not delete repository '%s': HTTP: %d, %s", id, resp.StatusCode, string(body))
	}
	return nil
}
