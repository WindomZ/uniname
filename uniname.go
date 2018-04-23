package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var version string

var (
	optHelp    bool
	optVersion bool
	optReplace bool
	optMD5     bool
	optSha1    bool
	optSha256  bool
	optSha512  bool
)

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

func displayUsage() {
	name := commandName()
	display(fmt.Sprintf("Usage:\n  %s [-r] [--md5|--sha1|--sha256|--sha512] <file path>", name))
	display(fmt.Sprintf("\nExample:\n  %s -r foo.png", name))
	display("\nOptional flags:")
	flag.PrintDefaults()
}

func displayVersion() {
	display(fmt.Sprintf("%s version %s", commandName(), version))
}

func init() {
	flag.Usage = displayUsage
}

func main() {
	flag.BoolVar(&optHelp, "h", false, "print help")
	flag.BoolVar(&optVersion, "v", false, "print version")
	flag.BoolVar(&optReplace, "r", false, "rename the input file")
	flag.BoolVar(&optMD5, "md5", true, "using md5sum")
	flag.BoolVar(&optSha1, "sha1", false, "using sha1sum")
	flag.BoolVar(&optSha256, "sha256", false, "using sha256sum")
	flag.BoolVar(&optSha512, "sha512", false, "using sha512sum")
	flag.Parse()

	if optVersion {
		displayVersion()
		os.Exit(0)
	}
	if optHelp || flag.NArg() == 0 {
		displayUsage()
		os.Exit(0)
	}

	srcFilePath := flag.Arg(0)

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
	case optSha1:
		mode = "sha1"
	case optSha256:
		mode = "sha256"
	case optSha512:
		mode = "sha512"
	default:
		mode = "md5"
	}

	dstFilePath, err := fileSum("sha1", srcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	dstFilePath = filepath.Join(filepath.Dir(srcFilePath), dstFilePath+filepath.Ext(srcFilePath))

	if optReplace {
		if err = os.Rename(srcFilePath, dstFilePath); err != nil {
			log.Fatal(err)
		}
		display(fmt.Sprintf("%s: '%s' rename to '%s'", mode, srcFilePath, dstFilePath))
	} else {
		display(fmt.Sprintf("%s: '%s' unique name is '%s'", mode, srcFilePath, dstFilePath))
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
