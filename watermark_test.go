package image

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestWatermark(t *testing.T) {
	watermarkFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/admpub/nging/public/assets/backend/images/nging-gear.png`)
	wm, err := NewWatermark(watermarkFile, 0, TopRight)
	if err != nil {
		t.Fatal(err)
	}
	srcFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/webx-top/image/webp/testdata/src.jpeg`)
	destinationFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/webx-top/image/webp/testdata/src_marked.jpeg`)
	err = wm.MarkFile(srcFile, destinationFile)
	if err != nil {
		t.Fatal(err)
	}

}

func TestRemoteWatermark(t *testing.T) {
	for i := 0; i < 3; i++ {
		watermarkFile := `https://img.nging.admpub.com/public/upload/default/2021/05/18/143523149248462848.png`
		wm, err := NewWatermark(watermarkFile, 0, TopRight)
		if err != nil {
			t.Fatal(err)
		}
		wm.debug = true
		srcFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/webx-top/image/webp/testdata/src.jpeg`)
		destinationFile := filepath.Join(os.Getenv("GOPATH"), `src/github.com/webx-top/image/webp/testdata/src_marked_`+strconv.Itoa(i)+`.jpeg`)
		fmt.Println(destinationFile)
		err = wm.MarkFile(srcFile, destinationFile)
		if err != nil {
			t.Fatal(err)
		}
	}
}
