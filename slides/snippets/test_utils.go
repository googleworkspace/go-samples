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
	"google.golang.org/api/slides/v1"
	"log"
)

func deleteFileOnCleanup(id string) {
	getServices().Drive.Files.Delete(id)
}

func createTestPresentation() string {
	slidesService := getServices().Slides
	presentation := &slides.Presentation{
		Title: "Test Presentation",
	}
	presentationCreateCall, _ := slidesService.Presentations.Create(presentation).Fields(
		"presentationId",
	).Do()
	return presentationCreateCall.PresentationId
}

func createTestSlide(presentationId string) string {
	requests := []*slides.Request{{
		CreateSlide: &slides.CreateSlideRequest{
			ObjectId:       "TestSlide",
			InsertionIndex: 0,
			SlideLayoutReference: &slides.LayoutReference{
				PredefinedLayout: "BLANK",
			},
		},
	}}

	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	response, err := getServices().Slides.Presentations.BatchUpdate(presentationId, body).Do()
	if err != nil {
		log.Fatalf("Unable to create test slide. %v", err)
	}
	return response.Replies[0].CreateSlide.ObjectId
}

func createTestTextbox(presentationId string, pageId string) string {
	slidesService := getServices().Slides
	boxId := "MyTextBox_01"
	pt350 := slides.Dimension{
		Magnitude: 350,
		Unit:      "PT",
	}
	requests := []*slides.Request{{
		CreateShape: &slides.CreateShapeRequest{
			ObjectId:  boxId,
			ShapeType: "TEXT_BOX",
			ElementProperties: &slides.PageElementProperties{
				PageObjectId: pageId,
				Size: &slides.Size{
					Height: &pt350,
					Width:  &pt350,
				},
				Transform: &slides.AffineTransform{
					ScaleX:     1,
					ScaleY:     1,
					TranslateX: 350,
					TranslateY: 350,
					Unit:       "PT",
				},
			},
		},
	}, {
		InsertText: &slides.InsertTextRequest{
			ObjectId:       boxId,
			InsertionIndex: 0,
			Text:           "New Box Text Inserted",
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	return response.Replies[0].CreateShape.ObjectId
}

func createTestSheetsChart(presentationId string, pageId string, spreadsheetId string, sheetChartId int64) string {
	slidesService := getServices().Slides
	chartId := "MyChart_01"
	emu4M := slides.Dimension{Magnitude: 4000000, Unit: "EMU"}
	requests := []*slides.Request{{
		CreateSheetsChart: &slides.CreateSheetsChartRequest{
			ObjectId:      chartId,
			SpreadsheetId: spreadsheetId,
			ChartId:       sheetChartId,
			LinkingMode:   "LINKED",
			ElementProperties: &slides.PageElementProperties{
				PageObjectId: pageId,
				Size: &slides.Size{
					Height: &emu4M,
					Width:  &emu4M,
				},
				Transform: &slides.AffineTransform{
					ScaleX:     1,
					ScaleY:     1,
					TranslateX: 100000,
					TranslateY: 100000,
					Unit:       "EMU",
				},
			},
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	return response.Replies[0].CreateSheetsChart.ObjectId
}
