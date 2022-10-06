package main

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func BatchUpdate(svc *sheets.Service, sheetId string) *sheets.BatchUpdateSpreadsheetResponse {
	// batch update request for add sheet
	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{{
			AddSheet: &sheets.AddSheetRequest{
				Properties: &sheets.SheetProperties{
					Title: "AddSheet",
				},
			},
		}},
	}

	resp, err := svc.Spreadsheets.BatchUpdate(sheetId, req).Do()
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
