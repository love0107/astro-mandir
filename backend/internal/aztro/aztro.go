package aztro

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Completely free API — no key needed!
const AZTRO_URL = "https://aztro.sameerkumar.website"

type RashifalData struct {
	Rashi         string `json:"rashi"`
	Description   string `json:"description"`
	LuckyNumber   string `json:"lucky_number"`
	LuckyColor    string `json:"lucky_color"`
	Mood          string `json:"mood"`
	Compatibility string `json:"compatibility"`
	Date          string `json:"date_range"`
}

// Hindi to English rashi mapping
var RashiMap = map[string]string{
	"mesh":      "aries",
	"vrishabh":  "taurus",
	"mithun":    "gemini",
	"kark":      "cancer",
	"simha":     "leo",
	"kanya":     "virgo",
	"tula":      "libra",
	"vrishchik": "scorpio",
	"dhanu":     "sagittarius",
	"makar":     "capricorn",
	"kumbh":     "aquarius",
	"meen":      "pisces",
}

func FetchRashifal(rashi string) (*RashifalData, error) {
	// Convert Hindi rashi to English
	sign, ok := RashiMap[rashi]
	if !ok {
		return nil, fmt.Errorf("invalid rashi: %s", rashi)
	}

	// Aztro API — POST request, no key needed
	url := fmt.Sprintf("%s/?sign=%s&day=today", AZTRO_URL, sign)

	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Description   string `json:"description"`
		LuckyNumber   string `json:"lucky_number"`
		Color         string `json:"color"`
		Mood          string `json:"mood"`
		Compatibility string `json:"compatibility"`
		DateRange     string `json:"date_range"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &RashifalData{
		Rashi:         rashi,
		Description:   result.Description,
		LuckyNumber:   result.LuckyNumber,
		LuckyColor:    result.Color,
		Mood:          result.Mood,
		Compatibility: result.Compatibility,
		Date:          result.DateRange,
	}, nil
}

// Hardcoded fallback if API fails
func HardcodedRashifal(rashi string) *RashifalData {
	return &RashifalData{
		Rashi:         rashi,
		Description:   "Aaj ka din aapke liye shubh hai. Mehnat rang laayegi. Parivar ka saath milega.",
		LuckyNumber:   "7",
		LuckyColor:    "Laal",
		Mood:          "Khush",
		Compatibility: "Mithun",
		Date:          "Aaj",
	}
}
