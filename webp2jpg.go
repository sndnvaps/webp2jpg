package webp2jpg

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

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
		fmt.Println(text)
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

	f, _ := os.Open(filename)
	defer f.Close()

	Ext := filepath.Ext(filename)
	switch Ext {
	case ".bmp":
		img, err = bmp.Decode(f)
	case ".gif":
		img, err = gif.Decode(f)
	case "jpeg", "jpg":
		img, err = jpeg.Decode(f)
	case ".png":
		img, err = png.Decode(f)
	case ".tiff":
		img, err = tiff.Decode(f)
	case ".webp":
		img, err = webp.Decode(f)

	default:
		text := fmt.Sprintf("The type:[%s] not in support list", Ext)
		err := errors.New(text)
		return nil, err
	}

	return img, nil
}

func RemovePathExt(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return path
}
