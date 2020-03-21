package webp

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestWebp(t *testing.T) {
	f, err := os.Open(`testdata/58.jpg`)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := Encode(f, 70)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(`testdata/58.jpg.webp`, b.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}