package mango_request

import (
	"encoding/json"

	"github.com/kilex/Mango-api-wrapper/v2/mango_objects"
)

func ParseResult(jsonResp string) *mango_objects.Result {
	result := mango_objects.Result{}
	err := json.Unmarshal([]byte(jsonResp), &result)
	if err != nil {
		panic(err)
	}

	return &result
}
