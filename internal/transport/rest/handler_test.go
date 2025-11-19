package rest

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest" // Специальный пакет для теста HTTP
	"testing"
)

type MockRepo struct{}

func (m *MockRepo) SaveNumber(ctx context.Context, num int) error {
	return nil
}

func (m *MockRepo) GetSortedNums(ctx context.Context) ([]int, error) {
	return []int{1, 2, 3}, nil
}

func TestHandler_SaveAndGetNumber(t *testing.T) {

	mockRepo := &MockRepo{}
	handler := NewHandler(mockRepo)

	requestBody := []byte(`{"number": 3}`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(requestBody))

	recorder := httptest.NewRecorder()

	handler.SaveAndGetNumber(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"sortedNums":[1,2,3]}`
	if body := recorder.Body.String(); body[0:len(expected)] != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
