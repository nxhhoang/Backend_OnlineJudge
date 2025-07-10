package utils_test

import (
	"judge/utils"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestSuccessfulDownload(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	result := utils.DownloadPackage(332909, 1154548)
	var expected error = nil

	if result != expected {
		t.Errorf("TestSuccessfulDownload expected %s; got %s", expected, result)
	}
}
