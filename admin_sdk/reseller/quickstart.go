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
// [START admin_sdk_reseller_quickstart]
package main

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/reseller/v1"
	"log"
)

func main() {
	ctx := context.Background()
	/* Load pre-authorized user credentials from the environment.
	   TODO(developer) - See https://developers.google.com/identity  and
	     https://cloud.google.com/docs/authentication/production for
	    guides on implementing OAuth2 for your application.
	*/
	client, err := google.DefaultClient(ctx, reseller.AppsOrderScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}

	srv, err := reseller.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve reseller Client %v", err)
	}

	r, err := srv.Subscriptions.List().MaxResults(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve subscriptions. %v", err)
	}

	if len(r.Subscriptions) == 0 {
		fmt.Println("No subscriptions found.")
	} else {
		fmt.Println("Subscriptions:")
		for _, s := range r.Subscriptions {
			fmt.Printf("%s (%s, %s)\n", s.CustomerId, s.SkuId, s.Plan.PlanName)
		}
	}
}

// [END admin_sdk_reseller_quickstart]
