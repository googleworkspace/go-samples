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
	"github.com/stretchr/testify/assert"
	"testing"
)

const IMAGE_URL = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
const TEMPLATE_PRESENTATION_ID = "1wJUN1B5CQ2wQOBzmz2apky48QNK1OsE2oNKHPMLpKDc"
const DATA_SPREADSHEET_ID = "14KaZMq2aCAGt5acV77zaA_Ps8aDt04G7T0ei4KiXLX8"
const CHART_ID = 1107320627
const CUSTOMER_NAME = "Fake Customer"

func TestCreatePresentation(t *testing.T) {
	presentationId := createPresentation()
	assert.NotNil(t, presentationId)
	deleteFileOnCleanup(presentationId)
}

func TestCopyPresentation(t *testing.T) {
	presentationId := createPresentation()
	copyId := copyPresentation(presentationId, "My Duplicate Presentation")
	assert.NotNil(t, copyId)
	deleteFileOnCleanup(presentationId)
	deleteFileOnCleanup(copyId)
}

func TestCreateSlide(t *testing.T) {
	presentationId := createPresentation()
	response := createSlide(presentationId)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Replies))
	assert.NotNil(t, response.Replies[0].CreateSlide.ObjectId)
}

func TestCreateTextBox(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	assert.NotNil(t, pageId)
	response := createTextBoxWithText(presentationId, pageId)
	assert.Equal(t, 2, len(response.Replies))
	boxId := response.Replies[0].CreateShape.ObjectId
	assert.NotNil(t, boxId)
}

func TestCreateImage(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	response := createImage(presentationId, pageId)
	assert.Equal(t, 1, len(response.Replies))
	imageId := response.Replies[0].CreateImage.ObjectId
	assert.NotNil(t, imageId)
}

func TestTextMerge(t *testing.T) {
	responses := textMerging(TEMPLATE_PRESENTATION_ID, DATA_SPREADSHEET_ID)
	for _, response := range responses {
		assert.NotNil(t, response.PresentationId)
		assert.Equal(t, int64(3), int64(len(response.Replies)))
		var numReplacements int64 = 0
		for _, reply := range response.Replies {
			numReplacements += reply.ReplaceAllText.OccurrencesChanged
		}
		assert.Equal(t, int64(4), int64(numReplacements))
		deleteFileOnCleanup(response.PresentationId)
	}
}

func TestImageMerge(t *testing.T) {
	response := imageMerging(TEMPLATE_PRESENTATION_ID, IMAGE_URL, CUSTOMER_NAME)
	presentationId := response.PresentationId
	assert.NotNil(t, presentationId)
	assert.Equal(t, 2, len(response.Replies))
	var numReplacements int64 = 0
	for _, reply := range response.Replies {
		numReplacements += reply.ReplaceAllShapesWithImage.OccurrencesChanged
	}
	assert.Equal(t, int64(2), numReplacements)
	deleteFileOnCleanup(response.PresentationId)
}

func TestSimpleTextReplace(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	boxId := createTestTextbox(presentationId, pageId)
	response := simpleTextReplace(presentationId, boxId, "MY NEW TEXT")
	assert.Equal(t, 2, len(response.Replies))
}

func TestTextStyleUpdate(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	boxId := createTestTextbox(presentationId, pageId)
	response := textStyleUpdate(presentationId, boxId)
	assert.Equal(t, 3, len(response.Replies))
}

func TestCreateBulletText(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	boxId := createTestTextbox(presentationId, pageId)
	response := createBulletedText(presentationId, boxId)
	assert.Equal(t, 1, len(response.Replies))
}

func TestCreateSheetsChart(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	response := createSheetsChart(presentationId, pageId, DATA_SPREADSHEET_ID, CHART_ID)
	assert.Equal(t, 1, len(response.Replies))
	assert.NotNil(t, response.Replies[0].CreateSheetsChart.ObjectId)
}

func TestRefreshSheetsChart(t *testing.T) {
	presentationId := createTestPresentation()
	pageId := createTestSlide(presentationId)
	chartId := createTestSheetsChart(presentationId, pageId, DATA_SPREADSHEET_ID, CHART_ID)
	response := createRefreshSheetsChart(presentationId, chartId)
	assert.Equal(t, 1, len(response.Replies))
}
