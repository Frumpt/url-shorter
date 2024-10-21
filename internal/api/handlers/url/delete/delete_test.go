package redirect

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-sorter/internal/api"
	"url-sorter/internal/api/handlers/url/delete/mocks"
	"url-sorter/internal/logger"
	"url-sorter/internal/storage/database"
)

func TestNew(t *testing.T) {
	imitationLogger := logger.NewImitationLogger()

	var tests = []struct {
		name      string
		URL       string
		alias     string
		errorResp string
		errorGet  error
	}{
		{
			name:  "Success redirect to URL",
			alias: "url",
			URL:   "https://google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testDelete := mocks.NewURLDelete(t)
			if tt.errorResp == "" || tt.errorGet != nil {
				testDelete.On(
					"DeleteURL",
					context.Background(),
					&database.DeleteURLParams{Alias: tt.alias},
				).
					Return(tt.errorGet).
					Once()
			}

			r := chi.NewRouter()
			r.Delete("/{alias}", New(imitationLogger, testDelete))

			srv := httptest.NewServer(r)
			defer srv.Close()

			ResponseCode, err := api.DeleteConfirm(srv.URL + "/" + tt.alias)
			require.NoError(t, err)

			require.Equal(t, http.StatusOK, ResponseCode)
		})
	}
}
