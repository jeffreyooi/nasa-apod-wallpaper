![Go Report Card](https://goreportcard.com/badge/github.com/jeffreyooi/nasa-apod-wallpaper)

# NASA APOD Wallpaper Setter

This is a NASA Astronomy Picture of the Day (APOD) wallpaper setter written in go.

The motivation to create this project is:
1. To learn Golang
2. To automatically change my desktop wallpaper daily to beautiful NASA APOD daily

## Before you run
You will need to generate an API key from https://api.nasa.gov/ and paste it on `apiKey` in `configs/config.json`. 

## Run the application

There are 2 ways to run the application, either run directly it with `go run`:
```sh
go run cmd/NasaWallpaper.go
```
or build it into a binary and run it:
```sh
go build cmd/NasaWallpaper.go

NasaWallpaper.exe
```

## Configurations
The configuration file is in `configs/config.json`. Currently there's only one option, that is to specify whether to download HD version of the file with `requestHDImage`.

## Supported platform
- Windows

## TODO
- [ ] Set as desktop wallpaper after download
- [ ] Make it run in background to automatically download and set wallpaper everyday
- [ ] Cross platform support (Linux, macOS)
