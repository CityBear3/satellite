package env

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func setValue(key string, value string) error {
	if err := os.Setenv(key, value); err != nil {
		return err
	}
	return nil
}

func TestGetStrEnv(t *testing.T) {
	if err := setValue("SOME_KEY", "test"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key          string
		defaultValue string
		required     bool
	}
	tests := []struct {
		name        string
		args        args
		want        string
		wantErr     bool
		expectedErr error
	}{
		// Test cases
		{
			name: "env variable exists",
			args: args{
				key:      "SOME_KEY",
				required: true,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "env variable does not exist",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: true,
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
		{
			name: "env variable does not exist and is not required",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: "default",
				required:     false,
			},
			want:    "default",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStrEnv(tt.args.key, tt.args.defaultValue, tt.args.required)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetIntEnv(t *testing.T) {
	if err := setValue("SOME_KEY", strconv.Itoa(123)); err != nil {
		t.Fatal(err)
	}
	if err := setValue("INVALID_EXISTENT_KEY", "invalid"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key          string
		defaultValue int
		required     bool
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		expectedErr error
	}{
		{
			name: "env variable exists",
			args: args{
				key:      "SOME_KEY",
				required: true,
			},
			want: 123,
		},
		{
			name: "env variable does not exist",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: true,
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
		{
			name: "env variable does not exist and is not required",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: 1,
				required:     false,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key:      "INVALID_EXISTENT_KEY",
				required: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIntEnv(tt.args.key, tt.args.defaultValue, tt.args.required)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.EqualError(t, err, tt.expectedErr.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetBoolEnv(t *testing.T) {
	if err := setValue("SOME_KEY", strconv.FormatBool(true)); err != nil {
		t.Fatal(err)
	}
	if err := setValue("INVALID_EXISTENT_KEY", "invalid"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key          string
		defaultValue bool
		required     bool
	}
	tests := []struct {
		name        string
		args        args
		want        bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "env variable exists",
			args: args{
				key:      "SOME_KEY",
				required: true,
			},
			want: true,
		},
		{
			name: "env variable does not exist",
			args: args{
				key:      "NON_EXISTENT_KEY",
				required: true,
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
		{
			name: "env variable does not exist and is not required",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: true,
				required:     false,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key:      "INVALID_EXISTENT_KEY",
				required: false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBoolEnv(tt.args.key, tt.args.defaultValue, tt.args.required)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					assert.EqualError(t, err, tt.expectedErr.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}