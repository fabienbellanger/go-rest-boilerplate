package lib

import (
	"testing"
)

// TestGetHTTPInternalServerError : Test de la fonction GetHTTPResponse
func TestGetHTTPResponse(t *testing.T) {
	response := GetHTTPResponse(200, "Success", nil)
	responseValid := HTTPResponse{200, "Success", nil}

	if response != responseValid {
		t.Errorf("GetHTTPResponse - got %+v: , want: %+v.", response, responseValid)
	}

	response = GetHTTPResponse(404, "Not found", 15)
	responseValid = HTTPResponse{404, "Not found", 15}

	if response != responseValid {
		t.Errorf("GetHTTPResponse - got %+v: , want: %+v.", response, responseValid)
	}
}

// TestGetHTTPInternalServerError : Test de la fonction GetHTTPInternalServerError
func TestGetHTTPInternalServerError(t *testing.T) {
	response := GetHTTPInternalServerError("Database Error")
	responseValid := HTTPResponse{500, "Database Error", nil}

	if response != responseValid {
		t.Errorf("GetHTTPInternalServerError - got %+v: , want: %+v.", response, responseValid)
	}

	response = GetHTTPInternalServerError("")
	responseValid = HTTPResponse{500, "Internal Server Error", nil}

	if response != responseValid {
		t.Errorf("GetHTTPInternalServerError - got %+v: , want: %+v.", response, responseValid)
	}
}
