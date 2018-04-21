# uniname - Unique Name

> A CLI tool - Rename file to unique name quickly

[![Build Status](https://travis-ci.org/WindomZ/uniname.svg?branch=master)](https://travis-ci.org/WindomZ/uniname)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/uniname)](https://goreportcard.com/report/github.com/WindomZ/uniname)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

## Usage
```
Usage:
  uniname [-r] [--md5|--sha1|--sha256|--sha512] <file path>

Example:
  uniname -r demo.png

Optional flags:
  -md5
        using md5sum, by default
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
uniname -r dst.png        # md5sum, rename file.
uniname -r --sha1 dst.png # sha1sum, rename file.
uniname --sha256 dst.png  # sha256sum, do nothing.
```

## Install
```bash
go get -u github.com/WindomZ/uniname
```

## License
[MIT](https://github.com/WindomZ/uniname/blob/master/LICENSE)
