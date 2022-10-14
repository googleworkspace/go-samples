/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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
