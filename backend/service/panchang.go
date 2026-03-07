package service

import (
	"context"
	"time"

	"github.com/love0107/astro-mandir/db"
	"github.com/love0107/astro-mandir/internal/panchang/prokerala"
)

type PanchaangService struct{}

func NewPanchaangService() *PanchaangService {
	return &PanchaangService{}
}

func (s *PanchaangService) GetToday(ctx context.Context) (*prokerala.PanchaangData, error) {
	today := time.Now().Format("2006-01-02")

	// Step 1 — check database first
	row, err := db.Queries.GetPanchang(ctx, today)
	if err == nil {
		return &prokerala.PanchaangData{
			Date:      row.Date,
			Tithi:     row.Tithi.String,
			Nakshatra: row.Nakshatra.String,
			Sunrise:   row.Sunrise.String,
			Sunset:    row.Sunset.String,
			Muhurat:   row.Muhurat.String,
			Vrat:      row.Vrat.String,
		}, nil
	}

	// Step 2 — return hardcoded for now
	// TODO: replace with real Prokerala API once key is ready
	return prokerala.HardcodedPanchang(), nil
}
