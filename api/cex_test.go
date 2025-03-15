package api_test // pacakge names for test file _test as the file name _test.go
import (
	"testing"

	"example.com/crypto-masters/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error not found")
	}
}
