package main

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func AppnedValue(svc *sheets.Service, sheetId string) *sheets.AppendValuesResponse {
	rng := "A1:B2"
	data := [][]string{{"a1", "a2"}, {"b1", "b2"}}
	values := make([][]interface{}, len(data))
	for i, v := range data {
		values[i] = make([]interface{}, len(v))
		for j, vv := range v {
			values[i][j] = vv
		}
	}

	vr := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Range:          "A1:B2",
		Values:         values,
	}

	resp, err := svc.Spreadsheets.Values.Append(sheetId, rng, vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
