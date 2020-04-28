# Samplr Installation

_:warning: This will override any `/usr/local/bin/samplr` pre-existing file._

For Mac OS and Linux, this will download the binary, place it in a `PATH` directory, and make it executable.

## Mac OS

```sh
wget https://github.com/unmultimedio/samplr/releases/download/v0.1.0/samplr-v0.1.0-mac \
  -O /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

## Linux

```sh
wget https://github.com/unmultimedio/samplr/releases/download/v0.1.0/samplr-v0.1.0-linux \
  -O /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

## From source

You don't want to mess with your secrets files, that's ok. **This tool does not record or upload information in any way**, binaries are automatically built and added to the releases using Github actions as defined [here](.github/workflows/release.yml), but if you want to make extra sure:

- Install [go](https://golang.org/dl/).
- Clone this repo and `cd` into it.
- Do `go mod vendor` to download dependencies.
- Do `go install` to build the executable, and place it in the `GOBIN` path.
- The command `samplr` should be available now, enjoy!
