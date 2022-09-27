package services

import (
	"github.com/patogallaiov/boltz-poc/config"
	"github.com/patogallaiov/boltz-poc/connectors"
	"github.com/patogallaiov/boltz-poc/storage"
)

type ConfigService struct {
	boltz connectors.BoltzConnector
	db    storage.DBConnector
}

func NewConfigService(appCfg config.Config, boltz connectors.BoltzConnector, db storage.DBConnector) *ConfigService {
	return &ConfigService{
		boltz,
		db,
	}
}

func (service *ConfigService) SaveConfig(request *storage.Config) (*storage.Config, error) {
	err := service.db.SaveConfig(request)
	if err != nil {
		return &storage.Config{}, err
	}
	return request, nil
}

func (service *ConfigService) GetConfigs() (any, error) {
	return service.db.GetConfigs()
}

func (service *ConfigService) GetConfig(keyConfig string) (any, error) {
	return service.db.GetConfig(keyConfig)
}
