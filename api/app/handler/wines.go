package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/guipassos/go-wine-house/api/app/model"
	"github.com/jinzhu/gorm"
)

func GetAllWines(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	wines := []model.Wine{}
	db.Find(&wines)
	respondJSON(w, http.StatusOK, wines)
}

func CreateWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	wine := model.Wine{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wine); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&wine).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, wine)
}

func GetWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	wine := getWineOr404(db, uint(id), w, r)
	if wine == nil {
		return
	}
	respondJSON(w, http.StatusOK, wine)
}

func UpdateWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	wine := getWineOr404(db, uint(id), w, r)
	if wine == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wine); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&wine).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, wine)
}

func DeleteWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	wine := getWineOr404(db, uint(id), w, r)
	if wine == nil {
		return
	}
	if err := db.Delete(&wine).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"message": "record was removed"})
}

func DisableWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	wine := getWineOr404(db, uint(id), w, r)
	if wine == nil {
		return
	}
	wine.Disable()
	if err := db.Save(&wine).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, wine)
}

func EnableWine(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusNotAcceptable, err.Error())
		return
	}

	wine := getWineOr404(db, uint(id), w, r)
	if wine == nil {
		return
	}
	wine.Enable()
	if err := db.Save(&wine).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, wine)
}

// getWineOr404 gets a wine instance if exists, or respond the 404 error otherwise
func getWineOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.Wine {
	wine := model.Wine{}
	if err := db.First(&wine, model.Wine{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &wine
}
