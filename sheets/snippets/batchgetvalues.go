package main

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func BatchGetValues(svc *sheets.Service, sheetId string) *sheets.BatchGetValuesResponse {
	c := svc.Spreadsheets.Values.BatchGet(sheetId)
	c.Ranges("A1:B2")
	resp, err := c.Do()
	if err != nil {
		log.Fatalln(err)
	}

	// print range and values
	// for _, v := range res.ValueRanges {
	// 	fmt.Println(v.Range)
	// 	fmt.Println(v.Values)
	// }

	return resp
}
