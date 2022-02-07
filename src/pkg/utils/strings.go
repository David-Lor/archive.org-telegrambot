package utils

import (
	"bytes"
	"io"
	"io/ioutil"
)

func BytesToIOReadCloser(data []byte) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(data))
}
