package main

import (
	"fmt"
	"context"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(start time.Time){
		fmt.Printf("fact=%s error=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())
	return s.next.GetCatFact(ctx)
}