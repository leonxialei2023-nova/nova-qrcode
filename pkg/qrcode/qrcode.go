package qrcode

import (
	"errors"

	lib "github.com/skip2/go-qrcode"
)

func encodeLevel(l Level) lib.RecoveryLevel {
	switch l {
	case Low:
		return lib.Low
	case Medium:
		return lib.Medium
	case High:
		return lib.High
	case Highest:
		return lib.Highest
	default:
		return lib.Medium
	}
}

func EncodePNG(text string, opt Options) ([]byte, error) {
	if text == "" {
		return nil, errors.New("empty qrcode text")
	}

	size := opt.Size
	if size <= 0 {
		size = 256
	}

	level := encodeLevel(opt.Level)

	return lib.Encode(text, level, size)
}
