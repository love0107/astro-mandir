package service

import (
	"context"
	"database/sql"

	"github.com/love0107/astro-mandir/db"
	generated "github.com/love0107/astro-mandir/db/generated"
	"github.com/love0107/astro-mandir/internal/kundali"
)

type KundaliService struct{}

func NewKundaliService() *KundaliService {
	return &KundaliService{}
}

func (s *KundaliService) Generate(ctx context.Context, req kundali.KundaliRequest) (*kundali.KundaliData, error) {
	// Call real Prokerala API
	result, err := kundali.FetchKundali(req)
	if err != nil {
		// Fallback to hardcoded
		result = kundali.HardcodedKundali(req)
	}

	// Save to DB
	db.Queries.CreateKundaliRequest(ctx, generated.CreateKundaliRequestParams{
		Name:  sql.NullString{String: req.Name, Valid: true},
		Dob:   sql.NullString{String: req.DOB, Valid: true},
		Tob:   sql.NullString{String: req.TOB, Valid: true},
		Place: sql.NullString{String: req.Latitude + "," + req.Longitude, Valid: true},
		Rashi: sql.NullString{String: result.ChandraRasi, Valid: true},
	})

	return result, nil
}
