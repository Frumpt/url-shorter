package save

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url-sorter/internal/api/handlers/url/save/mocks"
	"url-sorter/internal/storage/database"
)

func TestNew(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	var tests = []struct {
		name      string
		id        int32
		url       string
		alias     string
		errorResp string
	}{
		{
			name:      "Success with alias",
			id:        1,
			url:       "https://google.com",
			alias:     "g",
			errorResp: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testSave := mocks.NewURLSaver(t)
			if tt.errorResp == "" {
				testSave.On(
					"SaveURL",
					context.Background(),
					&database.SaveURLParams{ID: tt.id, Alias: tt.alias, Url: tt.url},
				).
					Return(&database.Url{ID: tt.id, Alias: tt.alias, Url: tt.url}, nil).
					Once()
			}

			handler := New(logger, testSave)
			input, err := json.Marshal(Request{ID: tt.id, Url: tt.url, Alias: tt.alias})
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader(input))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, rr.Code, http.StatusOK)

			bodyResp := rr.Body.String()

			var resp Response

			require.NoError(t, json.Unmarshal([]byte(bodyResp), &resp))

			require.Equal(t, tt.errorResp, resp.Error)

		})
	}
}
