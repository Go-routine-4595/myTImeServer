package service

import (
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Time() string {
	return time.Now().Format(time.RFC3339)
}
