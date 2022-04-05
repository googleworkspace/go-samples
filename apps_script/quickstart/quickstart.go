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
// [START apps_script_api_quickstart]
package main

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/script/v1"
	"log"
)

func main() {
	ctx := context.Background()
	/* Load pre-authorized user credentials from the environment.
	   TODO(developer) - See https://developers.google.com/identity  and
	     https://cloud.google.com/docs/authentication/production for
	    guides on implementing OAuth2 for your application.
	*/
	client, err := google.DefaultClient(ctx, script.ScriptProjectsScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}

	srv, err := script.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Script client: %v", err)
	}

	req := script.CreateProjectRequest{Title: "My Script"}
	createRes, err := srv.Projects.Create(&req).Do()
	if err != nil {
		// The API encountered a problem.
		log.Fatalf("The API returned an error: %v", err)
	}
	content := &script.Content{
		ScriptId: createRes.ScriptId,
		Files: []*script.File{{
			Name:   "hello",
			Type:   "SERVER_JS",
			Source: "function helloWorld() {\n  console.log('Hello, world!');}",
		}, {
			Name: "appsscript",
			Type: "JSON",
			Source: "{\"timeZone\":\"America/New_York\",\"exceptionLogging\":" +
				"\"CLOUD\"}",
		}},
	}
	updateContentRes, err := srv.Projects.UpdateContent(createRes.ScriptId,
		content).Do()
	if err != nil {
		// The API encountered a problem.
		log.Fatalf("The API returned an error: %v", err)
	}
	log.Printf("https://script.google.com/d/%v/edit", updateContentRes.ScriptId)
}

// [END apps_script_api_quickstart]
