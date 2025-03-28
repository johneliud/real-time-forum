package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/johneliud/real-time-forum/backend/logger"
)

// CreatePost handles the creation of a new post.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Error("Invalid request method", "method", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		logger.Error("Failed to create uploads directory", "err", err)
		http.Error(w, "Failed to create uploads directory", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		logger.Error("Failed to parse multipart form", "err", err)
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		if err.Error() == "http: no such file" {
			logger.Error("No file uploaded", "err", err)
			http.Error(w, "No file uploaded", http.StatusBadRequest)
			return
		} else {
			logger.Error("Failed to get file", "err", err)
			http.Error(w, "Failed to get file", http.StatusInternalServerError)
			return
		}
	}

	if file != nil {
		defer file.Close()

		if header.Size > 10<<20 {
			logger.Error("File is too large", "size", header.Size)
			http.Error(w, "File is too large", http.StatusBadRequest)
			return
		}

		_, err := validateMimeType(file)
		if err != nil {
			logger.Error("Invalid file type", "err", err)
			http.Error(w, "Invalid file type", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			logger.Error("Failed to seek file", "err", err)
			http.Error(w, "Failed to seek file", http.StatusInternalServerError)
			return
		}

		temporaryFile, err := os.Create("uploads/" + header.Filename)
		if err != nil {
			logger.Error("Failed to create temporary file", "err", err)
			http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
			return
		}
		defer temporaryFile.Close()

		_, err = temporaryFile.WriteTo(temporaryFile)
		if err != nil {
			logger.Error("Failed to write file to temporary file", "err", err)
			http.Error(w, "Failed to write file to temporary file", http.StatusInternalServerError)
			return
		}

		temporaryFilePath := temporaryFile.Name()
		url := fmt.Sprintf("%v", temporaryFilePath)
		logger.Info("File uploaded successfully", "url", url)
	}
}

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
