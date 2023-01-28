@ECHO off

ECHO Building Windows binary...
go env -w GO111MODULE=on
go build -o bin\athp7.exe src\main.go