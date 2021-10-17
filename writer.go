package d2editor

import (
	"github.com/vitalick/go-d2editor/consts"
	"io"
)

type writerWrapper struct {
	w        io.Writer
	Filesize int
	b        []byte
}

func (w *writerWrapper) Write(p []byte) (int, error) {
	w.b = append(w.b, p...)
	n := len(p)
	w.Filesize += n
	return n, nil
}

func (w *writerWrapper) EndWrite() (int, error) {
	bs := make([]byte, 4)
	consts.BinaryEndian.PutUint32(bs, uint32(w.Filesize))
	for i, b := range bs {
		w.b[i+8] = b
	}
	checksum := 0
	for _, b := range w.b {
		checksum = ChecksumAppend(b, checksum)
	}
	bs = make([]byte, 4)
	consts.BinaryEndian.PutUint32(bs, uint32(checksum))
	for i, b := range bs {
		w.b[i+12] = b
	}
	return w.w.Write(w.b)
}
