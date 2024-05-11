//go:generate mockgen -source=rocket.go -destination=rocket_mocks_test.go -package=rocket Store
// https://utkarshmani1997.medium.com/mocking-with-mockgen-43513e3091b5

package rocket

import (
	"context"
	"log"
)

// Rocket - should contain the definition of our rocket.
type Rocket struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Flights int    `json:"flights"`
}

// Store - defines the interface we expect our database.
// implementation to follow
type Store interface {
	GetRocketById(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service, responsible for updating
// the rocket inventory
type Service struct {
	Store Store `json:"store"`
}

// New - returns a new instance of our rocket service.
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketById - retrieves the rocket based on the id from the store.
func (s Service) GetRocketById(ctx context.Context, id string) (Rocket, error) {
	log.Println("Service | Get Rocket by id:", id)
	rkt, err := s.Store.GetRocketById(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store.
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	log.Println("Service | Insert Rocket")
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket -deletes a rocket from our inventory.
func (s Service) DeleteRocket(ctx context.Context, id string) error {
	log.Println("Service | Delete Rocket (id):", id)
	if err := s.Store.DeleteRocket(id); err != nil {
		return err
	}
	return nil
}
