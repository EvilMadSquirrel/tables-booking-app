package repository

import "github.com/jmoiron/sqlx"

type Reserve interface {
}

type BookingList interface {
}

type Booking interface {
}

type Repository struct {
	Reserve
	BookingList
	Booking
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
