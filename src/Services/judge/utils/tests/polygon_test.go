package utils_test

import (
	"fmt"
	"judge/utils"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestSuccessfulDownload(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found or failed to load")
	}

	result := utils.DownloadPackage(map[string]string{
		"problemId": "442306",
		"packageId": "1145478",
		"type":      "standard",
		"apiKey":    os.Getenv("POLYGON_API_KEY"),
		"time":      fmt.Sprintf("%d", time.Now().Unix()),
	})
	var expected error = nil

	if result != expected {
		t.Errorf("TestSuccessfulDownload expected %s; got %s", expected, result)
	}
}
