package service

import (
	"BookingPlatform/apartment-service/model"
	"BookingPlatform/apartment-service/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type ApartmentService struct {
	Repo   *repository.ApartmentRepository
	Logger *log.Logger
}

type AvailableApartment struct {
	Apartment  *model.Apartment
	TotalPrice int
	UnitPrice  int
}

func NewApartmentService(r *repository.ApartmentRepository, l *log.Logger) *ApartmentService {
	return &ApartmentService{r, l}
}

func (as *ApartmentService) Insert(apartment *model.Apartment, userRole string, userId string) (*model.Apartment, error) {
	if userRole == "HOST" {
		apartment.HostId, _ = primitive.ObjectIDFromHex(userId)
		return as.Repo.Insert(apartment)
	}
	return nil, fmt.Errorf("You are not authorized for this function!")
}

func (as *ApartmentService) GetAll() (model.Apartments, error) {
	apartments, err := as.Repo.GetAll()
	if err != nil {
		as.Logger.Println(err)
		return nil, err
	}
	return apartments, nil
}

func (as *ApartmentService) InsertPricelistItem(item *model.PricelistItem, apartmentId string, userRole string) error {
	if userRole == "HOST" {
		return as.Repo.InsertPricelistItem(item, apartmentId)
	}
	return fmt.Errorf("You are not authorized for this function!")
}

func (as *ApartmentService) FilterApartments(location string, guestsNumber int, start string, end string) ([]AvailableApartment, error) {
	apartments, err := as.GetAll()
	if err != nil {
		as.Logger.Println(err)
		return nil, err
	}

	var filteredApartments []*model.Apartment

	for i := 0; i < len(apartments); i++ {
		if apartments[i].Location == location && (apartments[i].MinGuestsNumber <= guestsNumber && apartments[i].MaxGuestsNumber >= guestsNumber) {
			filteredApartments = append(filteredApartments, apartments[i])
		}
	}

	if len(filteredApartments) == 0 {
		return nil, fmt.Errorf("No apartments by defined filter!")
	}

	var availableApartments []AvailableApartment

	for i := 0; i < len(filteredApartments); i++ {

		if len(filteredApartments[i].Pricelist) == 0 {
			return nil, fmt.Errorf("No apartments by defined filter!")
		}

		filterStartDate, _ := time.Parse("02-01-2006", start)
		filterEndDate, _ := time.Parse("02-01-2006", end)

		for j := 0; j < len(filteredApartments[i].Pricelist); j++ {
			availableStartDate, _ := time.Parse("02-01-2006", filteredApartments[i].Pricelist[j].AvailabilityStartDate)
			availableEndDate, _ := time.Parse("02-01-2006", filteredApartments[i].Pricelist[j].AvailabilityEndDate)
			if filterStartDate.After(availableStartDate) && filterEndDate.Before(availableEndDate) {
				if filteredApartments[i].Pricelist[j].UnitPrice == 0 {
					totalPrice := guestsNumber * filteredApartments[i].Pricelist[j].Price * int(filterEndDate.Sub(filterStartDate).Hours()/24)
					unitPrice := filteredApartments[i].Pricelist[j].Price
					availableApartments = append(availableApartments, AvailableApartment{filteredApartments[i], totalPrice, unitPrice})
				} else {
					totalPrice := filteredApartments[i].Pricelist[j].Price * int(filterEndDate.Sub(filterStartDate).Hours()/24)
					unitPrice := filteredApartments[i].Pricelist[j].Price
					availableApartments = append(availableApartments, AvailableApartment{filteredApartments[i], totalPrice, unitPrice})
				}
			}
		}

	}
	if len(availableApartments) == 0 {
		return nil, fmt.Errorf("No apartments by defined filter!")
	}
	return availableApartments, nil
}
