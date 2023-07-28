package driver

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	TestDataPath = "../testdata/config/%s"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	type test struct {
		name    string
		args    args
		wantErr bool
		want    Config
	}

	if err := os.Setenv("TEST_PASSWORD", "test"); err != nil {
		t.Error(err)
	}

	tests := []test{
		{
			name: "load config file",
			args: args{
				path: fmt.Sprintf(TestDataPath, "test_load_config.yml"),
			},
			wantErr: false,
			want: Config{
				ServerConfig: ServerConfig{
					Host:      "localhost",
					Port:      8081,
					IsDevelop: true,
				},
				DBConfig: DBConfig{
					Driver:   "mysql",
					Host:     "localhost",
					Port:     3307,
					DbName:   "satellite",
					User:     "satellite",
					Password: "test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.path)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
