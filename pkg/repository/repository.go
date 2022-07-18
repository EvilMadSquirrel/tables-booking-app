package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
