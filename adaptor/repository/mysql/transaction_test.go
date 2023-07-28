package mysql

import (
	"context"
	"database/sql"
	"github.com/CityBear3/satellite/tests/helper"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"reflect"
	"testing"
)

const errMsg = "Database transaction was not found."

func TestGetTransaction(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type test struct {
		name    string
		args    args
		wantErr bool
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
	boil.SetDB(db)

	ctx := context.Background()
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := []test{
		{
			name: "get transaction",
			args: args{
				ctx: context.WithValue(ctx, "tx", tx),
			},
			wantErr: false,
		},
		{
			name: "tx in ctx is invalid type",
			args: args{
				ctx: context.WithValue(ctx, "tx", "test"),
			},
			wantErr: true,
		},
		{
			name: "tx doesn't exist in ctx",
			args: args{
				ctx: ctx,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get, err := GetTransaction(tt.args.ctx)
			if !tt.wantErr {
				assert.Nil(t, err)
				assert.Equal(t, reflect.TypeOf(tx), reflect.TypeOf(get))
			} else {
				assert.Error(t, err, errMsg)
			}
		})
	}
}
