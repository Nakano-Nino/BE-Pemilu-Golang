package cloudinary

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/joho/godotenv"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cldSecret := os.Getenv("API_SECRET")
	cldName := os.Getenv("CLOUD_NAME")
	cldKey := os.Getenv("API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
