package apis

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dlbarduzzi/sentinel/tests"
)

type apiTestScenario struct {
	name            string
	url             string
	method          string
	body            io.Reader
	expectedStatus  int
	expectedContent []string
}

func (s *apiTestScenario) Test(t *testing.T) {
	t.Run(s.name, func(t *testing.T) {
		s.test(t)
	})
}

func (s *apiTestScenario) test(t *testing.T) {
	app, err := tests.NewTestApp()
	if err != nil {
		t.Fatalf("failed to initialize test app instance - %v", err)
	}

	router := newRouter(app)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(s.method, s.url, s.body)

	// Set default header.
	req.Header.Set("Content-Type", "application/json")

	mux := router.buildMux()
	mux.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != s.expectedStatus {
		t.Fatalf(
			"expected status code to be %d, got %d",
			s.expectedStatus, res.StatusCode,
		)
	}

	if len(s.expectedContent) == 0 {
		if len(rec.Body.Bytes()) != 0 {
			t.Fatalf(
				"expected empty content, got \n%v",
				rec.Body.String(),
			)
		}
	} else {
		var body string
		buf := new(bytes.Buffer)

		err := json.Compact(buf, rec.Body.Bytes())
		if err != nil {
			// Not a json payload.
			body = rec.Body.String()
		} else {
			// A valid json payload.
			body = buf.String()
		}

		for _, content := range s.expectedContent {
			if !strings.Contains(body, content) {
				t.Errorf(
					"expected content %v in response body \n%v",
					content, body,
				)
			}
		}
	}
}
