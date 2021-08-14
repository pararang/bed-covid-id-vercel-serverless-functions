package rest

import (
	"api-bed-covid/utils"
	"bytes"
	"log"
	"net/http"
	"sort"
)

const (
	statusOK    = "OK"
	statusError = "ERROR"
)

// Option ...
type Option struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

// MapStringInt ...
type MapStringInt map[string]int

// GetKeys returns the keys of a map as a string array.
func (m MapStringInt) GetKeys() []string {
	var keys = make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetListForOptions returns a list of objects from a map.
func (m MapStringInt) GetListForOptions() []Option {
	var list = make([]Option, 0)
	for k, v := range m {
		list = append(list, Option{ID: v, Label: k})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Label < list[j].Label
	})

	return list
}

// MapProvinceID ...
var MapProvinceID = MapStringInt{
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
	"PAPUA":                     92,
}

// Response ...
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func EncodeResponse(response Response) *bytes.Buffer {
	return utils.JSONIndentFormatter(response)
}

// ResponseSuccessWriter ...
func ResponseSuccessWriter(w http.ResponseWriter, message string, data interface{}) {

	if data == nil {
		log.Println("INFO: response with empty data")
	}

	var response = Response{
		Status:  statusOK,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(EncodeResponse(response).Bytes())
}

// ResponseFailWriter ...
func ResponseFailWriter(w http.ResponseWriter, message string) {

	log.Printf("ERROR: %s", message)

	var response = Response{
		Status:  statusError,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(EncodeResponse(response).Bytes())
}
