package handler

import (
	"api-bed-covid/service/rest"
	"api-bed-covid/service/scraper"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

// AvailableHandler ...
func AvailableHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	provinceName := strings.ToUpper(r.URL.Query().Get("province"))

	log.Printf("INFO: scraping for province %s", provinceName)

	provinceID, ok := rest.MapProvinceID[provinceName]
	if !ok {
		rest.ResponseFailWriter(w, fmt.Sprintf("unknown province %s", provinceName))
		return
	}

	scraperServices := scraper.New()

	hospitals, err := scraperServices.GetProvinceAvailability(provinceID)
	if err != nil {
		rest.ResponseFailWriter(w, err.Error())
		return
	}

	sort.Slice(hospitals, func(i, j int) bool {
		return hospitals[i].BedAvailable > hospitals[j].BedAvailable
	})

	rest.ResponseSuccessWriter(w, "Data ditemukan", hospitals)
}
