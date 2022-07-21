// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
)

// A group of Google services.
type Services struct {
	Drive  *drive.Service
	Slides *slides.Service
	Sheets *sheets.Service
}

// Gets Google services authenticated with Drive, Slides, and Sheets.
func getServices() *Services {
	ctx := context.Background()
	// Uses env GOOGLE_APPLICATION_CREDENTIALS
	client, err := google.DefaultClient(ctx,
		drive.DriveScope,
		slides.PresentationsScope,
		sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}
	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Drive client: %v", err)
	}
	slidesService, err := slides.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Slides client: %v", err)
	}
	sheetsService, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Sheets client: %v", err)
	}
	return &Services{
		Drive:  driveService,
		Slides: slidesService,
		Sheets: sheetsService,
	}
}
