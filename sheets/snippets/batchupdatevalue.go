package main

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func BatchUpdateValue(svc *sheets.Service, sheetId string) *sheets.BatchUpdateValuesResponse {
	data := [][]string{{"A1", "A2"}, {"B1", "B2"}}
	values := make([][]interface{}, len(data))
	for i, v := range data {
		values[i] = make([]interface{}, len(v))
		for j, vv := range v {
			values[i][j] = vv
		}
	}

	req := &sheets.BatchUpdateValuesRequest{
		Data: []*sheets.ValueRange{{
			MajorDimension: "ROWS",
			Range:          "A1:B2",
			Values:         values,
		}},
		ValueInputOption: "RAW",
	}

	resp, err := svc.Spreadsheets.Values.BatchUpdate(sheetId, req).Do()
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
