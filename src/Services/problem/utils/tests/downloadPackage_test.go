package utils_test

import (
	"log"
	"problem/utils/polygon"
	"testing"

	"github.com/joho/godotenv"
)

func TestSuccessfulDownload(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	result := polygon.DownloadPackage(332909, 1154548)
	var expected error = nil

	if result != expected {
		t.Errorf("TestSuccessfulDownload expected %s; got %s", expected, result)
	}
}
