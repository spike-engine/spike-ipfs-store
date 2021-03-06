package pinata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type PinResponse struct {
	IpfsHash    string `json:"IpfsHash,omitempty"`
	PinSize     int64  `json:"PinSize,omitempty"`
	Timestamp   string `json:"Timestamp,omitempty"`
	Error       string `json:"error,omitempty"`
	IsDuplicate bool   `json:"isDuplicate,omitempty"`
}

func (c *Client) PinFile(filepath string) (PinResponse, error) {
	b, w, err := createMultipartFormData(filepath)
	req, _ := http.NewRequest("POST", c.Node+ApiPinFile, &b)
	req.Header.Set("Authorization", "Bearer "+c.JWT)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return PinResponse{}, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PinResponse{}, err
	}
	fmt.Println("debug joy", string(content))
	var out PinResponse
	if err = json.Unmarshal(content, &out); err != nil {
		return PinResponse{}, err
	}
	return out, nil
}

func createMultipartFormData(filePath string) (bytes.Buffer, *multipart.Writer, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, err := os.Open(filePath)
	if err != nil {
		return b, w, err
	}
	if fw, err = w.CreateFormFile("file", formatFilename(filePath)); err != nil {
		return b, w, err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return b, w, err
	}
	w.Close()
	return b, w, nil
}

func formatFilename(path string) string {
	items := strings.Split(path, "/")
	if len(items) > 0 {
		return items[len(items)-1]
	}
	return ""
}
