#!/bin/bash
set -ex
export GOPATH="$PWD"
go get -v ...
