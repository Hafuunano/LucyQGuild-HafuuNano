// Package utils Transform => Make Image To Local Proxy Server, send Link.
package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

// ** Be aware this should use env. **

var localPath = os.Getenv("localpath")
var remotePath = os.Getenv("remotepath")

func ImageToTransform(contentName string, images image.Image) string {
	var buffer bytes.Buffer
	png.Encode(&buffer, images)
	getDest, err := os.Create(localPath + contentName)
	_, err = io.Copy(getDest, bytes.NewReader(buffer.Bytes()))
	if err != nil {
		panic(err)
	}
	fmt.Print(remotePath + contentName)
	return remotePath + contentName
}

func TransformLink(contentPlace string) string {
	getData, err := os.ReadFile(contentPlace)
	if err != nil {
		panic(err)
	}
	getSourceName := filepath.Base(contentPlace)
	getDest, err := os.Create(localPath + getSourceName)
	_, err = io.Copy(getDest, bytes.NewReader(getData))
	if err != nil {
		panic(err)
	}
	return remotePath + getSourceName
}

func UploadedCleaner() {
	os.RemoveAll(localPath)
}
