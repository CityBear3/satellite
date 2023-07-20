package model

import (
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImage_checkDataSize(t *testing.T) {
	const limit = 1024
	type args struct {
		archive *Archive
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "allowed size",
			args: args{
				archive: NewImageModel(ulid.Make(), "jpeg", 1024, ulid.Make(), []byte{}),
			},
			want: true,
		},
		{
			name: "not allowed size",
			args: args{
				archive: NewImageModel(ulid.Make(), "jpg", 2048, ulid.Make(), []byte{}),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.archive.CheckDataSize(limit)
			assert.Equal(t, tt.want, got)
		})
	}
}
