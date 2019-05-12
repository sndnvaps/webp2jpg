package main

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
	"golang.org/x/image/webp"

	cli "gopkg.in/urfave/cli.v1"
)

func Encode(img image.Image, filename, Type string) error {
	fw, _ := os.Create(filename)
	defer fw.Close()

	switch Type {
	case "jpeg", "jpg":
		jpeg.Encode(fw, img, nil)
	case "png":
		png.Encode(fw, img)
	case "gif":
		gif.Encode(fw, img, nil)
	case "bmp":
		bmp.Encode(fw, img)
	default:
		text := fmt.Sprintf("The type:[%s] not in support list", Type)
		err := errors.New(text)
		return err
	}

	fmt.Printf("Convert %s success\n", filename)

	return nil
}
func Convert(ctx *cli.Context) error {
	Type := ctx.String("type")
	src := ctx.String("source")

	//fmt.Printf("source file = %s\n", src)

	if !strings.HasSuffix(src, ".webp") {
		cli.ShowAppHelp(ctx)
	}

	filenameOnly := strings.TrimSuffix(src, ".webp")
	//fmt.Printf("filenameonly %s\n", filenameOnly)
	var NewFileName string

	switch Type {
	case "jpeg", "jpg":
		NewFileName = filenameOnly + ".jpg"
	case "png":
		NewFileName = filenameOnly + ".png"
	case "gif":
		NewFileName = filenameOnly + ".gif"
	case "bmp":
		NewFileName = filenameOnly + ".bmp"
	}

	f, _ := os.Open(src)
	defer f.Close()

	img, _ := webp.Decode(f)

	return Encode(img, NewFileName, Type)
}

func main() {

	app := cli.NewApp()
	app.Name = "webp2jpg"
	app.Usage = "convert webp image to [bmp|jpeg|png|gif]"
	app.Version = "0.0.2"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "sndnvaps",
			Email: "sndnvaps@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "type,t",
			Usage: "Convert to the type of image,support bmp,jpeg,png,gif",
		},
		cli.StringFlag{
			Name:  "source,s",
			Usage: "The file to convert,look like test.webp",
		},
	}

	app.Action = Convert

	app.Run(os.Args)
}
