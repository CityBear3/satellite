package mysql

import (
	"database/sql"
	"testing"

	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/testutils/helper"
	"github.com/stretchr/testify/assert"
)

func TestConvertToSqlTx(t *testing.T) {
	type args struct {
		rtx repository.ITx
	}

	db, err := helper.GetTestDB()
	if err != nil {
		t.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		args        args
		want        *sql.Tx
		expectedErr error
	}{
		{
			name: "normal case",
			args: args{
				rtx: tx,
			},
			want: tx,
		},
		{
			name: "invalid type",
			args: args{
				rtx: nil,
			},
			expectedErr: apperrs.UnexpectedError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToSqlTx(tt.args.rtx)
			if tt.expectedErr != nil {
				assert.Error(t, err, tt.expectedErr.Error())
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
