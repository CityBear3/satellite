package validations

import (
	"fmt"
	"github.com/CityBear3/satellite/pb/image/v1"
	"github.com/CityBear3/satellite/pkg/apperrs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	TestDataPath = "../../../testdata"
)

func getImage(name string, buf []byte) error {
	path := fmt.Sprintf("%s/%s", TestDataPath, name)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	_, err = file.Read(buf)
	if err != nil {
		return err
	}

	return nil
}

func TestValidateUploadImageStream(t *testing.T) {
	type args struct {
		meta *image.Meta
		data []byte
	}

	data := make([]byte, 1024)
	if err := getImage("validation.jpg", data); err != nil {
		t.Fatal(err)
	}

	invalidData := make([]byte, 1024)
	if err := getImage("validation_invalid.txt", invalidData); err != nil {
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
				&image.Meta{Id: uuid.New().String()},
				data,
			},
			want: nil,
		},
		{
			name: "meta info is not exist on request",
			args: args{
				nil,
				data,
			},
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidMetaInfoMsg),
		},
		{
			name: "data is not exist on request",
			args: args{
				&image.Meta{Id: uuid.New().String()},
				[]byte{},
			},
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidDataMsg),
		},
		{
			name: "invalid file extension.",
			args: args{
				&image.Meta{Id: uuid.New().String()},
				invalidData,
			},
			want: apperrs.NewError(apperrs.BadRequest, apperrs.InvalidFileExtMsg),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUploadImageStream(tt.args.meta, tt.args.data)
			if tt.want != nil {
				assert.EqualError(t, err, tt.want.Error())
			} else {
				assert.Equal(t, nil, err)
			}
		})
	}
}
