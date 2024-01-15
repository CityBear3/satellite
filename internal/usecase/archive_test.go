package usecase_test

import (
	"context"
	"testing"

	"github.com/CityBear3/satellite/internal/adaptor/repository/mock"
	"github.com/CityBear3/satellite/internal/domain/entity"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/primitive/archive"
	"github.com/CityBear3/satellite/internal/domain/primitive/device"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/internal/usecase/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestArchiveInteractor_CreateArchive(t *testing.T) {
	type args struct {
		ctx    context.Context
		req    usecase.CreateArchiveInput
		device entity.Device
	}

	ctrl := gomock.NewController(t)
	mockArchiveRepository := mock_repository.NewMockIArchiveRepository(ctrl)
	mockEventRepository := mock_repository.NewMockIEventRepository(ctrl)
	mockITxManager := mock_usecase.NewMockITxManager(ctrl)

	sut := usecase.NewArchiveInteractor(mockArchiveRepository, mockEventRepository, mockITxManager)

	contentType, err := archive.NewContentType("image/png")
	if err != nil {
		t.Fatal(err)
	}

	data, err := archive.NewData([]byte("test"))
	if err != nil {
		t.Fatal(err)
	}

	deviceName, err := device.NewDeviceName("test")
	if err != nil {
		t.Fatal(err)
	}

	device := entity.NewDevice(primitive.NewID(), deviceName, nil, primitive.NewID())

	tests := []struct {
		name      string
		args      args
		wantErr   error
		assertion assert.ErrorAssertionFunc
		runMock   func(ctx context.Context, req usecase.CreateArchiveInput, device entity.Device)
	}{
		{
			name: "success to create archive",
			args: args{
				ctx: context.Background(),
				req: usecase.CreateArchiveInput{
					ArchiveEventID: primitive.NewID(),
					ContentType:    contentType,
					Data:           data,
				},
				device: device,
			},

			assertion: assert.NoError,
			runMock: func(ctx context.Context, req usecase.CreateArchiveInput, device entity.Device) {
				mockITxManager.
					EXPECT().
					DoInTx(ctx, gomock.Any()).
					DoAndReturn(func(ctx context.Context, f func(ctx2 context.Context) error) (interface{}, error) {
						return gomock.Any(), f(ctx)
					})

				mockEventRepository.
					EXPECT().
					GetArchiveEvent(ctx, req.ArchiveEventID).
					Return(entity.NewArchiveEvent(req.ArchiveEventID, device.ID, device.ClientID), nil)

				mockArchiveRepository.
					EXPECT().Save(ctx, gomock.Any()).
					Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.runMock(tt.args.ctx, tt.args.req, tt.args.device)
			err = sut.CreateArchive(tt.args.ctx, tt.args.req, tt.args.device)
			tt.assertion(t, err)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
