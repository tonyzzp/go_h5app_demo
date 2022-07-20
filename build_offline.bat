@echo off
echo compile offline
set GOOS=js
set GOARCH=wasm
go build -o docs/web/app.wasm
set GOOS=
set GOARCH=
go build 
go_h5app_demo.exe -type offline -website github
