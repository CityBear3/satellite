package validations

import (
	"fmt"
	"testing"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/pb/archive/v1"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

const (
	TestDataPath = "../../../../testdata/validation/%s"
)

func TestValidateCreateArchive(t *testing.T) {
	type args struct {
		meta        *archivePb.CreateArchiveMetaInfo
		data        []byte
		contentType string
	}

	data := make([]byte, 1024)
	if err := helper.ReadFileData(fmt.Sprintf(TestDataPath, "ok.jpg"), data); err != nil {
		t.Fatal(err)
	}

	invalidData := make([]byte, 1024)
	if err := helper.ReadFileData(fmt.Sprintf(TestDataPath, "invalid.txt"), invalidData); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "no error occurred",
			args: args{
				&archivePb.CreateArchiveMetaInfo{ArchiveEventId: ulid.Make().String()},
				data,
				"image/jpeg",
			},
			want: nil,
		},
		{
			name: "meta info is not exist on request",
			args: args{
				nil,
				data,
				"image/jpeg",
			},
			want: apperrs.InvalidMetaInfoError,
		},
		{
			name: "data is not exist on request",
			args: args{
				&archivePb.CreateArchiveMetaInfo{ArchiveEventId: ulid.Make().String()},
				[]byte{},
				"image/jpeg",
			},
			want: apperrs.InvalidFileSizeError,
		},
		{
			name: "invalid data size",
			args: args{
				&archivePb.CreateArchiveMetaInfo{ArchiveEventId: ulid.Make().String()},
				[]byte{},
				"image/jpeg",
			},
			want: apperrs.InvalidFileSizeError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateArchive(tt.args.meta, tt.args.data)
			if tt.want != nil {
				assert.EqualError(t, err, tt.want.Error())
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}
