package prokerala

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PanchaangData struct {
	Date      string `json:"date"`
	Day       string `json:"day"`
	Tithi     string `json:"tithi"`
	Nakshatra string `json:"nakshatra"`
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	Muhurat   string `json:"muhurat"`
	Vrat      string `json:"vrat"`
}

// FetchFromAPI — calls Prokerala API
func FetchFromAPI(date string) (*PanchaangData, error) {
	url := fmt.Sprintf(
		"https://api.prokerala.com/v2/astrology/panchang?datetime=%sT06:00:00&coordinates=26.8467,80.9462&ayanamsa=1",
		date,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result PanchaangData
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// HardcodedPanchang — fallback if API fails
func HardcodedPanchang() *PanchaangData {
	now := time.Now()
	return &PanchaangData{
		Date:      now.Format("2006-01-02"),
		Day:       now.Weekday().String(),
		Tithi:     "Tritiya",
		Nakshatra: "Rohini",
		Sunrise:   "6:42 AM",
		Sunset:    "6:18 PM",
		Muhurat:   "11:30 AM - 12:30 PM",
		Vrat:      "Mangala Gauri Vrat",
	}
}
