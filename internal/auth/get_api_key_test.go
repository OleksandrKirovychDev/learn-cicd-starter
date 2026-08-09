package auth

import (
	"testing"
)

func TestGetApiGe(t *testing.T) {
	type test struct {
		name    string
		headers map[string][]string
		want    string
		wantErr bool
	}

	tests := []test{
		{
			name: "valid api key",
			headers: map[string][]string{
				"Authorization": {"ApiKey my-api-key"},
			},
			want: "my-api-key",
			wantErr: false,
		},
		{
			name: "missing authorization header",
			headers: map[string][]string{},
			want: "",
			wantErr: true,
		},
		{
			name: "malformed authorization header",
			headers: map[string][]string{
				"Authorization": {"Bearer my-api-key"},
			},
			want: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := GetAPIKey(tt.headers)

		if (err != nil) != tt.wantErr {
			t.Errorf("%s: GetAPIKey() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}

		if got != tt.want {
			t.Errorf("%s: GetAPIKey() got = %v, want %v", tt.name, got, tt.want)
		}
	}
}