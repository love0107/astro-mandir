package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get your free key from:
// console.cloud.google.com
// → Create Project
// → Enable YouTube Data API v3
// → Create Credentials → API Key
const API_KEY = "YOUR_YOUTUBE_API_KEY"

type BhajanData struct {
	Title     string `json:"title"`
	YoutubeID string `json:"youtube_id"`
	EmbedURL  string `json:"embed_url"`
}

// SearchBhajan — searches YouTube for bhajan
func SearchBhajan(query string) (*BhajanData, error) {
	url := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/search?part=snippet&q=%s&type=video&videoEmbeddable=true&key=%s&maxResults=1",
		query, API_KEY,
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

	var result struct {
		Items []struct {
			ID struct {
				VideoID string `json:"videoId"`
			} `json:"id"`
			Snippet struct {
				Title string `json:"title"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return HardcodedBhajan(), nil
	}

	videoID := result.Items[0].ID.VideoID
	title := result.Items[0].Snippet.Title

	return &BhajanData{
		Title:     title,
		YoutubeID: videoID,
		EmbedURL:  fmt.Sprintf("https://www.youtube.com/embed/%s", videoID),
	}, nil
}

// HardcodedBhajan — fallback if API key not set
func HardcodedBhajan() *BhajanData {
	return &BhajanData{
		Title:     "Jai Ganesh Jai Ganesh Deva",
		YoutubeID: "pGCMNFfcLH8",
		EmbedURL:  "https://www.youtube.com/embed/pGCMNFfcLH8",
	}
}