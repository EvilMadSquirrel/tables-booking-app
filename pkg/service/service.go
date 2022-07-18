package service

import "tables-booking-app/pkg/repository"

type Reserve interface {
}

type BookingList interface {
}

type Booking interface {
}

type Service struct {
	Reserve
	BookingList
	Booking
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
