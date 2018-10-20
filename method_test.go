package restpec

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

type User struct {
	ID   int
	Name string
}

func TestOk(t *testing.T) {
	nr := httptest.NewRecorder()
	var response Response

	statusCode := 200
	message := "Well Done"
	data := []User{
		User{
			ID:   1,
			Name: "Gede",
		},
		User{
			ID:   2,
			Name: "Ayu",
		},
	}
	meta := map[string]interface{}{
		"query": "?page=1",
	}
	Ok(nr, message, data, meta)

	if err := json.Unmarshal(nr.Body.Bytes(), &response); err != nil {
		t.Errorf("not valid JSON: %v", err)
	}

	if nr.Code != statusCode {
		t.Errorf("got %v want %v", nr.Code, statusCode)
	}

	dataMarshal, _ := json.Marshal(data)
	responseDataMarshal, _ := json.Marshal(response.Data)

	if !reflect.DeepEqual(response.Meta, meta) ||
		!reflect.DeepEqual(responseDataMarshal, dataMarshal) {
		t.Fail()
	}
}

func TestBadRequest(t *testing.T) {
	nr := httptest.NewRecorder()
	var response Response

	statusCode := 400
	message := "Bad Gateway"
	BadRequest(nr, message, nil, nil)

	if err := json.Unmarshal(nr.Body.Bytes(), &response); err != nil {
		t.Errorf("not valid JSON: %v", err)
	}

	if nr.Code != statusCode {
		t.Errorf("got %v want %v", nr.Code, statusCode)
	}
}
