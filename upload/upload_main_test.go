package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "resources/0c77348b078036a0.gif"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))

	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	upload, _ := os.Open(uploadFilePath)
	origin, _ := os.Open(path)
	defer upload.Close()
	defer origin.Close()

	uploadData := []byte{}
	originData := []byte{}

	upload.Read(uploadData)
	origin.Read(originData)

	assert.Equal(originData, uploadData)
}
