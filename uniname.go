package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

var version string

func display(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func commandName() (name string) {
	name = filepath.Base(os.Args[0])
	if name == "" {
		name = "uniname"
	}
	return
}

func main() {
	app := cli.NewApp()
	app.Name = commandName()
	app.Usage = "rename file to unique name quickly!"
	app.UsageText = fmt.Sprintf("%s [-r] [--md5|--sha1|--sha256|--sha512] <file path>", app.Name)
	app.Version = version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "replace, r",
			Usage: "rename the input file",
		},
		cli.BoolTFlag{
			Name:  "md5",
			Usage: "using md5sum",
		},
		cli.BoolFlag{
			Name:  "sha1",
			Usage: "using sha1sum",
		},
		cli.BoolFlag{
			Name:  "sha256",
			Usage: "using sha256sum",
		},
		cli.BoolFlag{
			Name:  "sha512",
			Usage: "using sha512sum",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}

		srcFilePath := c.Args().Get(0)

		if f, err := os.Stat(srcFilePath); err != nil {
			if os.IsNotExist(err) {
				log.Fatal(os.ErrNotExist)
			}
			log.Fatal(err)
		} else if f.IsDir() {
			log.Fatal(os.ErrNotExist)
		}

		var mode string
		switch {
		case c.Bool("sha1"):
			mode = "sha1"
		case c.Bool("sha256"):
			mode = "sha256"
		case c.Bool("sha512"):
			mode = "sha512"
		default:
			mode = "md5"
		}

		dstFilePath, err := fileSum("sha1", srcFilePath)
		if err != nil {
			log.Fatal(err)
		}

		dstFilePath = filepath.Join(filepath.Dir(srcFilePath), dstFilePath+filepath.Ext(srcFilePath))

		if c.Bool("r") {
			if err = os.Rename(srcFilePath, dstFilePath); err != nil {
				log.Fatal(err)
			}
			display(fmt.Sprintf("%s: '%s' rename to '%s'", mode, srcFilePath, dstFilePath))
		} else {
			display(fmt.Sprintf("%s: '%s' unique name is '%s'", mode, srcFilePath, dstFilePath))
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func fileSum(mode, file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	switch mode {
	case "sha1":
		h := sha1.New()
		h.Write(data)
		return strings.ToLower(hex.EncodeToString(h.Sum(nil))), nil
	case "sha256":
		h := sha256.New()
		h.Write(data)
		return strings.ToLower(hex.EncodeToString(h.Sum(nil))), nil
	case "sha512":
		h := sha512.New()
		h.Write(data)
		return strings.ToLower(hex.EncodeToString(h.Sum(nil))), nil
	default:
		h := md5.New()
		h.Write(data)
		return strings.ToLower(hex.EncodeToString(h.Sum(nil))), nil
	}
}
