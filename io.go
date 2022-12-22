package wen

import (
	"errors"
	"io"
	"time"
)

// 同 io.ReadAtLeast，加了超时
func ReadAtLeast(r io.Reader, buf []byte, min int, timeout time.Duration) (n int, err error) {
	if len(buf) < min {
		return 0, io.ErrShortBuffer
	}
	defer func() {
		if n >= min {
			err = nil
		} else if n > 0 && err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()
	tout := time.After(timeout)
	for {
		select {
		case <-tout:
			err = errors.New("read timeout")
			return
		default:
			if n < min && err == nil {
				var nn int
				nn, err = r.Read(buf[n:])
				n += nn
			} else {
				return
			}
		}
	}
}

// 同 io.ReadFull，加了超时
func ReadFull(r io.Reader, buf []byte, timeout time.Duration) (n int, err error) {
	return ReadAtLeast(r, buf, len(buf), timeout)
}
