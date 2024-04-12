package webp

import (
	"os"
	"testing"
)

// go test --tags=libwebp -v --count=1 -run ./...
// go test --tags=vips -v --count=1 -run ./...
func TestWebp(t *testing.T) {
	f, err := os.Open(`testdata/src.jpeg`)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := Encode(f, 70)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(`testdata/src.jpeg.webp`, b.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
