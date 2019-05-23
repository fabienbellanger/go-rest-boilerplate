package lib

import (
	"testing"
)

// TestGetHTTPInternalServerError
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

// BenchmarkGetHTTPResponse
func BenchmarkGetHTTPResponse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetHTTPResponse(200, "Success", nil)
	}
}

// TestGetHTTPInternalServerError
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
