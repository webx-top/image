package image

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWatermark(t *testing.T) {
	watermarkFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/admpub/nging/public/assets/backend/images/nging-gear.png`)
	wm, err := NewWatermark(watermarkFile,0,TopLeft)
	if err != nil {
		t.Fatal(err)
	}
	destinationFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/webx-top/image/webp/testdata/1586766691290258.png`)
	err = wm.MarkFile(destinationFile)
	if err != nil {
		t.Fatal(err)
	}
}