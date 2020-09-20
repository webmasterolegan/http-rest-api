package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHellow(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hellow", nil)
	s.handleHellow().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hellow")
}
