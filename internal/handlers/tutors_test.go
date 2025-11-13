package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListTutors(t *testing.T) {
	rec := httptest.NewRecorder()
	ListTutors(rec, httptest.NewRequest("GET", "/api/tutors", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
