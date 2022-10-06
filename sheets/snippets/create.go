package main

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func Create(svc *sheets.Service) string {
	newSheet := &sheets.Spreadsheet{
		Sheets: []*sheets.Sheet{{
			Properties: &sheets.SheetProperties{
				Title: "ExampleSheet1",
			},
		}, {Properties: &sheets.SheetProperties{
			Title: "ExampleSheet2",
		}}},
		Properties: &sheets.SpreadsheetProperties{
			Title: "ExampleSpreadsheet1",
		},
	}
	resp, err := svc.Spreadsheets.Create(newSheet).Do()
	if err != nil {
		log.Fatalln(err)
	}
	return resp.SpreadsheetId
}
