package validations

import (
	"fmt"
	"github.com/CityBear3/satellite/pb/image/v1"
	"github.com/CityBear3/satellite/pkg/apperrs"
	"github.com/CityBear3/satellite/tests/helper"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestDataPath = "../../../testdata/validation/%s"
)

func TestValidateUploadImageStream(t *testing.T) {
	type args struct {
		meta        *imagePb.Meta
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
				&imagePb.Meta{Id: ulid.Make().String()},
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
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidMetaInfoMsg),
		},
		{
			name: "data is not exist on request",
			args: args{
				&imagePb.Meta{Id: ulid.Make().String()},
				[]byte{},
				"image/jpeg",
			},
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidDataMsg),
		},
		{
			name: "invalid file extension.",
			args: args{
				&imagePb.Meta{Id: ulid.Make().String()},
				invalidData,
				"text/plain",
			},
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidFileExtMsg),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUploadImageStream(tt.args.meta, tt.args.data, tt.args.contentType)
			if tt.want != nil {
				assert.EqualError(t, err, tt.want.Error())
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}
