package magicbell

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

var (
	validConfig = Config{
		APIKey:    "key",
		APISecret: "secret",
	}
)
