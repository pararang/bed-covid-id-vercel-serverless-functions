package handler

import (
	"api-bed-covid/service"
	"log"
	"net/http"
)

// ListProvinceHandler ...
func ListProvinceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Println("INFO: get list of provinces")

	service.JSONResponseSuccess(w, "Data ditemukan", service.MapProvinceID.GetKeys())
}
