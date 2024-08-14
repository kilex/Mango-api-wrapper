package mango_events

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kilex/Mango-api-wrapper/v2/mango_objects"
	"github.com/kilex/Mango-api-wrapper/v2/mango_request"
)

func PressDTMF(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	dtmf := mango_objects.DTMF{}
	err := json.Unmarshal([]byte(request.Json), &dtmf)
	if err != nil {
		panic(err)
	}

	dtmf.Time = time.Unix(dtmf.Timestamp, 0)

	Events.AddDTMF(&dtmf)
}
