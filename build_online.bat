@echo off
echo compile online
set GOOS=js
set GOARCH=wasm
go build -o web/app.wasm
set GOOS=
set GOARCH=
go build 
go_h5app_demo.exe
