package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/love0107/astro-mandir/db"
	generated "github.com/love0107/astro-mandir/db/generated"
	"github.com/love0107/astro-mandir/internal/panchang/prokerala"
)

type PanchaangService struct{}

func NewPanchaangService() *PanchaangService {
	return &PanchaangService{}
}

func (s *PanchaangService) GetToday(ctx context.Context, lat string, lng string) (*prokerala.PanchaangData, error) {
	today := time.Now().Format("2006-01-02")

	// Step 1 — check DB first
	row, err := db.Queries.GetPanchang(ctx, today)
	if err == nil {
		return &prokerala.PanchaangData{
			Date:      row.Date,
			Tithi:     row.Tithi.String,
			Nakshatra: row.Nakshatra.String,
			Sunrise:   row.Sunrise.String,
			Sunset:    row.Sunset.String,
			Vrat:      row.Vrat.String,
		}, nil
	}

	// Step 2 — call API with user coordinates
	data, err := prokerala.FetchPanchang(today, lat, lng)
	if err != nil {
		return prokerala.HardcodedPanchang(), nil
	}

	// Step 3 — save to DB
	db.Queries.InsertPanchang(ctx, generated.InsertPanchangParams{
		Date:      data.Date,
		Vrat:      sql.NullString{String: data.Vrat, Valid: true},
		Tithi:     sql.NullString{String: data.Tithi, Valid: true},
		Nakshatra: sql.NullString{String: data.Nakshatra, Valid: true},
		Sunrise:   sql.NullString{String: data.Sunrise, Valid: true},
		Sunset:    sql.NullString{String: data.Sunset, Valid: true},
		Muhurat:   sql.NullString{String: "", Valid: false},
		Festival:  sql.NullString{String: "", Valid: false},
	})

	return data, nil
}
