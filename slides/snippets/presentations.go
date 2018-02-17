package snippets

import (
	"fmt"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/slides/v1"
	"log"
)

func createPresentation() string {
	slidesService := getServices().Slides
	// [START createPresentation]
	// Create a presentation and request a PresentationId.
	p := &slides.Presentation{
		Title: "Title",
	}
	presentation, err := slidesService.Presentations.Create(p).Fields(
		"presentationId",
	).Do()
	if err != nil {
		log.Fatalf("Unable to create presentation. %v", err)
	}
	fmt.Printf("Created presentation with ID: %s", presentation.PresentationId)
	// [END createPresentation]
	return presentation.PresentationId
}

func copyPresentation(id string, title string) string {
	driveService := getServices().Drive
	// [START copyPresentation]
	// Copy a presentation.
	file := drive.File{
		Title: title,
	}
	presentationCopyFile, err := driveService.Files.Copy(id, &file).Do()
	if err != nil {
		log.Fatalf("Unable to copy presentation. %v", err)
	}
	presentationCopyId := presentationCopyFile.Id
	// [END copyPresentation]
	return presentationCopyId
}

func createSlide(presentationId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START createSlide]
	// Add a slide at index 1 using the predefined "TITLE_AND_TWO_COLUMNS" layout
	// and the ID "MyNewSlide_001".
	slideId := "MyNewSlide_001"
	requests := []*slides.Request{{
		CreateSlide: &slides.CreateSlideRequest{
			ObjectId:       slideId,
			InsertionIndex: 1,
			SlideLayoutReference: &slides.LayoutReference{
				PredefinedLayout: "TITLE_AND_TWO_COLUMNS",
			},
		},
	}}

	// If you wish to populate the slide with elements, add create requests here,
	// using the slide ID specified above.

	// Execute the request.
	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	response, err := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	if err != nil {
		log.Fatalf("Unable to create slide. %v", err)
	}
	fmt.Printf("Created slide with ID: %s", response.Replies[0].CreateSlide.ObjectId)
	// [END createSlide]
	return *response
}

func createTextBoxWithText(presentationId string, slideId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START createTextBoxWithText]
	// Create a new square text box, using a supplied object ID.
	textBoxId := "MyTextBox_01"
	pt350 := slides.Dimension{
		Magnitude: 350,
		Unit:      "PT",
	}
	requests := []*slides.Request{{
		// Create a new square text box, using a supplied object ID.
		CreateShape: &slides.CreateShapeRequest{
			ObjectId:  textBoxId,
			ShapeType: "TEXT_BOX",
			ElementProperties: &slides.PageElementProperties{
				PageObjectId: slideId,
				Size: &slides.Size{
					Height: &pt350,
					Width:  &pt350,
				},
				Transform: &slides.AffineTransform{
					ScaleX:     1.0,
					ScaleY:     1.0,
					TranslateX: 350.0,
					TranslateY: 100.0,
					Unit:       "PT",
				},
			},
		},
	}, {
		// Insert text into the box, using the object ID given to it.
		InsertText: &slides.InsertTextRequest{
			ObjectId:       textBoxId,
			InsertionIndex: 0,
			Text:           "New Box Text Inserted",
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	response, err := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	if err != nil {
		log.Errorf("Unable to create text box. %v", err)
	}
	fmt.Printf("Created text box with ID: %s", response.Replies[0].CreateShape.ObjectId)
	// [END createTextBoxWithText]
	return *response
}

func createImage(presentationId string, slideId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	imageURL := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
	// [START createImage]
	// Temporarily upload a local image file to Drive, in order to obtain a URL
	// for the image. Alternatively, you can provide the Slides service a URL of
	// an already hosted image.
	//
	// We will use an existing image under the variable: imageURL.
	//
	// Create a new image, using the supplied object ID, with content downloaded from imageURL.
	imageId := "MyImageId_01"
	emu4M := slides.Dimension{Magnitude: 4000000, Unit: "EMU"}
	requests := []*slides.Request{{
		CreateImage: &slides.CreateImageRequest{
			ObjectId: imageId,
			Url:      imageURL,
			ElementProperties: &slides.PageElementProperties{
				PageObjectId: slideId,
				Size: &slides.Size{
					Height: &emu4M,
					Width:  &emu4M,
				},
				Transform: &slides.AffineTransform{
					ScaleX:     1.0,
					ScaleY:     1.0,
					TranslateX: 100000.0,
					TranslateY: 100000.0,
					Unit:       "EMU",
				},
			},
		},
	}}

	// Execute the request.
	body := &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	}
	response, err := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	if err != nil {
		log.Fatalf("Unable to create image. %v", err)
	} else {
		fmt.Printf("Created image with ID: %s", response.Replies[0].CreateImage.ObjectId)
	}
	// [END createImage]
	return *response
}

func textMerging(templatePresentationId string, dataSpreadsheetId string) []slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	driveService := getServices().Drive
	sheetsService := getServices().Sheets
	responses := make([]slides.BatchUpdatePresentationResponse, 0)

	// [START textMerging]
	// Use the Sheets API to load data, one record per row.
	dataRangeNotation := "Customers!A2:M6"
	sheetsResponse, _ := sheetsService.Spreadsheets.Values.Get(dataSpreadsheetId, dataRangeNotation).Do()
	values := sheetsResponse.Values

	// For each record, create a new merged presentation.
	for _, row := range values {
		customerName := row[2].(string)
		caseDescription := row[5].(string)
		totalPortfolio := row[11].(string)

		// Duplicate the template presentation using the Drive API.
		copyTitle := customerName + " presentation"
		file := drive.File{
			Title: copyTitle,
		}
		presentationFile, _ := driveService.Files.Copy(templatePresentationId, &file).Do()
		presentationId := presentationFile.Id

		// Create the text merge (replaceAllText) requests for this presentation.
		requests := []*slides.Request{{
			ReplaceAllText: &slides.ReplaceAllTextRequest{
				ContainsText: &slides.SubstringMatchCriteria{
					Text:      "{{customer-name}}",
					MatchCase: true,
				},
				ReplaceText: customerName,
			},
		}, {
			ReplaceAllText: &slides.ReplaceAllTextRequest{
				ContainsText: &slides.SubstringMatchCriteria{
					Text:      "{{case-description}}",
					MatchCase: true,
				},
				ReplaceText: caseDescription,
			},
		}, {
			ReplaceAllText: &slides.ReplaceAllTextRequest{
				ContainsText: &slides.SubstringMatchCriteria{
					Text:      "{{total-portfolio}}",
					MatchCase: true,
				},
				ReplaceText: totalPortfolio,
			},
		}}

		// Execute the requests for this presentation.
		body := &slides.BatchUpdatePresentationRequest{
			Requests: requests,
		}
		response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
		// [START_EXCLUDE silent]
		responses = append(responses, *response)
		// [END_EXCLUDE silent]

		// Count total number of replacements made.
		var numReplacements int64 = 0
		for _, resp := range response.Replies {
			numReplacements += resp.ReplaceAllText.OccurrencesChanged
		}

		fmt.Printf("Created merged presentation for %s with ID %s\n", customerName, presentationId)
		fmt.Printf("Replaced %d text instances.\n", numReplacements)
	}
	// [END textMerging]
	return responses
}

func imageMerging(templatePresentationId string, imageURL string, customerName string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	driveService := getServices().Drive
	logoURL := imageURL
	customerGraphicURL := imageURL

	// [START imageMerging]
	// Duplicate the template presentation using the Drive API.
	copyTitle := customerName + " presentation"
	file := drive.File{
		Title: copyTitle,
	}
	presentationFile, _ := driveService.Files.Copy(templatePresentationId, &file).Do()
	presentationId := presentationFile.Id

	// Create the image merge (replaceAllShapesWithImage) requests.
	requests := []*slides.Request{{
		ReplaceAllShapesWithImage: &slides.ReplaceAllShapesWithImageRequest{
			ImageUrl:      logoURL,
			ReplaceMethod: "CENTER_INSIDE",
			ContainsText: &slides.SubstringMatchCriteria{
				Text:      "{{company-logo}}",
				MatchCase: true,
			},
		},
	}, {
		ReplaceAllShapesWithImage: &slides.ReplaceAllShapesWithImageRequest{
			ImageUrl:      customerGraphicURL,
			ReplaceMethod: "CENTER_INSIDE",
			ContainsText: &slides.SubstringMatchCriteria{
				Text:      "{{customer-graphic}}",
				MatchCase: true,
			},
		},
	}}

	// Execute the requests for this presentation.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()

	// Count total number of replacements made.
	numReplacements := 0
	for _, resp := range response.Replies {
		numReplacements += resp.ReplaceAllShapesWithImage.OccurrencesChanged
	}
	fmt.Printf("Created merged presentation with ID %s\n", presentationId)
	fmt.Printf("Replaced %d shapes instances with images.\n", numReplacements)
	// [END imageMerging]
	return *response
}

func simpleTextReplace(presentationId string, shapeId string, replacementText string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START simpleTextReplace]
	// Remove existing text in the shape, then insert the new text.
	requests := []*slides.Request{{
		DeleteText: &slides.DeleteTextRequest{
			ObjectId: shapeId,
			TextRange: &slides.Range{
				Type: "All",
			},
		},
	}, {
		InsertText: &slides.InsertTextRequest{
			ObjectId:       shapeId,
			InsertionIndex: 0,
			Text:           replacementText,
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	fmt.Printf("Replaced text in shape with ID: %s", shapeId)
	// [END simpleTextReplace]
	return *response
}

func textStyleUpdate(presentationId string, shapeId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START textStyleUpdate]
	// Update the text style so that the first 5 characters are bolded
	// and italicized, and the next 5 are displayed in blue 14 pt Times
	// New Roman font, and the next five are hyperlinked.
	requests := []*slides.Request{{
		UpdateTextStyle: &slides.UpdateTextStyleRequest{
			ObjectId: shapeId,
			TextRange: &slides.Range{
				Type:            "FIXED_RANGE",
				StartIndex:      0,
				EndIndex:        5,
				ForceSendFields: []string{"StartIndex"},
			},
			Style: &slides.TextStyle{
				Bold:   true,
				Italic: true,
			},
			Fields: "bold,italic",
		},
	}, {
		UpdateTextStyle: &slides.UpdateTextStyleRequest{
			ObjectId: shapeId,
			TextRange: &slides.Range{
				Type:       "FIXED_RANGE",
				StartIndex: 5,
				EndIndex:   10,
			},
			Style: &slides.TextStyle{
				FontFamily: "Times New Roman",
				FontSize: &slides.Dimension{
					Magnitude: 14.0,
					Unit:      "PT",
				},
				ForegroundColor: &slides.OptionalColor{
					OpaqueColor: &slides.OpaqueColor{
						RgbColor: &slides.RgbColor{
							Blue:  1.0,
							Green: 0.0,
							Red:   0.0,
						},
					},
				},
			},
			Fields: "foregroundColor,fontFamily,fontSize",
		},
	}, {
		UpdateTextStyle: &slides.UpdateTextStyleRequest{
			ObjectId: shapeId,
			TextRange: &slides.Range{
				Type:       "FIXED_RANGE",
				StartIndex: 10,
				EndIndex:   15,
			},
			Style: &slides.TextStyle{
				Link: &slides.Link{
					Url: "www.example.com",
				},
			},
			Fields: "link",
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	fmt.Printf("Updated text style for shape with ID: %s", shapeId)
	// [END textStyleUpdate]
	return *response
}

func createBulletedText(presentationId string, shapeId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START createBulletedText]
	// Add arrow-diamond-disc bullets to all text in the shape.
	requests := []*slides.Request{{
		CreateParagraphBullets: &slides.CreateParagraphBulletsRequest{
			ObjectId: shapeId,
			TextRange: &slides.Range{
				Type: "ALL",
			},
			BulletPreset: "BULLET_ARROW_DIAMOND_DISC",
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	fmt.Printf("Added a linked Sheets chart with ID %s", shapeId)
	// [END textStyleUpdate]
	return *response
}

func createSheetsChart(presentationId string, pageId string, spreadsheetId string, sheetChartId int64) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START createSheetsChart]
	// Embed a Sheets chart (indicated by the spreadsheetId and sheetChartId) onto
	// a page in the presentation. Setting the linking mode as "LINKED" allows the
	// chart to be refreshed if the Sheets version is updated.
	emu4M := slides.Dimension{Magnitude: 4000000, Unit: "EMU"}
	presentationChartId := "MyEmbeddedChart"
	requests := []*slides.Request{{
		CreateSheetsChart: &slides.CreateSheetsChartRequest{
			ObjectId:      presentationChartId,
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
					ScaleX:     1.0,
					ScaleY:     1.0,
					TranslateX: 100000.0,
					TranslateY: 100000.0,
					Unit:       "EMU",
				},
			},
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	fmt.Printf("Added a linked Sheets chart with ID %s", presentationChartId)
	// [END createSheetsChart]
	return *response
}

func createRefreshSheetsChart(presentationId string, presentationChartId string) slides.BatchUpdatePresentationResponse {
	slidesService := getServices().Slides
	// [START refreshSheetsChart]
	requests := []*slides.Request{{
		RefreshSheetsChart: &slides.RefreshSheetsChartRequest{
			ObjectId: presentationChartId,
		},
	}}

	// Execute the requests.
	body := &slides.BatchUpdatePresentationRequest{Requests: requests}
	response, _ := slidesService.Presentations.BatchUpdate(presentationId, body).Do()
	fmt.Printf("Refreshed a linked Sheets chart with ID %s", presentationChartId)
	// [END createSheetsChart]
	return *response
}
