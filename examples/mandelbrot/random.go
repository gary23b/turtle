package main

import (
	"golang.org/x/exp/rand"
)

type SavedRandoms struct {
	savedList []float64
	index     int
}

func NewSavedRandoms() *SavedRandoms {
	savedCount := 1000000 // 1 million
	ret := &SavedRandoms{
		savedList: make([]float64, 0, savedCount),
		index:     0,
	}

	for range savedCount {
		ret.savedList = append(ret.savedList, rand.Float64())
	}
	return ret
}

func (s *SavedRandoms) RandFloat64() float64 {
	ret := s.savedList[s.index]
	s.index++
	if s.index >= len(s.savedList) {
		s.index = 0
	}
	return ret
}
