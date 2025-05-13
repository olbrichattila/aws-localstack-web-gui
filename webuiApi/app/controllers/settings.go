package controller

import (
	"encoding/json"
	"webuiApi/app/repositories/database"
	"webuiApi/app/repositories/domain"

	"github.com/olbrichattila/gofra/pkg/app/request"
)

type settingsRequest struct {
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
}

// SettingsAction function can take any parameters defined in the Di config
func GetSettingsAction(data database.Database) (domain.Setting, error) {
	return data.GetSettings()
}

func SaveSettingsAction(r request.Requester, data database.Database) (domain.Setting, error) {
	var req settingsRequest
	if err := json.Unmarshal([]byte(r.Body()), &req); err != nil {
		return domain.Setting{}, err
	}

	setting := domain.Setting{
		Version:              "1",
		Region:               req.Region,
		Endpoint:             req.Endpoint,
		UsePathStyleEndpoint: true,
		Credentials: domain.Credentials{
			Key:    req.Key,
			Secret: req.Secret,
		},
	}
	if err := data.SaveSettings(setting); err != nil {
		return domain.Setting{}, err
	}

	return data.GetSettings()
}
