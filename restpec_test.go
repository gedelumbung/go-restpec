package restpec

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestRestpec(t *testing.T) {
	nr := httptest.NewRecorder()
	var response Response

	statusCode := 502
	message := "Bad Gateway"
	restpec(nr, statusCode, message, nil, nil)

	if err := json.Unmarshal(nr.Body.Bytes(), &response); err != nil {
		t.Errorf("not valid JSON: %v", err)
	}

	if nr.Code != statusCode {
		t.Errorf("got %v want %v", nr.Code, statusCode)
	}
}
