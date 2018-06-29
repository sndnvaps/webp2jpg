package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"strings"

	"golang.org/x/image/webp"
)

func main() {

	args := os.Args

	if args == nil || len(args) != 3 {
		usage()
		return
	}

	source := args[1]
	target := args[2]

	if !strings.HasSuffix(source, "webp") {
		usage()
		return
	}
	f, _ := os.Open(source)

	img, _ := webp.Decode(f)

	fw, _ := os.Create(target)

	jpeg.Encode(fw, img, nil)

}

func usage() {
	fmt.Printf("输入错误，请按照下面的格式输入: \n")
	fmt.Printf("使用: webp2jpg source_image.webp output_image\n")
}
