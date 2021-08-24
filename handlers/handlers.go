package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hellofreshdevtests/GitKaran-devops-test/database"
	"github.com/hellofreshdevtests/GitKaran-devops-test/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

type APIEnv struct {
	DB *gorm.DB
}

// POST mapping - Create a new config
func (a *APIEnv) CreateConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	config := models.Config{}

	//Decode Json Request
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.DB.Create(&config).Error; err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, config)
}

// GET mapping - GetAllConfigs will return all configs
func (a *APIEnv) GetAllConfigs(w http.ResponseWriter, r *http.Request) {

	result, err := database.GetAllConfigs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// GET mapping - GetConfig will return a config by its name
func (a *APIEnv) GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	// get request params
	params := mux.Vars(r)
	name := params["name"]

	result, exists, err := database.GetConfigByName(name, a.DB)
	if !exists {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// GET mapping - SearchConfigByKey will return config by query search
func (a *APIEnv) SearchConfigByKey(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query()

	var pathKey, pathValue string
	for k, v := range path {
		pathKey = k
		pathValue = v[len(v)-1]
	}

	result, exists, err := database.SearchConfigByKey(pathKey, pathValue, a.DB)
	if !exists {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// DELETE mapping - DeleteConfig deletes config
func (a *APIEnv) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	if !isString(name) {
		respondWithError(w, http.StatusBadRequest, "Requested config name is not valid")
		return
	}

	_, exists, err := database.GetConfigByName(name, a.DB)
	if !exists {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = database.DeleteConfig(name, a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "deleted sucessfully"})
}

// PUT mapping - UpdateConfig updates existing config
func (a *APIEnv) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	name := params["name"]
	updatedConfig := models.Config{}

	if !isString(name) {
		respondWithError(w, http.StatusBadRequest, "Requested config name is not valid")
		return
	}

	existingConfig, exists, err := database.GetConfigByName(name, a.DB)
	if !exists {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// decode the json request
	decodeErr := json.NewDecoder(r.Body).Decode(&updatedConfig)

	if decodeErr != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = database.UpdateConfig(&existingConfig, &updatedConfig, a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusAccepted, updatedConfig)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func isString(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}
