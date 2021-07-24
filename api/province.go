package handler

import (
	"api-bed-covid/service/rest"
	"log"
	"net/http"
)

// ListProvinceHandler ...
func ListProvinceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Println("INFO: get list of provinces")

	rest.ResponseSuccessWriter(w, "Data ditemukan", rest.MapProvinceID.GetKeys())
}
