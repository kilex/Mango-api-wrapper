package mango_commands

import (
	"github.com/google/uuid"
	"github.com/kilex/Mango-api-wrapper/mango_objects"
	"github.com/kilex/Mango-api-wrapper/mango_request"
)

// Перевод вызова
func TransferCall(callID, method, toNumber, initiator, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID: uuidCall,
		CallID:    callID,
		ToNumber:  toNumber,
		Method:    method,
		Initiator: initiator,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/transfer", apiKey, apiSing, apiUrl)
	return mango_request.ParseResult(jsonResp)
}
