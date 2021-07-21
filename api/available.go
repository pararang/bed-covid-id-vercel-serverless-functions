package handler

import (
	"api-bed-covid/service"
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

	provinceID, ok := service.MapProvinceID[provinceName]
	if !ok {
		service.JSONResponseFail(w, fmt.Sprintf("unknown province %s", provinceName))
		return
	}

	hospitals, err := service.ScrapeProvince(provinceID)
	if err != nil {
		service.JSONResponseFail(w, err.Error())
		return
	}

	sort.Slice(hospitals, func(i, j int) bool {
		return hospitals[i].BedAvailable > hospitals[j].BedAvailable
	})

	service.JSONResponseSuccess(w, "Data ditemukan", hospitals)
}
