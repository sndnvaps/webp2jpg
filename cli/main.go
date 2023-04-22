package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/sndnvaps/webp2jpg"
	"github.com/urfave/cli/v2"
)

func Convert(ctx *cli.Context) error {
	Type := ctx.String("type")
	src := ctx.String("source")
	if Type == "" || src == "" {
		fmt.Println("You must specify a type")
		cli.ShowAppHelpAndExit(ctx, -1)
	}

	//fmt.Printf("source file = %s\n", src)

	filenameOnly := webp2jpg.RemovePathExt(src)
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
		fmt.Println(text)
	}

	img, _ := webp2jpg.Decode(src)
	/*
		if err != nil {
			fmt.Println(err.Error())
			cli.ShowAppHelp(ctx)
		}
	*/
	return webp2jpg.Encode(img, NewFileName, Type)
}

func main() {

	app := cli.NewApp()
	app.Name = "webp2jpg"
	app.Usage = "convert webp image to [bmp|gif|jpeg|png|tiff]"
	app.Version = "0.1.0"
	app.Authors = []*cli.Author{
		{
			Name:  "sndnvaps",
			Email: "sndnvaps@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "Convert to the type of image,support bmp,gif,jpeg,png,tiff",
		},
		&cli.StringFlag{
			Name:    "source",
			Aliases: []string{"s"},
			Usage:   "The file to convert,look like test.webp",
		},
	}

	app.Action = Convert

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Run(os.Args)
}
