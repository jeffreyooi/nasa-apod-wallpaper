package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jeffreyooi/nasa-apod-wallpaper/apod"
)

const (
	// ConfigPath ...
	ConfigPath = "./configs/config.json"
)

func main() {
	bytes, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
	}
	var config apod.Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		panic(err)
	}

	if outputPath, err := apod.DownloadAPOD(config.APIKey, "", config.RequestHD); err != nil {
		panic(err)
	} else if err := apod.SetWallpaper(outputPath); err != nil {
		panic(err)
	} else if err := os.Remove(outputPath); err != nil {
		log.Fatalf("Failed to cleanup image at path %s after setting wallpaper", outputPath)
	}

	log.Println("Wallpaper set successfully")
}
