package webp2jpg

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
)

// Encode(img Image.Image, filename,Type string)
/*
 * Encode image.Image stream to Type
 * Type support : bmp,gif,jpeg,jpg,png,tiff
 * filename: save name of the file
 */
func Encode(img image.Image, filename, Type string) error {
	fw, _ := os.Create(filename)
	defer fw.Close()

	switch Type {
	case "bmp":
		bmp.Encode(fw, img)
	case "gif":
		gif.Encode(fw, img, nil)
	case "jpeg", "jpg":
		jpeg.Encode(fw, img, nil)
	case "png":
		png.Encode(fw, img)
	case "tiff":
		tiff.Encode(fw, img, nil)
	default:
		text := fmt.Sprintf("The type:[%s] not in support list", Type)
		err := errors.New(text)
		return err
	}

	fmt.Printf("Convert %s success\n", filename)

	return nil
}

//func Decode(filename string) (img image.Image,err error)
/* filename: origin of *.webp file
 * img: what we decode *.webp to
 * err: error info
 */
func Decode(filename string) (img image.Image, err error) {
	if !strings.HasSuffix(filename, ".webp") {
		err = errors.New("filename not contain *.webp")
		return nil, err
	}
	f, _ := os.Open(filename)
	defer f.Close()

	img, _ = webp.Decode(f)
	return img, nil
}
