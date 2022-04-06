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

	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
)

func getCourse(client *http.Client) {
	// [START classroom_get_course]
	ctx := context.Background()
	srv, err := classroom.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	id := "123456"
	course, err := srv.Courses.Get(id).Do()
	if err != nil {
		log.Fatalf("Course unable to be retrieved %v", err)
	}
	fmt.Printf("Course with ID %v found.", course.Id)
	// [END classroom_get_course]
}

func main() {
	ctx := context.Background()
	/* Load pre-authorized user credentials from the environment.
	   TODO(developer) - See https://developers.google.com/identity  and
	     https://cloud.google.com/docs/authentication/production for
	    guides on implementing OAuth2 for your application.
	*/
	client, err := google.DefaultClient(ctx, classroom.ClassroomCoursesScope)
	if err != nil {
		log.Fatalf("Failed Default authentication: %v", err)
	}
	getCourse(client)
}
