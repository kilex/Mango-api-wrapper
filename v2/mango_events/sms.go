package mango_events

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kilex/Mango-api-wrapper/v2/mango_objects"
	"github.com/kilex/Mango-api-wrapper/v2/mango_request"
)

func ResultSMS(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	sms := mango_objects.SMS{}
	err := json.Unmarshal([]byte(request.Json), &sms)
	if err != nil {
		panic(err)
	}

	sms.Time = time.Unix(sms.Timestamp, 0)
	Events.AddSMS(&sms)
}
