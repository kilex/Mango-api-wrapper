package mango_commands

import (
	"github.com/google/uuid"
	"github.com/kilex/Mango-api-wrapper/v2/mango_objects"
	"github.com/kilex/Mango-api-wrapper/v2/mango_request"
)

// Маршрутизация вызова
func RoutingCall(callID, toNumber, sipHeaders, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID:  uuidCall,
		CallID:     callID,
		ToNumber:   toNumber,
		SipHeaders: sipHeaders,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/route", apiKey, apiSing, apiUrl)
	return mango_request.ParseResult(jsonResp)
}
