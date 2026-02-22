@echo off
echo Building for Linux...

REM Устанавливаем переменные окружения
set GOOS=linux
set GOARCH=amd64

REM Собираем приложение
go build -o vendor-server cmd/main.go

REM Очищаем переменные
set GOOS=
set GOARCH=

echo Build complete: ./vendor-server