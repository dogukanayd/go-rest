package config

import (
	"testing"
)

func TestGetDBDSN(t *testing.T) {
	type args struct {
		dbName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "It should return correct connection DSN",
			args: args{
				dbName: "unitTest",
			},
			want: connections["mysql"] + "/" + "unitTest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDBDSN(tt.args.dbName); got != tt.want {
				t.Errorf("GetDBDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
