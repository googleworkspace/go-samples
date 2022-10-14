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
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func insertEvent(client *http.Client) {

	// date time should be in rcf339 format
	calendarEvent := &calendar.Event{
		Summary:     "Google event to learn more about the products",
		Location:    "Mumbai",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2022-10-14T09:00:00-07:00",
			TimeZone: "Asia/Kolkata",
		},
		End: &calendar.EventDateTime{
			DateTime: "2022-10-14T17:00:00-07:00",
			TimeZone: "Asia/Kolkata",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"}, //for reccuring events
		Attendees: []*calendar.EventAttendee{
			{Email: "john@example.com"},
			{Email: "sha@example.com"},
			{Email: "example@example.com"},
		},
		// need to mention ConferenceData -> create request -> requestID -> ConferenceSolutionKey
		// for creating event as google meet
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: "some-string-which-needs-to-be-changed-eachtime-for-every-new-google-meet-link",
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		},
	}

	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	// specify the calendar id for which events will be created
	calendarId := "primary"
	createdEvent, err := srv.Events.Insert(calendarId, calendarEvent).
		SendUpdates("all").ConferenceDataVersion(1).Do()
	if err != nil {
		log.Fatalf("Unable to create calendar event %v", err)
	}

	fmt.Printf("Event created: %s\n", createdEvent.HtmlLink)
}

func main() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	insertEvent(client)
	// patchEvent(client)
	// deleteEvent(client)
	// updateEvent(client)
}
