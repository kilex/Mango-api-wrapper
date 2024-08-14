package mango_client

import (
	"encoding/json"

	"github.com/kilex/Mango-api-wrapper/v2/mango_objects"
	"github.com/kilex/Mango-api-wrapper/v2/mango_request"
)

func GetUsers(apiKey, apiSing, apiUrl string) *mango_objects.Users {

	jsonResp := mango_request.RequestToMango("{}", "config/users/request", apiKey, apiSing, apiUrl)

	users := mango_objects.Users{}

	err := json.Unmarshal([]byte(jsonResp), &users)
	if err != nil {
		panic(err)
	}

	return &users
}
