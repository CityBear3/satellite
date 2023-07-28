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
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Test cases
		{
			name: "env variable exists",
			args: args{
				key: "SOME_KEY",
			},
			want: "test",
		},
		{
			name: "env variable does not exist and return default value",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: "default",
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetStrEnv(tt.args.key, tt.args.defaultValue)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetRequiredStrEnv(t *testing.T) {
	if err := setValue("SOME_KEY", "test"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key string
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
				key: "SOME_KEY",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "env variable does not exist",
			args: args{
				key: "NON_EXISTENT_KEY",
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRequiredStrEnv(tt.args.key)
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
				key: "SOME_KEY",
			},
			want: 123,
		},
		{
			name: "env variable does not exist and return default value",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key: "INVALID_EXISTENT_KEY",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIntEnv(tt.args.key, tt.args.defaultValue)
			if tt.wantErr {
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

func TestGetRequiredIntEnv(t *testing.T) {
	if err := setValue("SOME_KEY", strconv.Itoa(123)); err != nil {
		t.Fatal(err)
	}
	if err := setValue("INVALID_EXISTENT_KEY", ""); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key string
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
				key: "SOME_KEY",
			},
			want: 123,
		},
		{
			name: "env variable does not exist",
			args: args{
				key: "NON_EXISTENT_KEY",
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key: "INVALID_EXISTENT_KEY",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRequiredIntEnv(tt.args.key)
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
				key: "SOME_KEY",
			},
			want: true,
		},
		{
			name: "env variable does not exist and return default value",
			args: args{
				key:          "NON_EXISTENT_KEY",
				defaultValue: false,
			},
			want: false,
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key: "INVALID_EXISTENT_KEY",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBoolEnv(tt.args.key, tt.args.defaultValue)
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

func TestGetRequiredBoolEnv(t *testing.T) {
	if err := setValue("SOME_KEY", strconv.FormatBool(true)); err != nil {
		t.Fatal(err)
	}
	if err := setValue("INVALID_EXISTENT_KEY", "invalid"); err != nil {
		t.Fatal(err)
	}
	type args struct {
		key string
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
				key: "SOME_KEY",
			},
			want: true,
		},
		{
			name: "env variable does not exist",
			args: args{
				key: "NON_EXISTENT_KEY",
			},
			wantErr:     true,
			expectedErr: errors.New(fmt.Sprintf(ValueNotFoundError, "NON_EXISTENT_KEY")),
		},
		{
			name: "env variable exist and is invalid",
			args: args{
				key: "INVALID_EXISTENT_KEY",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRequiredBoolEnv(tt.args.key)
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
