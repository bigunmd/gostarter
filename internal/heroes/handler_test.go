//go:build unit

package heroes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bigunmd/gostarter/pkg/util/tests"
	"github.com/stretchr/testify/require"
)

func setupTestMux(ctx context.Context) *http.ServeMux {
	svc := setupTestService(ctx)

	mux := http.NewServeMux()
	mux.Handle("GET /healthz", HandleHealthz())
	mux.Handle("POST /v1/heroes", HandleCreateHero(svc))

	return mux
}

func TestHandler(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx = tests.SetupTestLogger(t).WithContext(ctx)

	mux := setupTestMux(ctx)

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name            string
		args            func(t *testing.T) args
		expectedCode    int
		expectedBody    string
		expectedBodyErr string
		expectedHeaders map[string]string
	}{
		{
			name: "must return http.StatusOK, empty body",
			args: func(t *testing.T) args {
				req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
				require.NoError(t, err)

				return args{
					req: req,
				}
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "must return http.StatusCreated, empty body and Location header with new resource",
			args: func(t *testing.T) args {
				h := Hero{
					Name: "test",
				}
				b, err := json.Marshal(h)
				require.NoError(t, err)
				req, err := http.NewRequest(http.MethodPost, "/v1/heroes", bytes.NewReader(b))
				require.NoError(t, err)

				return args{
					req: req,
				}
			},
			expectedCode: http.StatusCreated,
			expectedHeaders: map[string]string{
				"Location": "/v1/heroes/test",
			},
		},
		{
			name: "must return http.StatusConfilict and empty body",
			args: func(t *testing.T) args {
				h := Hero{
					Name: "test",
				}
				b, err := json.Marshal(h)
				require.NoError(t, err)
				req, err := http.NewRequest(http.MethodPost, "/v1/heroes", bytes.NewReader(b))
				require.NoError(t, err)

				return args{
					req: req,
				}
			},
			expectedCode:    http.StatusConflict,
			expectedBodyErr: ErrAlreadyExists.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			resp := httptest.NewRecorder()
			mux.ServeHTTP(resp, tArgs.req)

			require.Equal(t, tt.expectedCode, resp.Result().StatusCode)
			if tt.expectedBodyErr != "" {
				require.Contains(t, resp.Body.String(), tt.expectedBodyErr)
			} else {
				require.Equal(t, tt.expectedBody, resp.Body.String())
			}
			for k, v := range tt.expectedHeaders {
				require.Equal(t, v, resp.Result().Header.Get(k))
			}
		})
	}
}
