#! /bin/bash
go list ./... | grep -v vendor | xargs protobuild