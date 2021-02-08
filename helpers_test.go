package magicbell

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func runGlobalTest(config Config, fn func()) {
	defer func() { api = nil }()
	Init(config)
	fn()
}

func runServer(t *testing.T, path string, method string, status int, fn func(config Config)) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path, r.URL.Path)
		assert.Equal(t, method, r.Method)

		dataPath := "testdata/api/500.txt"
		if status != 500 {
			dataPath = fmt.Sprintf("testdata/api%s_%s_%d.json", strings.ToLower(path), strings.ToLower(method), status)
		}

		data, err := ioutil.ReadFile(dataPath)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(status)
		if _, err := w.Write(data); err != nil {
			panic(err)
		}
	}))
	defer srv.Close()

	fn(validConfig.withBaseURL(srv.URL))
}

func assertAPIError(code APIErrorCode, message string) func(*testing.T, error) {
	return func(t *testing.T, err error) {
		require.Error(t, err)
		assert.True(t, IsAPIErrors(err))

		errs := err.(APIErrors)
		assert.Len(t, errs, 1)
		assert.Equal(t, APIError{
			Code:    code,
			Message: message,
		}, errs[0])
	}
}

func assertInternalServerError(t *testing.T, err error) {
	require.Error(t, err)

	assert.True(t, IsInternalServerError(err))
	assert.Equal(t, InternalServerError{
		StatusCode: 500,
		Body:       "Internal server error\n",
	}, err.(InternalServerError))
}

func assertNoError(t *testing.T, err error) {
	require.NoError(t, err)
}

var (
	validConfig = Config{
		APIKey:    "key",
		APISecret: "secret",
	}
)
