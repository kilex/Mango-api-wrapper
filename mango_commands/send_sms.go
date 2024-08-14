package mango_commands

import (
	"github.com/google/uuid"
	"github.com/kilex/Mango-api-wrapper/mango_objects"
	"github.com/kilex/Mango-api-wrapper/mango_request"
)

// Отправка SMS
func SendSms(fromExtension, toNumber, smsSender, text, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	sms := mango_objects.SMS{
		CommandID:     uuidCall,
		Text:          text,
		FromExtension: fromExtension,
		ToNumber:      toNumber,
		SmsSender:     smsSender,
	}

	jsonResp := mango_request.RequestToMango(sms.ToJSON(), "commands/sms", apiKey, apiSing, apiUrl)
	return mango_request.ParseResult(jsonResp)
}
