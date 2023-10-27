package entity

import (
	"testing"
	"time"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/stretchr/testify/assert"
)

func TestArchiveEvent_CheckCorrectCall(t *testing.T) {

	type args struct {
		deviceID primitive.ID
		clientID primitive.ID
		now      time.Time
	}

	sut := NewArchiveEvent(primitive.NewID(), primitive.NewID(), primitive.NewID())

	tests := []struct {
		name        string
		args        args
		expectedErr error
		assertFunc  assert.ErrorAssertionFunc
	}{
		{
			name: "normal case",
			args: args{
				deviceID: sut.DeviceID,
				clientID: sut.ClientID,
				now:      time.Now(),
			},
			assertFunc: assert.NoError,
		},
		{
			name: "invalid device id",
			args: args{
				deviceID: primitive.NewID(),
				clientID: sut.ClientID,
				now:      time.Now(),
			},
			expectedErr: apperrs.InvalidEventDeviceIDError,
			assertFunc:  assert.Error,
		},
		{
			name: "invalid client id",
			args: args{
				deviceID: sut.DeviceID,
				clientID: primitive.NewID(),
				now:      time.Now(),
			},
			expectedErr: apperrs.InvalidEventClientIDError,
			assertFunc:  assert.Error,
		},
		{
			name: "event time out",
			args: args{
				deviceID: sut.DeviceID,
				clientID: sut.ClientID,
				now:      time.Now().Add(31 * time.Second),
			},
			expectedErr: apperrs.EventTimeOutError,
			assertFunc:  assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := sut.CheckCorrectCall(tt.args.deviceID, tt.args.clientID, tt.args.now)
			if tt.expectedErr != nil {
				tt.assertFunc(t, err, tt.expectedErr.Error())
				return
			}

			tt.assertFunc(t, err)
		})
	}
}
