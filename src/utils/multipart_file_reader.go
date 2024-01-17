package utils

import (
	"bytes"
	"io"
	"mime/multipart"
)

func ReadMultipartFile(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}

	defer f.Close()

	buf := bytes.NewBuffer(nil)

	if _, err := io.Copy(buf, f); err != nil {
		return "", err
	}

	return string(buf.String()), nil
}
