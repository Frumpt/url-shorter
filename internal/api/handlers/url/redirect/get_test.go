package redirect

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"url-sorter/internal/api"
	"url-sorter/internal/api/handlers/url/redirect/mocks"
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

			testGet := mocks.NewURLGetter(t)
			if tt.errorResp == "" || tt.errorGet != nil {
				testGet.On(
					"GetURL",
					context.Background(),
					&database.GetURLParams{Alias: tt.alias},
				).
					Return(tt.URL, tt.errorGet).
					Once()
			}

			r := chi.NewRouter()
			r.Get("/{alias}", New(imitationLogger, testGet))

			srv := httptest.NewServer(r)
			defer srv.Close()

			redirectToURL, err := api.GetRedirect(srv.URL + "/" + tt.alias)
			require.NoError(t, err)

			require.Equal(t, tt.URL, redirectToURL)
		})
	}
}
