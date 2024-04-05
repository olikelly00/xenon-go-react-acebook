package controllers

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/go-resty/resty/v2"
)

func UploadFileToHostingService(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	client := resty.New()
	api_key := os.Getenv("IMGBB_API_KEY")
	client.SetFormData(map[string]string{
		"key": api_key,
	})

	resp, err := client.R().
		SetFileReader("image", fileHeader.Filename, file).
		Post("https://api.imgbb.com/1/upload")
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("failed to upload image: %s", resp.String())
	}

	var imgResponse struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	err = json.Unmarshal(resp.Body(), &imgResponse)
	if err != nil {
		return "", err
	}

	return imgResponse.Data.URL, nil
}
