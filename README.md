# Samplr

Autogenerate `.sample` files (and keep them updated!) after you `.gitignore` your real ones.

## Installation

_:warning: This will override any `/usr/local/bin/samplr` pre-existing file._

For Mac OS and Linux, this will download the binary, place it in a `PATH` directory, and make it executable.

### Mac OS

```sh
wget https://github.com/unmultimedio/samplr/releases/download/v0.1.0/samplr-v0.1.0-mac \
  -O /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

### Linux

```sh
wget https://github.com/unmultimedio/samplr/releases/download/v0.1.0/samplr-v0.1.0-linux \
  -O /usr/local/bin/samplr && \
  chmod +x /usr/local/bin/samplr
```

### From source

You don't want to mess with your secrets files, that's ok. **This tool does not record or upload information in any way**, binaries are automatically built and added to the releases using Github actions as defined [here](.github/workflows/release.yml), but if you want to make extra sure:

- Install [go](https://golang.org/dl/).
- Clone this repo and `cd` into it.
- Do `go mod vendor` to download dependencies.
- Do `go install` to build the executable, and place it in the `GOBIN` path.
- The command `samplr` should be available now, enjoy!

## Why?

We all have `.gitignored` files in our repos. Many of those for security purposes, like configuration files with secrets or frequently-changing URLs that we don't want to change, commit and clutter in every PR.

**So, what do we normally do?** We make a copy of the original file, and append a `.sample` to it. This serves the purpose for new repo clones to get a configuration file structure, and we just need to fill the secrets or URLs locally.

Our real file:

```yml
# configuration.yml
some: variables
for: local
or: stage
development: that
contains: secrets
like: VERY_SECURE_PASSWORD
```

Sample file that we commit:

```yml
# configuration.yml.sample
some: variables
for: local
or: stage
development: that
contains: secrets
like: <INSERT_PASS>
```

So we can gitignore the real one:

```gitignore
# .gitignore
configuration.yml
```

**And what's the issue with that?** With time we modify our code and settings in the real file, and it's not so long before we realize that we have a very old `.sample` file that poorly reflects the real one. That complicates new clones or new peers onboardings.

**Ok, so what's the idea?** I don't like making manual changes to the `.sample` file everytime I change something in the original file. I know you don't either. Sometimes we just forget. What about if we could only edit the original one and autogenerate the sample?

## How it works?

![alt text](https://i.kym-cdn.com/entries/icons/facebook/000/031/991/cover3.jpg "You son of a bitch, I'm in")

You need to use the special keyword `#samplr#` to obscure your secret lines, and tell samplr what it should use to replace it.

So if you have

```yml
# configuration.yml
public: foo
secret: VERY_SECURE_PASSWORD
```

And you want

```yml
# configuration.sample.yml
public: foo
secret: <INSERT_PASS>
```

You need to do

```yml
# configuration.yml
public: foo
#samplr#secret: <INSERT_PASS>
secret: VERY_SECURE_PASSWORD
```

Focus on the line:

```yml
#samplr#secret: <INSERT_PASS>
```

That line will print everything after the keyword and will ignore the immediate following line.

> Important: you still need to manually `.gitignore` your original file.

## Samplr command

When run, it will scan all of your directory files, and will generate sample files for files that meet these requirements:
- file path matches with the [configuration file](.samplr.yml) regex matches list
- it includes the `#samplr#` keyword at least once

## Sample extension

Generated sample files does not include the `.sample` at the end of the file, but before the real extension. This helps editor linting and coloring. This means your file `configuration.yml` will generate a sample `configuration.sample.yml` instead of `configuration.yml.sample`. Files with no extension will have `.sample` appended at the end like `Dockerfile` to `Dockerfile.sample`.
