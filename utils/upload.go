package utils

import (
	"Restful_Go/pkg/cloudinary"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
