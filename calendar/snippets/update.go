package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func updateEvent(client *http.Client) {
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// specify the calendar id for which events will be created
	calendarId := "primary"
	// Retrieve the event from the API
	eventID := "id of the created event"
	createdEvent, err := srv.Events.Get(calendarId, eventID).Do()
	if err != nil {
		log.Fatalf("Unable to get calendar event %v", err)
	}

	// make the updates required
	createdEvent.Summary = "Set the update value"
	updatedEvent, err := srv.Events.Update(calendarId, eventID, createdEvent).
		SendUpdates("all").ConferenceDataVersion(1).Do()
	if err != nil {
		log.Fatalf("Unable to update calendar event %v", err)
	}

	fmt.Printf("Event updated: %s\n", updatedEvent.HtmlLink)
}
