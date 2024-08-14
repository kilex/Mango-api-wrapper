package mango_commands

import (
	"github.com/google/uuid"
	"github.com/kilex/Mango-api-wrapper/mango_objects"
	"github.com/kilex/Mango-api-wrapper/mango_request"
)

// Включение записи разговора
func TurnOnRecordingCall(callID, callPartyNumber, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID:       uuidCall,
		CallID:          callID,
		CallPartyNumber: callPartyNumber,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/recording/start", apiKey, apiSing, apiUrl)
	return mango_request.ParseResult(jsonResp)
}
