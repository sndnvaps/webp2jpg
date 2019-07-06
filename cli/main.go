package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sndnvaps/webp2jpg"
	cli "gopkg.in/urfave/cli.v1"
)

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
	case "bmp":
		NewFileName = filenameOnly + ".bmp"
	case "gif":
		NewFileName = filenameOnly + ".gif"
	case "jpeg", "jpg":
		NewFileName = filenameOnly + ".jpg"
	case "png":
		NewFileName = filenameOnly + ".png"
	case "tiff":
		NewFileName = filenameOnly + ".tiff"
	default:
		text := fmt.Sprintf("The type:[%s] not in support list", Type)
		err := errors.New(text)
		return err
	}

	img, _ := webp2jpg.Decode(src)

	return webp2jpg.Encode(img, NewFileName, Type)
}

func main() {

	app := cli.NewApp()
	app.Name = "webp2jpg"
	app.Usage = "convert webp image to [bmp|gif|jpeg|png|tiff]"
	app.Version = "0.0.3"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "sndnvaps",
			Email: "sndnvaps@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "type,t",
			Usage: "Convert to the type of image,support bmp,gif,jpeg,png,tiff",
		},
		cli.StringFlag{
			Name:  "source,s",
			Usage: "The file to convert,look like test.webp",
		},
	}

	app.Action = Convert

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
