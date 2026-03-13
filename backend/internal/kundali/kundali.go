package kundali

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	accessToken string
	tokenExpiry time.Time
	tokenMu     sync.Mutex
)

type KundaliRequest struct {
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	TOB       string `json:"tob"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type KundaliData struct {
	Name          string       `json:"name"`
	Nakshatra     string       `json:"nakshatra"`
	NakshatraLord string       `json:"nakshatra_lord"`
	ChandraRasi   string       `json:"chandra_rasi"`
	SooryaRasi    string       `json:"soorya_rasi"`
	Zodiac        string       `json:"zodiac"`
	Color         string       `json:"color"`
	BirthStone    string       `json:"birth_stone"`
	Gender        string       `json:"gender"`
	BestDirection string       `json:"best_direction"`
	MangalDosha   bool         `json:"mangal_dosha"`
	MangalDesc    string       `json:"mangal_description"`
	Yogas         []YogaDetail `json:"yogas"`
}

type YogaDetail struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getAccessToken() (string, error) {
	tokenMu.Lock()
	defer tokenMu.Unlock()

	if accessToken != "" && time.Now().Before(tokenExpiry) {
		return accessToken, nil
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", os.Getenv("PROKERALA_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("PROKERALA_CLIENT_SECRET"))

	resp, err := http.Post(
		"https://api.prokerala.com/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	json.Unmarshal(body, &result)

	accessToken = result.AccessToken
	tokenExpiry = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second)

	return accessToken, nil
}

func FetchKundali(req KundaliRequest) (*KundaliData, error) {
	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	// Combine date and time
	// DOB: "1990-05-15", TOB: "14:30"
	datetime := fmt.Sprintf("%sT%s:00%%2B05:30", req.DOB, req.TOB)

	// Default coordinates to Lucknow if not provided
	coordinates := req.Latitude + "," + req.Longitude
	if req.Latitude == "" || req.Longitude == "" {
		coordinates = "26.8467,80.9462"
	}

	apiURL := fmt.Sprintf(
		"https://api.prokerala.com/v2/astrology/kundli?ayanamsa=1&coordinates=%s&datetime=%s&la=hi",
		coordinates, datetime,
	)

	httpReq, _ := http.NewRequest("GET", apiURL, nil)
	httpReq.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Exact response structure from Prokerala
	var result struct {
		Status string `json:"status"`
		Data   struct {
			NakshatraDetails struct {
				Nakshatra struct {
					Name string `json:"name"`
					Lord struct {
						VedicName string `json:"vedic_name"`
					} `json:"lord"`
				} `json:"nakshatra"`
				ChandraRasi struct {
					Name string `json:"name"`
				} `json:"chandra_rasi"`
				SooryaRasi struct {
					Name string `json:"name"`
				} `json:"soorya_rasi"`
				Zodiac struct {
					Name string `json:"name"`
				} `json:"zodiac"`
				AdditionalInfo struct {
					Color         string `json:"color"`
					BirthStone    string `json:"birth_stone"`
					Gender        string `json:"gender"`
					BestDirection string `json:"best_direction"`
				} `json:"additional_info"`
			} `json:"nakshatra_details"`
			MangalDosha struct {
				HasDosha    bool   `json:"has_dosha"`
				Description string `json:"description"`
			} `json:"mangal_dosha"`
			YogaDetails []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"yoga_details"`
		} `json:"data"`
	}

	json.Unmarshal(body, &result)

	if result.Status != "ok" {
		return nil, fmt.Errorf("kundali API error")
	}

	// Build yogas list
	yogas := []YogaDetail{}
	for _, y := range result.Data.YogaDetails {
		yogas = append(yogas, YogaDetail{
			Name:        y.Name,
			Description: y.Description,
		})
	}

	return &KundaliData{
		Name:          req.Name,
		Nakshatra:     result.Data.NakshatraDetails.Nakshatra.Name,
		NakshatraLord: result.Data.NakshatraDetails.Nakshatra.Lord.VedicName,
		ChandraRasi:   result.Data.NakshatraDetails.ChandraRasi.Name,
		SooryaRasi:    result.Data.NakshatraDetails.SooryaRasi.Name,
		Zodiac:        result.Data.NakshatraDetails.Zodiac.Name,
		Color:         result.Data.NakshatraDetails.AdditionalInfo.Color,
		BirthStone:    result.Data.NakshatraDetails.AdditionalInfo.BirthStone,
		Gender:        result.Data.NakshatraDetails.AdditionalInfo.Gender,
		BestDirection: result.Data.NakshatraDetails.AdditionalInfo.BestDirection,
		MangalDosha:   result.Data.MangalDosha.HasDosha,
		MangalDesc:    result.Data.MangalDosha.Description,
		Yogas:         yogas,
	}, nil
}

// Fallback if API fails
func HardcodedKundali(req KundaliRequest) *KundaliData {
	return &KundaliData{
		Name:          req.Name,
		Nakshatra:     "Rohini",
		NakshatraLord: "Chandra",
		ChandraRasi:   "Vrishabh",
		SooryaRasi:    "Mesh",
		Zodiac:        "Taurus",
		Color:         "White",
		BirthStone:    "Pearl",
		BestDirection: "North",
		MangalDosha:   false,
		MangalDesc:    "Not Manglik",
		Yogas:         []YogaDetail{},
	}
}
