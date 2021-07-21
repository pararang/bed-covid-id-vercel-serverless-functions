package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// HospitalSummary is a summary of a hospital
type HospitalSummary struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DetailURL    string `json:"detailURL"`
	BedAvailable int    `json:"bedAvailable"`
	PatientQueue int    `json:"patientQueue"`
	Hotline      string `json:"hotline"`
	Note         string `json:"note"`
	LastUpdate   string `json:"lastUpdate"`
}

// HospitalDetail is a detail of a hospital
type HospitalDetail struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Hotline string `json:"hotline"`
	Room    []Room `json:"rooms"`
}

// IsEmpty ...
func (hd *HospitalDetail) IsEmpty() bool {
	return hd.Name == "" && hd.Address == "" && hd.Hotline == "" && len(hd.Room) == 0 //TODO: create proper way to check empty
}

// Room is a room of a hospital
type Room struct {
	Name       string `json:"name"`
	Capacity   int    `json:"capacity"`
	Empty      int    `json:"empty"`
	Queue      int    `json:"queue"`
	LastUpdate string `json:"lastUpdate"`
}

func readPage(url string) (goQueryDoc *goquery.Document, err error) {
	log.Printf("INFO: Read page %s", url)

	response, err := http.Get(url)
	if err != nil {
		return goQueryDoc, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return goQueryDoc, fmt.Errorf("HTTP error: %d %s", response.StatusCode, response.Status)

	}

	return goquery.NewDocumentFromReader(response.Body)
}

func getHospitalCode(siranapHospitalURL string) (code string, err error) {
	u, err := url.Parse(siranapHospitalURL)
	if err != nil {
		return code, err
	}

	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return code, err
	}

	hospitalCodeParam, ok := m["kode_rs"]
	if !ok {
		return code, fmt.Errorf("not found query param kode_rs")
	}

	if len(hospitalCodeParam) == 0 {
		return code, fmt.Errorf("query param kode_rs is empty")
	}

	return hospitalCodeParam[0], nil
}

// ScrapeProvince ...
func ScrapeProvince(provinceID int) (data []HospitalSummary, err error) {

	domHTML, err := readPage(fmt.Sprintf("http://yankes.kemkes.go.id/app/siranap/rumah_sakit?jenis=1&propinsi=%dprop&kabkota", provinceID))
	if err != nil {
		return data, err
	}

	domHTML.Find(".cardRS").Each(func(i int, sel *goquery.Selection) {
		var hospital = new(HospitalSummary)

		hospital.Name = sel.Find("h5").Text()

		siranapHospitalURL, exist := sel.Find("a[href]").Attr("href")
		if !exist {
			log.Println("INFO: not found selector siranap hospital detail URL")
		}

		hospital.Code, err = getHospitalCode(siranapHospitalURL)
		if err != nil {
			log.Printf("INFO: failed get hospital code, err: %s", err.Error())
		}

		hospital.DetailURL = siranapHospitalURL

		sel.Find("p").Each(func(i int, subSel *goquery.Selection) {
			text := strings.TrimSpace(subSel.Text())

			if i == 0 {
				hospital.Address = text
			}

			if i == 2 && text != "Bed IGD Penuh!" {
				bedAvailText := subSel.Find("b").Text()
				if bedAvailText != "" {
					hospital.BedAvailable, _ = strconv.Atoi(bedAvailText)
				}
			}

			if i == 3 && strings.HasPrefix(text, "dengan antrian") {
				inLineElements := strings.Split(text, " ")
				if len(inLineElements) == 4 {
					hospital.PatientQueue, _ = strconv.Atoi(inLineElements[2])
				}
			}

			if i == 4 {
				hospital.LastUpdate = strings.Replace(text, "diupdate ", "", 1)
			}

			if i == 5 {
				hospital.Note = text
			}

		})

		sel.Find(".card-footer").Each(func(i int, footerSel *goquery.Selection) {
			hotline := footerSel.Find("span").Text()
			if hotline != "hotline tidak tersedia" {
				hospital.Hotline = hotline
			}
		})

		data = append(data, *hospital)
	})

	return data, nil
}

// ScrapeHospital ...
func ScrapeHospital(hospitalCode string) (data HospitalDetail, err error) {

	var getHospitalName = func(titleText, address, hotline string) string {
		titleText = strings.Replace(titleText, address, "", 1)
		titleText = strings.Replace(titleText, hotline, "", 1)
		return titleText
	}

	provinceID := hospitalCode[0:2]

	domHTML, err := readPage(fmt.Sprintf("https://yankes.kemkes.go.id/app/siranap/tempat_tidur?kode_rs=%s&jenis=1&propinsi=%sprop&kabkota=", hospitalCode, provinceID))
	if err != nil {
		return data, err
	}

	titleSelector := domHTML.Find("p[class=mb-0]").First()

	data.Address = strings.TrimSpace(titleSelector.Find("small").First().Text())
	data.Hotline = strings.TrimSpace(titleSelector.Find("i").First().Text())
	data.Name = strings.TrimSpace(getHospitalName(titleSelector.Text(), data.Address, data.Hotline))

	var rooms []Room
	domHTML.Find(".card").Each(func(i int, card *goquery.Selection) {
		var room = new(Room)

		description := strings.Split(card.Find("p[class=mb-0]").Text(), "Update")
		if len(description) == 2 {
			room.Name = strings.TrimSpace(description[0])
			room.LastUpdate = strings.TrimSpace(description[1])
		}

		card.Find(".text-center").Each(func(i int, cardData *goquery.Selection) {

			text := strings.TrimSpace(cardData.Text())

			if strings.HasPrefix(text, "Tempat Tidur") {
				room.Capacity, err = strconv.Atoi(strings.TrimSpace(strings.Replace(text, "Tempat Tidur", "", 1)))
				if err != nil {
					log.Printf("INFO: error on convert room capacity, err: %s", err.Error())
				}
			}

			if strings.HasPrefix(text, "Kosong") {
				room.Empty, err = strconv.Atoi(strings.TrimSpace(strings.Replace(text, "Kosong", "", 1)))
				if err != nil {
					log.Printf("INFO: error on convert room empty bed, err: %s", err.Error())
				}
			}

			if strings.HasPrefix(text, "Antrian") {
				room.Empty, err = strconv.Atoi(strings.TrimSpace(strings.Replace(text, "Antrian", "", 1)))
				if err != nil {
					log.Printf("INFO: error on convert room queue, err: %s", err.Error())
				}
			}
		})

		rooms = append(rooms, *room)
	})

	data.Room = rooms

	return data, nil
}
