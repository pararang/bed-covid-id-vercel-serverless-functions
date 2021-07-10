package service

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var MapProvinceID = map[string]int{
	"ACEH":                      11,
	"SUMATERA UTARA":            12,
	"SUMATERA BARAT":            13,
	"RIAU":                      14,
	"JAMBI":                     15,
	"SUMATERA SELATAN":          16,
	"BENGKULU":                  17,
	"LAMPUNG":                   18,
	"KEPULAUAN BANGKA BELITUNG": 19,
	"KEPULAUAN RIAU":            21,
	"DKI JAKARTA":               31,
	"JAWA BARAT":                32,
	"JAWA TENGAH":               33,
	"DI YOGYAKARTA":             34,
	"JAWA TIMUR":                35,
	"BANTEN":                    36,
	"BALI":                      51,
	"NUSA TENGGARA BARAT":       52,
	"NUSA TENGGARA TIMUR":       53,
	"KALIMANTAN BARAT":          61,
	"KALIMANTAN TENGAH":         62,
	"KALIMANTAN SELATAN":        63,
	"KALIMANTAN TIMUR":          64,
	"KALIMANTAN UTARA":          65,
	"SULAWESI UTARA":            71,
	"SULAWESI TENGAH":           72,
	"SULAWESI SELATAN":          73,
	"SULAWESI TENGGARA":         74,
	"GORONTALO":                 75,
	"SULAWESI BARAT":            76,
	"MALUKU":                    81,
	"MALUKU UTARA":              82,
	"PAPUA BARAT":               91,
	"PAPUA":                     94,
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func EncodeResponse(response Response) *bytes.Buffer {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	_ = enc.Encode(response)

	return buf
}

func JSONResponseSuccess(w http.ResponseWriter, message string, data interface{}) {

	if data == nil {
		log.Println("INFO: response with empty data")
	}

	var response = Response{
		Status:  "OK",
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(EncodeResponse(response).Bytes())
}

func JSONResponseFail(w http.ResponseWriter, message string) {

	log.Printf("ERROR: %s", message)

	var response = Response{
		Status:  "ERROR",
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(EncodeResponse(response).Bytes())
}
