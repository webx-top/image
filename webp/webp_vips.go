//go:build vips
// +build vips

package webp

import (
	"bytes"
	"io"

	"github.com/davidbyttow/govips/v2/vips"
)

func Encode(r io.Reader, quality float32) (buf *bytes.Buffer, err error) {
	var img *vips.ImageRef
	img, err = vips.NewImageFromReader(r)
	if err != nil {
		return
	}
	defer img.Close()
	var b []byte
	b, _, err = img.ExportWebp(&vips.WebpExportParams{
		Lossless: false,
		Quality:  int(quality),
	})
	if err != nil {
		return
	}
	buf = bytes.NewBuffer(b)
	return
}
