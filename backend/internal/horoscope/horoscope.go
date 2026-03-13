package horoscope

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

type RashifalData struct {
	Rashi       string `json:"rashi"`
	Description string `json:"description"`
	LuckyNumber string `json:"lucky_number"`
	LuckyColor  string `json:"lucky_color"`
	Mood        string `json:"mood"`
}

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

func FetchRashifal(rashi string) (*RashifalData, error) {
	sign, ok := RashiMap[rashi]
	if !ok {
		return nil, fmt.Errorf("invalid rashi: %s", rashi)
	}

	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")
	apiURL := fmt.Sprintf(
		"https://api.prokerala.com/v2/horoscope/daily?datetime=%sT06:00:00%%2B05:30&sign=%s&type=general&ayanamsa=1",
		today, sign,
	)

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result struct {
		Status string `json:"status"`
		Data   struct {
			DailyPrediction struct {
				SignID     int    `json:"sign_id"`
				SignName   string `json:"sign_name"`
				Date       string `json:"date"`
				Prediction string `json:"prediction"`
			} `json:"daily_prediction"`
		} `json:"data"`
	}
	json.Unmarshal(body, &result)

	if result.Status != "ok" {
		return nil, fmt.Errorf("rashifal API error")
	}

	return &RashifalData{
		Rashi:       rashi,
		Description: result.Data.DailyPrediction.Prediction,
	}, nil
}

func HardcodedRashifal(rashi string) *RashifalData {
	return &RashifalData{
		Rashi:       rashi,
		Description: "Aaj ka din aapke liye shubh hai. Mehnat rang laayegi.",
		LuckyNumber: "7",
		LuckyColor:  "Laal",
		Mood:        "Khush",
	}
}
