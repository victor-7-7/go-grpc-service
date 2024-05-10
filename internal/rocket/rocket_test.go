package rocket

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocketByID(id).
			Return(Rocket{Id: id}, nil)
		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketById(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.Id)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(Rocket{
				Id: id,
			}).
			Return(Rocket{
				Id: id,
			}, nil)

		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.InsertRocket(
			context.Background(),
			Rocket{Id: id},
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.Id)
	})

	t.Run("tests delete rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocket(id).
			Return(nil)

		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(context.Background(), id)
		assert.NoError(t, err)
	})
}
