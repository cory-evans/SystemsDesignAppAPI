package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/CoryEvans2324/SystemsDesignAppAPI/database"
	"github.com/CoryEvans2324/SystemsDesignAppAPI/models"
	"gorm.io/gorm/clause"
)

func Index(w http.ResponseWriter, r *http.Request) {

}

func UploadTracks(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tracks := models.LoadFromFile(file)
	log.Println(len(tracks))
	tx := database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"description", "status", "object_type_description", "shape_length", "geometry"}),
	})
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetTracks(w http.ResponseWriter, r *http.Request) {
	var tracks []models.Track
	result := database.DB.Limit(100).Find(&tracks)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

	log.Println(len(tracks))

	json.NewEncoder(w).Encode(tracks)
}