package model

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
