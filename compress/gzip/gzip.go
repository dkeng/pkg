package gzip

import (
	"compress/gzip"
	"io"
)

var (
	// BufferSize 缓冲区大小 32KB
	BufferSize = 32768
)

// Execute 执行压缩
func Execute(r io.Reader, w io.Writer) error {
	wGzip := gzip.NewWriter(w)
	defer wGzip.Close()
	fileBuffer := make([]byte, BufferSize)
	for {
		lenght, err := r.Read(fileBuffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		wGzip.Write(fileBuffer[0:lenght])
	}
	return nil
}
