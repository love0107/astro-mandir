package service

import (
	"context"
	"fmt"

	"github.com/love0107/astro-mandir/internal/horoscope"
)

type RashifalService struct{}

func NewRashifalService() *RashifalService {
	return &RashifalService{}
}

func (s *RashifalService) GetRashifal(ctx context.Context, rashi string) (*horoscope.RashifalData, error) {
	_, ok := horoscope.RashiMap[rashi]
	if !ok {
		return nil, fmt.Errorf("invalid rashi: %s", rashi)
	}

	data, err := horoscope.FetchRashifal(rashi)
	if err != nil {
		return horoscope.HardcodedRashifal(rashi), nil
	}

	return data, nil
}
