package util

import "encoding/json"

// parseJson parses the JSON-encoded data and stores the result in the value pointed to by result.
// interface{} here to handle different types like T
// unmarshal parse the json and put it into the second argument type
func ParseJson(jsonBodyBytes []byte, result interface{}) error {
	return json.Unmarshal(jsonBodyBytes, result)
}
