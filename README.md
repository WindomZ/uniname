# uniname - Unique Name

> A CLI tool - Rename file to unique name quickly

[![Build Status](https://travis-ci.org/WindomZ/uniname.svg?branch=master)](https://travis-ci.org/WindomZ/uniname)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/uniname)](https://goreportcard.com/report/github.com/WindomZ/uniname)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Usage
```
Usage:
  uniname [-r] [--sha1|--sha256|--sha512] <input file>
  uniname [--sha1|--sha256|--sha512] <input file> -d <output directory>

Example:
  uniname -r foo.png
  uniname -r --sha256 foo.png
  uniname --sha512 foo.png -d foo/images

Optional flags:
  -d string
        rename to the specified directory
  -h    print help
  -md5
        using md5sum (default true)
  -r    rename the input file
  -sha1
        using sha1sum
  -sha256
        using sha256sum
  -sha512
        using sha512sum
  -v    print version
```

## Examples
```bash
uniname -r foo.png                    # md5sum, rename file.
uniname -r --sha1 foo.png             # sha1sum, rename file.
uniname --sha256 foo.png              # sha256sum, do nothing.
uniname --sha512 foo.png -d ./images  # sha512sum, create file to ./images.
```

## Install
If you have a Golang environment:
```bash
go get -u github.com/WindomZ/uniname
```

Or download the latest binary [release](https://github.com/WindomZ/uniname/releases)

## Changelog
See [CHANGELOG.md](https://github.com/WindomZ/uniname/blob/master/CHANGELOG.md#readme)

## Contributing
Welcome to pull requests, report bugs, suggest ideas and discuss, 
i would love to hear what you think on [issues page](https://github.com/WindomZ/uniname/issues).

If you like it then you can put a :star: on it.

## License
[MIT](https://github.com/WindomZ/uniname/blob/master/LICENSE)
