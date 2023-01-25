package main

import "io"

type ArgContainer struct {
	Offset      uint64
	LimitedSize uint64
}

func GocryptfsEncrypt(dir io.Reader) (io.ReadCloser, error) {
	return nil, nil
}

func GocryptfsDecrypt(sif io.Reader, offset, limitedSize uint64) (io.ReadCloser, error) {
	return nil, nil
}
