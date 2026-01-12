# Google Workspace Go Samples

[![Build Status](https://github.com/googleworkspace/go-samples/actions/workflows/test.yml/badge.svg)](https://github.com/googleworkspace/go-samples/actions/workflows/test.yml)
[![Lint Status](https://github.com/googleworkspace/go-samples/actions/workflows/lint.yml/badge.svg)](https://github.com/googleworkspace/go-samples/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/googleworkspace/go-samples)](https://goreportcard.com/report/github.com/googleworkspace/go-samples)

Go samples for [Google Workspace APIs](https://developers.google.com/gsuite/) docs.

To run the quickstarts, download a `credentials.json` file in the `quickstart` folder by following the instructions in `quickstart/README.md`.

## APIs

| API                     | GoDoc                                                                                                                                         | Quickstart                                                                 | Snippets                                                      |
| ----------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------- |
| **Admin SDK Directory** | [![GoDoc](https://godoc.org/google.golang.org/api/admin/directory/v1?status.svg)](https://godoc.org/google.golang.org/api/admin/directory/v1) | [Link](https://developers.google.com/admin-sdk/directory/v1/quickstart/go) | -                                                             |
| **Admin SDK Reports**   | [![GoDoc](https://godoc.org/google.golang.org/api/admin/reports/v1?status.svg)](https://godoc.org/google.golang.org/api/admin/reports/v1)     | [Link](https://developers.google.com/admin-sdk/reports/v1/quickstart/go)   | -                                                             |
| **Admin SDK Reseller**  | [![GoDoc](https://godoc.org/google.golang.org/api/reseller/v1?status.svg)](https://godoc.org/google.golang.org/api/reseller/v1)               | [Link](https://developers.google.com/admin-sdk/reseller/v1/quickstart/go)  | -                                                             |
| **Apps Script**         | [![GoDoc](https://godoc.org/google.golang.org/api/script/v1?status.svg)](https://godoc.org/google.golang.org/api/script/v1)                   | [Link](https://developers.google.com/apps-script/api/quickstart/go)        | -                                                             |
| **Calendar**            | [![GoDoc](https://godoc.org/google.golang.org/api/calendar/v1?status.svg)](https://godoc.org/google.golang.org/api/calendar/v1)               | [Link](https://developers.google.com/calendar/quickstart/go)               | -                                                             |
| **Classroom**           | [![GoDoc](https://godoc.org/google.golang.org/api/classroom/v1?status.svg)](https://godoc.org/google.golang.org/api/classroom/v1)             | [Link](https://developers.google.com/classroom/quickstart/go)              | -                                                             |
| **Docs**                | [![GoDoc](https://godoc.org/google.golang.org/api/docs/v1?status.svg)](https://godoc.org/google.golang.org/api/docs/v1)                       | [Link](https://developers.google.com/docs/api/quickstart/go)               | -                                                             |
| **Drive V3**            | [![GoDoc](https://godoc.org/google.golang.org/api/drive/v3?status.svg)](https://godoc.org/google.golang.org/api/drive/v3)                     | [Link](https://developers.google.com/drive/v3/web/quickstart/go)           | [Link](https://developers.google.com/drive/v3/web/about-sdk)  |
| **Gmail**               | [![GoDoc](https://godoc.org/google.golang.org/api/gmail/v1?status.svg)](https://godoc.org/google.golang.org/api/gmail/v1)                     | [Link](https://developers.google.com/gmail/api/quickstart/go)              | -                                                             |
| **People**              | [![GoDoc](https://godoc.org/google.golang.org/api/people/v1?status.svg)](https://godoc.org/google.golang.org/api/people/v1)                   | [Link](https://developers.google.com/people/quickstart/go)                 | -                                                             |
| **Sheets**              | [![GoDoc](https://godoc.org/google.golang.org/api/sheets/v4?status.svg)](https://godoc.org/google.golang.org/api/sheets/v4)                   | [Link](https://developers.google.com/sheets/api/quickstart/go)             | -                                                             |
| **Slides**              | [![GoDoc](https://godoc.org/google.golang.org/api/slides/v1?status.svg)](https://godoc.org/google.golang.org/api/slides/v1)                   | [Link](https://developers.google.com/slides/quickstart/go)                 | [Link](https://developers.google.com/slides/how-tos/overview) |
| **Tasks**               | [![GoDoc](https://godoc.org/google.golang.org/api/tasks/v1?status.svg)](https://godoc.org/google.golang.org/api/tasks/v1)                     | [Link](https://developers.google.com/google-apps/tasks/quickstart/go)      | -                                                             |

## Development

Use the following commands to maintain the repository:

### Format

```bash
go fmt ./...
```

### Build

```bash
go build -v ./...
```

### Vet

```bash
go vet ./...
```

### Tidy

```bash
go mod tidy
```
