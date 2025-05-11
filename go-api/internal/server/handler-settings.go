package server

import (
	"api/internal/domain"
	"encoding/json"
	"net/http"
)

type settingsRequest struct {
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
}

func (s *server) getSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.saveSettingsHandler(w, r)
		return
	}

	setting, err := s.db.GetSettings()
	if err != nil {
		// TODO log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(setting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) saveSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var settingsRequest settingsRequest

	if err := json.NewDecoder(r.Body).Decode(&settingsRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	setting := domain.Setting{
		Version:              "1",
		Region:               settingsRequest.Region,
		Endpoint:             settingsRequest.Endpoint,
		UsePathStyleEndpoint: true,
		Credentials: domain.Credentials{
			Key:    settingsRequest.Key,
			Secret: settingsRequest.Secret,
		},
	}

	err := s.db.SaveSettings(setting)
	if err != nil {
		// TODO log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(setting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
