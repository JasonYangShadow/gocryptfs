package main

import (
	"io"
)

type ArgContainer struct {
	Offset      int64
	LimitedSize int64
}

func convertToInternlArgContainer(in *ArgContainer) *argContainer {
	arg := parseCliOpts([]string{"gocryptfs"})
	if in != nil {
		arg.offset = in.Offset
		arg.limitedSize = in.LimitedSize
	}
	arg.apptainer = true
	return &arg
}

func GocryptfsEncrypt(dir io.Reader) (io.ReadCloser, []byte, func(), error) {
	arg := convertToInternlArgContainer(nil)
	return mountEncrypt(dir, arg)
}

func GocryptfsDecrypt(sif io.Reader, offset, limitedSize int64) (io.ReadCloser, []byte, func(), error) {
	arg := &ArgContainer{
		Offset:      offset,
		LimitedSize: limitedSize,
	}

	internalArg := convertToInternlArgContainer(arg)
	return mountDecrypt(sif, internalArg)
}
