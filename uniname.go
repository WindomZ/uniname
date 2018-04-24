package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var version string

var (
	optHelp      bool
	optVersion   bool
	optRename    bool
	optOutputDir string
	optMD5       bool
	optSha1      bool
	optSha256    bool
	optSha512    bool
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
	display(fmt.Sprintf("Usage:\n  %[1]s [-r] [--sha1|--sha256|--sha512] <input file>"+
		"\n  %[1]s [--sha1|--sha256|--sha512] <input file> -d <output directory>", name))
	display(fmt.Sprintf("\nExample:\n  %[1]s -r foo.png"+
		"\n  %[1]s -r --sha256 foo.png\n  %[1]s --sha512 foo.png -d foo/images", name))
	display("\nOptional flags:")
	flag.PrintDefaults()
}

func displayVersion() {
	display(fmt.Sprintf("%s version %s", commandName(), version))
}

func init() {
	flag.Usage = displayUsage
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

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	if err = dstFile.Sync(); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.BoolVar(&optHelp, "h", false, "print help")
	flag.BoolVar(&optVersion, "v", false, "print version")
	flag.BoolVar(&optRename, "r", false, "rename the input file")
	flag.StringVar(&optOutputDir, "d", "", "rename to the specified directory")
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
			log.Fatal(err)
		}
		log.Fatal(err)
	} else if f.IsDir() {
		log.Fatal(os.ErrNotExist)
	}

	if optOutputDir != "" {
		if optOutputDir == "." {
			optRename = true
			optOutputDir = ""
		} else if f, err := os.Stat(optOutputDir); err != nil {
			if err = os.MkdirAll(optOutputDir, 0700); err != nil {
				log.Fatal(err)
			}
		} else if !f.IsDir() {
			optOutputDir = filepath.Dir(optOutputDir)
		}
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

	dstFileName, err := fileSum(mode, srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	dstFileName += filepath.Ext(srcFilePath)

	if optRename {
		dstFilePath := filepath.Join(filepath.Dir(srcFilePath), dstFileName)
		if err = os.Rename(srcFilePath, dstFilePath); err != nil {
			log.Fatal(err)
		}
		display(fmt.Sprintf("%s: '%s' rename to '%s'", mode, srcFilePath, dstFileName))
	} else if optOutputDir != "" {
		dstFilePath := filepath.Join(optOutputDir, dstFileName)
		if err = copyFile(srcFilePath, dstFilePath); err != nil {
			log.Fatal(err)
		}
		display(fmt.Sprintf("%s: '%s' copy to '%s'", mode, srcFilePath, dstFilePath))
	} else {
		display(fmt.Sprintf("%s: unique name is '%s'", mode, dstFileName))
	}
}
