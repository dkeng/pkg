package gzip

import (
	"bytes"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
)

const (
	inputFileName = ""
	outFileName   = ""
)

func TestExecute(t *testing.T) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()
	var gzipBuffer = new(bytes.Buffer)
	Execute(inputFile, gzipBuffer)
	err = ioutil.WriteFile(outFileName, gzipBuffer.Bytes(), syscall.O_CREAT)
	if err != nil {
		t.Fatal(err)
	}
}
