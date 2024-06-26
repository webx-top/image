//go:build !vips
// +build !vips

package webp

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"strings"

	"github.com/admpub/errors"
	"github.com/admpub/log"
	"github.com/pcpl2/go-webp"
	"golang.org/x/image/bmp"
)

func Encode(r io.Reader, quality float32) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	var (
		img image.Image
		b   = make([]byte, 512)
	)
	_, err = r.Read(b)
	if err != nil {
		return
	}
	if rs, ok := r.(io.ReadSeeker); ok {
		rs.Seek(0, 0)
	}
	contentType := http.DetectContentType(b)
	switch {
	case strings.Contains(contentType, "jpeg"):
		img, err = jpeg.Decode(r)
	case strings.Contains(contentType, "png"):
		img, err = png.Decode(r)
	case strings.Contains(contentType, "bmp"):
		img, err = bmp.Decode(r)
	case strings.Contains(contentType, "gif"):
		log.Warn("Gif support is not perfect!")
		img, err = gif.Decode(r)
	default:
		return nil, errors.WithMessage(ErrUnsupportedImageType, contentType)
	}
	if err != nil {
		return
	}
	options := &webp.Options{
		Lossless: false,
		Quality:  quality,
	}
	err = webp.Encode(buf, img, options)
	return
}
