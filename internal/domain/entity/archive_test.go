package entity

import (
	"testing"

	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/stretchr/testify/assert"
)

func TestArchive_CheckCorrectCall(t *testing.T) {
	type args struct {
		client Client
	}

	contentType, err := primitive.NewContentType("image/jpeg")
	if err != nil {
		t.Fatal(err)
	}

	clientName, err := primitive.NewClientName("test")
	if err != nil {
		t.Fatal(err)
	}

	deviceName, err := primitive.NewDeviceName("test")
	if err != nil {
		t.Fatal(err)
	}

	archive := NewArchive(primitive.NewID(), primitive.NewID(), contentType, primitive.NewID())
	device := NewDevice(archive.DeviceID, deviceName, nil, primitive.NewID())

	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "normal case",
			args: args{
				client: NewClient(device.ClientID, clientName, nil, []Device{device}),
			},
			expectedErr: nil,
		},
		{
			name: "invalid case",
			args: args{
				client: NewClient(device.ClientID, clientName, nil, []Device{}),
			},
			expectedErr: apperrs.InvalidClientCallingArchiveError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := archive.CheckCorrectCall(tt.args.client)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}
