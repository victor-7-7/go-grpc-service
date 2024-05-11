//+acceptance

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	rocket "go-course/grpc-service/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

//func (s *RocketTestSuite) SetupSuite() {}

func (s *RocketTestSuite) TestAddRocket() {
	log.Println("TestAddRocket")
	s.T().Run("Adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		resp, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "6f79f166-2a3b-4d20-83c6-4c4201e5b257",
					Name: "Rocket 1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), "6f79f166-2a3b-4d20-83c6-4c4201e5b257", resp.Rocket.Id)
	})

	s.T().Run("Validates the new rocket's Id is uuid", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "231c6ad0-4f4e-4214-a616-4609882baafa", //"not-a-valid-uuid"
					Name: "Rocket 1",
					Type: "Falcon Heavy",
				},
			},
		)

		assert.NoError(s.T(), err)

		st := status.Convert(err)
		assert.NotEqual(s.T(), codes.InvalidArgument, st.Code())
	})
}

// Для выполнения теста надо запустить сервис локально.
// Сначала поднимаем докер-компоуз: docker compose up --build
// Если порт 5432 занят - находим PID процесса (локальный postgres),
// который слушает порт: sudo lsof -i tcp:5432. И гасим: sudo kill xxxx.
// Затем в другом терминале командуем: go test ./test -tags=acceptance -v.

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}

//func (s *RocketTestSuite) TearDownSuite() {}
