package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/logger"
)

// validateMimeType validates the mime type of the file.
func validateMimeType(file multipart.File) (string, error) {
	allowedMimeTypes := map[string]string{
		"image/jpeg": ".jpeg",
		"image/jpg":  ".jpg",
		"image/png":  ".png",
		"image/gif":  ".gif",
	}

	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		logger.Error("Failed to read file", "err", err)
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	extension, ok := allowedMimeTypes[mimeType]
	if !ok {
		return "", fmt.Errorf("invalid mime type: %v", mimeType)
	}
	return extension, nil
}
