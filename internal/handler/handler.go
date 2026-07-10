package handler

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type ModelHandler[M gorm.Model] interface {
	GetById(id uint) (M, error)
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
