@echo off
go tool templ generate
go build -o ./tmp/main.exe ./cmd/main.go


