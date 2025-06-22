package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey(t *testing.T) {
	tests := map[string]struct {
		header http.Header
		want   string
		result bool
	}{
		"valid header": {
			header: http.Header{"Authorization": []string{"ApiKey the_api_key"}},
			want:   "the_api_key",
			result: true,
		},
		"invalid header": {
			header: http.Header{"Authorization": []string{"ApiKeyV9 the_api_key"}},
			want:   "",
			result: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got, err := GetAPIKey(test.header); (err == nil) != test.result {
				t.Errorf("expected: %q; got: %q", test.want, got)
			}
		})
	}
}
