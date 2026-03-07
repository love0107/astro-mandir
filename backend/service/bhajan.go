package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/love0107/astro-mandir/db"
	"github.com/love0107/astro-mandir/internal/youtube"
)

type BhajanService struct{}

func NewBhajanService() *BhajanService {
	return &BhajanService{}
}

func (s *BhajanService) GetTodayBhajan(ctx context.Context, rashi string) (*youtube.BhajanData, error) {
	today := time.Now().Format("2006-01-02")

	// Convert plain string to sql.NullString
	nullRashi := sql.NullString{
		String: rashi,
		Valid:  rashi != "",
	}

	// Step 1 — check DB
	row, err := db.Queries.GetBhajanByRashi(ctx, nullRashi)
	if err == nil {
		return &youtube.BhajanData{
			Title:     row.Title,
			YoutubeID: row.YoutubeID,
			EmbedURL:  "https://www.youtube.com/embed/" + row.YoutubeID,
		}, nil
	}

	// Step 2 — search YouTube
	query := buildBhajanQuery(rashi, today)
	data, err := youtube.SearchBhajan(query)
	if err != nil {
		return youtube.HardcodedBhajan(), nil
	}

	return data, nil
}

func buildBhajanQuery(rashi string, date string) string {
	// Personalize query based on rashi
	rashiBhajan := map[string]string{
		"mesh":      "hanuman+bhajan",
		"vrishabh":  "lakshmi+bhajan",
		"mithun":    "ganesh+bhajan",
		"kark":      "shiv+bhajan",
		"simha":     "surya+bhajan",
		"kanya":     "durga+bhajan",
		"tula":      "saraswati+bhajan",
		"vrishchik": "kali+bhajan",
		"dhanu":     "vishnu+bhajan",
		"makar":     "shani+bhajan",
		"kumbh":     "saraswati+bhajan",
		"meen":      "krishna+bhajan",
	}

	if query, ok := rashiBhajan[rashi]; ok {
		return query + "+morning+aarti"
	}

	// Default if rashi not provided
	return "aaj+ka+bhajan+morning+bhakti"
}
