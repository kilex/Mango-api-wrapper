package mango_events

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
)

func EndCall(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	summaryCall := mango_objects.Call{}
	err := json.Unmarshal([]byte(request.Json), &summaryCall)
	if err != nil {
		panic(err)
	}

	summaryCall.CreateTime = time.Unix(summaryCall.CreateTimeStamp, 0)
	summaryCall.ForwardTime = time.Unix(summaryCall.ForwardTimeStamp, 0)
	summaryCall.TalkTime = time.Unix(summaryCall.TalkTimeStamp, 0)
	summaryCall.EndTime = time.Unix(summaryCall.EndTimeStamp, 0)

	Events.AddEndCall(&summaryCall)
}
