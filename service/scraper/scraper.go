package scraper

import (
	"api-bed-covid/model"
	"api-bed-covid/storage"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nedpals/supabase-go"
)

type Scraper interface {
	GetProvinceAvailability(provinceID int) ([]model.HospitalSummary, error)
	GetHospitalDetail(hospitalCode string) (model.HospitalDetail, error)
}

type scraper struct {
	cacheClient *supabase.Client
	cacheDB     *string
}

var redis = storage.NewRedis()

func New() scraper {

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseDB := os.Getenv("SUPABASE_DB")

	return scraper{
		cacheClient: supabase.CreateClient(supabaseUrl, supabaseKey),
		cacheDB:     &supabaseDB,
	}
}

// scanHospitalSummaryFromCardSelector get hospital summary for data available bed
func (s *scraper) scanHospitalSummaryFromCardSelector(hospitals *[]model.HospitalSummary, card *goquery.Selection, wg *sync.WaitGroup) {
	var err error
	var hospital = new(model.HospitalSummary)

	defer wg.Done()

	hospital.Name = card.Find("h5").Text()

	siranapHospitalURL, exist := card.Find("a[href]").Attr("href")
	if !exist {
		log.Println("INFO: not found selector siranap hospital detail URL")
	}

	hospital.Code, err = s.getHospitalCodeFromDetailURL(siranapHospitalURL)
	if err != nil {
		log.Printf("INFO: failed get hospital code, err: %s", err.Error())
	}

	hospital.DetailURL = siranapHospitalURL

	card.Find("p").Each(func(i int, subSel *goquery.Selection) {
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

	card.Find(".card-footer").Each(func(i int, footerSel *goquery.Selection) {
		hotline := footerSel.Find("span").Text()
		if hotline != "hotline tidak tersedia" {
			hospital.Hotline = hotline
		}
	})

	*hospitals = append(*hospitals, *hospital)
}

func (s *scraper) GetProvinceAvailability(provinceID int) ([]model.HospitalSummary, error) {
	var data = make([]model.HospitalSummary, 0)
	var url = fmt.Sprintf("http://yankes.kemkes.go.id/app/siranap/rumah_sakit?jenis=1&propinsi=%dprop&kabkota", provinceID)

	domHTML, err := s.readPage(url)
	if err != nil {
		return data, err
	}

	cachedData, _ := redis.GetScrapedAvailableHospitals(url)
	if len(cachedData) > 0 {
		err := json.Unmarshal([]byte(cachedData), &data)
		if err == nil {
			log.Printf("INFO: Return data from cached data")
			return data, err
		}
	}

	var wg sync.WaitGroup

	domHTML.Find(".cardRS").Each(func(i int, sel *goquery.Selection) {
		wg.Add(1)
		go s.scanHospitalSummaryFromCardSelector(&data, sel, &wg)
	})

	wg.Wait()

	go func() {
		err = redis.SetScrapedAvailableHospitals(url, data)
		if err != nil {
			log.Println("ERROR: Error set to redis")
		}
	}()

	return data, nil
}

// scanRoomsDetailFromCardSelector get rooms detail from card selector
func (s *scraper) scanRoomsDetailFromCardSelector(rooms *[]model.Room, card *goquery.Selection, wg *sync.WaitGroup) {
	var err error
	var room = new(model.Room)

	defer wg.Done()

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

	*rooms = append(*rooms, *room)
}

func (s *scraper) GetHospitalDetail(hospitalCode string) (model.HospitalDetail, error) {
	var getHospitalName = func(titleText, address, hotline string) string {
		titleText = strings.Replace(titleText, address, "", 1)
		titleText = strings.Replace(titleText, hotline, "", 1)
		return titleText
	}

	var data model.HospitalDetail
	var provinceID = hospitalCode[0:2]
	var url = fmt.Sprintf("https://yankes.kemkes.go.id/app/siranap/tempat_tidur?kode_rs=%s&jenis=1&propinsi=%sprop&kabkota=", hospitalCode, provinceID)

	domHTML, err := s.readPage(url)
	if err != nil {
		return data, err
	}

	cachedData, _ := redis.GetScrapedDetailHospital(url)
	if len(cachedData) > 0 {
		err := json.Unmarshal([]byte(cachedData), &data)
		if err == nil {
			log.Printf("INFO: Return data from cached data")
			return data, err
		}
	}

	titleSelector := domHTML.Find("p[class=mb-0]").First()

	data.Address = strings.TrimSpace(titleSelector.Find("small").First().Text())
	data.Hotline = strings.TrimSpace(titleSelector.Find("i").First().Text())
	data.Name = strings.TrimSpace(getHospitalName(titleSelector.Text(), data.Address, data.Hotline))

	var rooms = make([]model.Room, 0)
	var wg sync.WaitGroup

	domHTML.Find(".card").Each(func(i int, card *goquery.Selection) {
		wg.Add(1)
		go s.scanRoomsDetailFromCardSelector(&rooms, card, &wg)
	})

	wg.Wait()

	data.Room = rooms

	go func() {
		err = redis.SetScrapedDetailHospital(url, data)
		if err != nil {
			log.Println("ERROR: Error set to redis")
		}
	}()

	return data, nil
}

func (s *scraper) readPage(url string) (goQueryDoc *goquery.Document, err error) {
	log.Printf("INFO: Read page %s", url)

	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 120 * time.Second,
	}

	client := &http.Client{
		Transport: t,
	}

	response, err := client.Get(url)
	if err != nil {
		return goQueryDoc, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return goQueryDoc, fmt.Errorf("HTTP error: %d %s", response.StatusCode, response.Status)

	}

	return goquery.NewDocumentFromReader(response.Body)
}

func (s *scraper) getHospitalCodeFromDetailURL(detailURL string) (code string, err error) {
	u, err := url.Parse(detailURL)
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
