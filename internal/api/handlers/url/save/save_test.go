package save

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-sorter/internal/api/handlers/url/save/mocks"
	"url-sorter/internal/logger"
	"url-sorter/internal/storage/database"
)

func TestNew(t *testing.T) {
	imitationLogger := logger.NewImitationLogger()

	var tests = []struct {
		name      string
		id        int32
		url       string
		alias     string
		errorResp string
		errorSave error
	}{
		{
			name:      "Success with alias",
			id:        1,
			url:       "https://google.com",
			alias:     "g",
			errorResp: "",
		},
		{
			name:      "Success without alias",
			id:        1,
			url:       "https://google.com",
			alias:     "",
			errorResp: "",
		},
		{
			name:      "Fail URL value",
			id:        1,
			url:       "gowqreqwem",
			alias:     "",
			errorResp: "validation error: failed Url is a value URL",
		},
		{
			name:      "Fail empty URL",
			id:        1,
			url:       "",
			alias:     "",
			errorResp: "validation error: failed Url is a required field",
		},
		{
			name:      "Fail empty ID",
			id:        0,
			url:       "https://google.com",
			alias:     "",
			errorResp: "validation error: failed ID is a required field",
		},
		{
			name:      "Fail empty URL and empty ID",
			id:        0,
			url:       "",
			alias:     "",
			errorResp: "validation error: failed ID is a required field, validation error: failed Url is a required field",
		},
		{
			name:      "Fail some error in SaveURL",
			id:        1,
			url:       "https://google.com",
			alias:     "",
			errorResp: "some error",
			errorSave: errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testSave := mocks.NewURLSaver(t)
			if tt.errorResp == "" || tt.errorSave != nil {
				testSave.On(
					"SaveURL",
					context.Background(),
					&database.SaveURLParams{ID: tt.id, Alias: tt.alias, Url: tt.url},
				).
					Return(&database.Url{ID: tt.id, Alias: tt.alias, Url: tt.url}, tt.errorSave).
					Once()
			}

			handler := New(imitationLogger, testSave)
			input, err := json.Marshal(Request{ID: tt.id, Url: tt.url, Alias: tt.alias})
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader(input))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// TODO: Why only 200?
			require.Equal(t, rr.Code, http.StatusOK)

			bodyResp := rr.Body.String()

			var resp Response

			require.NoError(t, json.Unmarshal([]byte(bodyResp), &resp))

			require.Equal(t, tt.errorResp, resp.Error)

		})
	}
}
