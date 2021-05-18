package webp

import (
	"io/ioutil"
	"os"
	"testing"
)

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
	err = ioutil.WriteFile(`testdata/src.jpeg.webp`, b.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
