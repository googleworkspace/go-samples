#!/bin/bash
set -ex
export GOOGLE_APPLICATION_CREDENTIALS="$PWD/../application_credentials.json";
export GOPATH="$PWD"
go test -v ./...
