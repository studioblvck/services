package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/event"
)

// Publish a message to the event. Specify a topic to group messages for a specific topic.
func PublishAmessage() {
	eventService := event.NewEventService(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := eventService.Publish(&event.PublishRequest{
		Message: map[string]interface{}{
			"type": "signup",
			"user": "john",
			"id":   "1",
		},
		Topic: "user",
	})
	fmt.Println(rsp, err)
}