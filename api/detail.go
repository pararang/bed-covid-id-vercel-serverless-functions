package handler

import (
	"api-bed-covid/service"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// DetailHospitalHandler ...
func DetailHospitalHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	hospitalCode := strings.ToUpper(r.URL.Query().Get("code"))

	log.Printf("INFO: scraping for hospital %s", hospitalCode)

	if len(hospitalCode) < 2 {
		service.JSONResponseFail(w, fmt.Sprintf("invalid hospital code: %s", hospitalCode))
		return
	}

	detail, err := service.ScrapeHospital(hospitalCode)
	if err != nil {
		service.JSONResponseFail(w, err.Error())
		return
	}

	if detail.IsEmpty() {
		service.JSONResponseSuccess(w, "Data tidak ditemukan", nil)
		return
	}

	service.JSONResponseSuccess(w, "Data ditemukan", detail)
}
