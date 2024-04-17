//go:generate mockgen -source=rocket.go -destination=rocket_mocks_test.go -package=rocket Store
// https://utkarshmani1997.medium.com/mocking-with-mockgen-43513e3091b5

package rocket

import "context"

// Rocket - should contain the definition of our rocket
type Rocket struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Flights int    `json:"flights"`
}

// Store - defines the interface we expect our database
// implementation to follow
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service, responsible for updating
// the rocket inventory
type Service struct {
	Store Store `json:"store"`
}

// New - returns a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

// GetRocketByID - retrieves the rocket based on the ID from the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// InsertRocket - inserts a new rocket into the store.
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil
}

// DeleteRocket -deletes a rocket from our inventory
func (s Service) DeleteRocket(id string) error {
	if err := s.Store.DeleteRocket(id); err != nil {
		return err
	}
	return nil
}
