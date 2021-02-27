package snippets

import (
	"log"

	slides "google.golang.org/api/slides/v1"
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
