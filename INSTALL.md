# Samplr Installation

_:warning: This will override any `/usr/local/bin/samplr` pre-existing file, which means the same command will be used to update._

## Mac OS

```sh
curl https://github.com/unmultimedio/samplr/releases/download/v0.2.1/samplr-v0.2.1-mac \
  -L -o /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

## Linux

```sh
curl https://github.com/unmultimedio/samplr/releases/download/v0.2.1/samplr-v0.2.1-linux \
  -L -o /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

## From source

You don't want to mess with your secrets files, that's ok. **This tool does not record or upload information in any way**, binaries are automatically built and added to the releases using Github actions as defined [here](.github/workflows/release.yml), but if you want to make extra sure:

- Install [go](https://golang.org/dl/).
- Clone this repo and `cd` into it.
- Download dependencies with `go mod vendor`.
- Go-install it with `go install` (this places it in `$GOPATH/bin`).
- (Optional) Build it with `go build` an and make it accessible somewhere in your `$PATH`.
- The command `samplr` should be available now, enjoy!
