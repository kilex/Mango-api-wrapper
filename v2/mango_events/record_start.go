package mango_events

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kilex/Mango-api-wrapper/mango_objects"
	"github.com/kilex/Mango-api-wrapper/mango_request"
)

func RecordStart(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	record := mango_objects.Record{}
	err := json.Unmarshal([]byte(request.Json), &record)
	if err != nil {
		panic(err)
	}

	record.Time = time.Unix(record.Timestamp, 0)
	Events.AddStartRecord(&record)
}
