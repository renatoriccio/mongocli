// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_integrations.go -package=mocks github.com/mongodb/mongocli/internal/store IntegrationCreator

type IntegrationCreator interface {
	CreateIntegration(string, string, *atlas.ThirdPartyIntegration) (*atlas.ThirdPartyIntegrations, error)
}

// CreateIntegration encapsulate the logic to manage different cloud providers
func (s *Store) CreateIntegration(projectID, integrationType string, integration *atlas.ThirdPartyIntegration) (*atlas.ThirdPartyIntegrations, error) {
	switch s.service {
	case config.CloudService:
		resp, _, err := s.client.(*atlas.Client).Integrations.Replace(context.Background(), projectID, integrationType, integration)
		return resp, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}