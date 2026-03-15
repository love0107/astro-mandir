package prokerala

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

type PanchaangData struct {
	Date      string `json:"date"`
	Day       string `json:"day"`
	Tithi     string `json:"tithi"`
	Nakshatra string `json:"nakshatra"`
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	Moonrise  string `json:"moonrise"`
	Yoga      string `json:"yoga"`
	Vrat      string `json:"vrat"`
	Muhurat   string `json:"muhurat"`
}

// Prokerala exact response structure
type prokeralaResponse struct {
	Status string `json:"status"`
	Data   struct {
		Vaara     string `json:"vaara"`
		Sunrise   string `json:"sunrise"`
		Sunset    string `json:"sunset"`
		Moonrise  string `json:"moonrise"`
		Nakshatra []struct {
			Name string `json:"name"`
		} `json:"nakshatra"`
		Tithi []struct {
			Name   string `json:"name"`
			Paksha string `json:"paksha"`
		} `json:"tithi"`
		Yoga []struct {
			Name string `json:"name"`
		} `json:"yoga"`
	} `json:"data"`
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	accessToken = result.AccessToken
	tokenExpiry = time.Now().Add(time.Duration(result.ExpiresIn-60) * time.Second)

	return accessToken, nil
}

func FetchPanchang(date string, lat string, lng string) (*PanchaangData, error) {
	token, err := getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("token error: %v", err)
	}

	// Use dynamic coordinates
	coordinates := lat + "," + lng

	apiURL := fmt.Sprintf(
		"https://api.prokerala.com/v2/astrology/panchang?datetime=%sT06:00:00%%2B05:30&coordinates=%s&ayanamsa=1",
		date, coordinates,
	)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result prokeralaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.Status != "ok" {
		return nil, fmt.Errorf("prokerala error")
	}

	// Extract first tithi and nakshatra
	tithi := ""
	nakshatra := ""
	yoga := ""

	if len(result.Data.Tithi) > 0 {
		tithi = result.Data.Tithi[0].Name + " — " + result.Data.Tithi[0].Paksha
	}
	if len(result.Data.Nakshatra) > 0 {
		nakshatra = result.Data.Nakshatra[0].Name
	}
	if len(result.Data.Yoga) > 0 {
		yoga = result.Data.Yoga[0].Name
	}

	// Format time — extract only HH:MM from ISO string
	sunrise := formatTime(result.Data.Sunrise)
	sunset := formatTime(result.Data.Sunset)
	moonrise := formatTime(result.Data.Moonrise)

	return &PanchaangData{
		Date:      date,
		Day:       result.Data.Vaara,
		Tithi:     tithi,
		Nakshatra: nakshatra,
		Sunrise:   sunrise,
		Sunset:    sunset,
		Moonrise:  moonrise,
		Yoga:      yoga,
		Vrat:      "",
	}, nil
}

// formatTime — extract HH:MM from ISO string
// "2026-01-01T06:59:15+05:30" → "06:59 AM"
func formatTime(iso string) string {
	if len(iso) < 16 {
		return iso
	}
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		return iso
	}
	return t.Format("03:04 PM")
}

func HardcodedPanchang() *PanchaangData {
	now := time.Now()
	return &PanchaangData{
		Date:      now.Format("2006-01-02"),
		Day:       now.Weekday().String(),
		Tithi:     "Tritiya",
		Nakshatra: "Rohini",
		Sunrise:   "06:42 AM",
		Sunset:    "06:18 PM",
		Moonrise:  "07:15 PM",
		Yoga:      "Subha",
		Vrat:      "Mangala Gauri Vrat",
	}
}
