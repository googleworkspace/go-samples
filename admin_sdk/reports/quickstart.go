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
// [START admin_sdk_reports_quickstart]
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/reports/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	/* Load pre-authorized user credentials from the environment.
	   TODO(developer) - See https://developers.google.com/identity  and
	     https://cloud.google.com/docs/authentication/production for
	    guides on implementing OAuth2 for your application.
	*/
	client, err := google.DefaultClient(ctx, admin.AdminReportsAuditReadonlyScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}

	srv, err := admin.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve reports Client %v", err)
	}

	r, err := srv.Activities.List("all", "login").MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve logins to domain. %v", err)
	}

	if len(r.Items) == 0 {
		fmt.Println("No logins found.")
	} else {
		fmt.Println("Logins:")
		for _, a := range r.Items {
			t, err := time.Parse(time.RFC3339Nano, a.Id.Time)
			if err != nil {
				fmt.Println("Unable to parse login time.")
				// Set time to zero.
				t = time.Time{}
			}
			fmt.Printf("%s: %s %s\n", t.Format(time.RFC822), a.Actor.Email,
				a.Events[0].Name)
		}
	}
}

// [END admin_sdk_reports_quickstart]
