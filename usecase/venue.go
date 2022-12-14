package usecase

import (
	"github.com/scarlet0725/prism-api/framework"
	"github.com/scarlet0725/prism-api/infrastructure/repository"
	"github.com/scarlet0725/prism-api/model"
)

type Venue interface {
	CreateVenue(*model.Venue) (*model.Venue, error)
	UpdateVenue(*model.Venue) (*model.Venue, error)
	//GetVenueByName(string) (*model.Venue, error)
	GetVenueByID(string) (*model.Venue, error)
}

type venueUsecase struct {
	db     repository.DB
	ramdom framework.RandomID
}

func NewVenueUsecase(db repository.DB, r framework.RandomID) Venue {
	return &venueUsecase{
		db:     db,
		ramdom: r,
	}
}

func (a *venueUsecase) CreateVenue(venue *model.Venue) (*model.Venue, error) {
	id := a.ramdom.Generate(venueIDLength)

	venue.VenueID = id
	venue.ID = 0
	venue.IsOpen = true

	venue, err := a.db.CreateVenue(venue)

	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (a *venueUsecase) UpdateVenue(venue *model.Venue) (*model.Venue, error) {
	venue, err := a.db.UpdateVenue(venue)

	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (a *venueUsecase) GetVenueByID(id string) (*model.Venue, error) {
	venue, err := a.db.GetVenueByID(id)

	if err != nil {
		return nil, err
	}

	return venue, nil
}
