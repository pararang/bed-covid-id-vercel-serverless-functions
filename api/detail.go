package handler

import (
	"api-bed-covid/service/rest"
	"api-bed-covid/service/scraper"
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
		rest.ResponseFailWriter(w, fmt.Sprintf("invalid hospital code: %s", hospitalCode))
		return
	}

	scraperServices := scraper.New()

	detail, err := scraperServices.GetHospitalDetail(hospitalCode)
	if err != nil {
		rest.ResponseFailWriter(w, err.Error())
		return
	}

	if detail.IsEmpty() {
		rest.ResponseSuccessWriter(w, "Data tidak ditemukan", nil)
		return
	}

	rest.ResponseSuccessWriter(w, "Data ditemukan", detail)
}
