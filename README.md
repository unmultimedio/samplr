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

We all `.gitignore` files in our repos. Many of those for security purposes, like configuration files with secrets or frequently-changing URLs that we don't want to commit and clutter in every PR everytime we change them.

**So, what do we normally do?** We make a copy of the original file, and append a `.sample` to it. This serves the purpose for new repo clones to get a configuration file structure, and we just need to fill the secrets or URLs locally.

<table>
<tr>
<th> Actual file </th> <th> File that we commit </th> <th> So we can do </th>
</tr>

<tr valign="top">
<td>

```yml
# configuration.yml
some: variables
url: http://aws.url/changes.frequently
foo: bar
token: MY.S3CUR3.T0K3N
```

</td>
<td>

```yml
# configuration.yml.sample
some: variables
url: REPLACE_ME_DEV_URL
foo: bar
token: REPLACE_ME_TOKEN
```

</td>
<td>

```gitignore
# .gitignore
configuration.yml
```

</td>
</tr>
</table>


**And what's the issue with that?** It gets outdated really fast! When was the last time you found missing or renamed configs and secrets when onboarding someone?

**Ok, so what's the idea?** I don't like making manual changes to the `.sample` file everytime I change something in the original. I know you don't either. Sometimes we just forget. What about if we could only edit the original one and the samples get updated automatically?

<img src="https://i.kym-cdn.com/entries/icons/facebook/000/031/991/cover3.jpg" alt="You son of a bitch, I'm in" width="400"/>

## Usage

### 1. Samplr config

Setup your `.samplr.yml` configuration file at the root of your project. It uses regex matches for file paths. This one is a good start:

```yml
# Example of .samplr.yml

includes:
- \.yaml$
- \.yml$

excludes:
- \.sample
```

You can find more sample configs [here](./examples/.samplr.sample.yml).

> For a file to be sampled, it has to meet **both** requirements, be included **and** not excluded.

### 2. Setup your original files

Use the special keywords in your original file, and samplr will take care of autogenerate the sample files.

> Important: you still need to manually `.gitignore` your original file.

#### Keyword #samplr#

Regular use case for team collaboration. It will replace the next line with the content right after. The samplr line itself will also be rendered.

<table>
<tr>
<th> Original file </th> <th> Sample file </th>
</tr>

<tr valign="top">
<td>

```yml
# configuration.yml
some: variables
#samplr#token: REPLACE_ME_TOKEN
token: MY.S3CUR3.T0K3N
foo: bar
```

</td>
<td>

```yml
# configuration.sample.yml
some: variables
#samplr#token: REPLACE_ME_TOKEN
token: REPLACE_ME_TOKEN
foo: bar
```

</td>
</tr>
</table>

#### Keyword #hsamplr#

As in "hide this samplr comment". It will replace the next line with the content right after. The samplr line itself will **not** be rendered.

<table>
<tr>
<th> Original file </th> <th> Sample file </th>
</tr>

<tr valign="top">
<td>

```yml
# configuration.yml
some: variables
#hsamplr#token: REPLACE_ME_TOKEN
token: MY.S3CUR3.T0K3N
foo: bar
```

</td>
<td>

```yml
# configuration.sample.yml
some: variables
token: REPLACE_ME_TOKEN
foo: bar
```

</td>
</tr>
</table>

#### Keyword #ssamplr#

As in "this is a secret samplr comment". It will hide the samplr line itself. The next line will be rendered normally.

<table>
<tr>
<th> Original file </th> <th> Sample file </th>
</tr>

<tr valign="top">
<td>

```yml
# configuration.yml
some: variables
#ssamplr# secret token: MY.S3CUR3.T0K3N
foo: bar
```

</td>
<td>

```yml
# configuration.sample.yml
some: variables
foo: bar
```

</td>
</tr>
</table>

### 3. Samplr command

Run `samplr` at the root of your project. When run, it will scan all of your directory files, and will generate samples if:

- file path matches with the [configuration file](#1-samplr-command) settings.
- it includes at least one samplr keyword.

### 4. Check the samples

Sample files does not include the `.sample` at the end of the file (if exists), but before the real extension _(which helps editor linting and coloring)_.

| Original filename | Generated sample filename |
| ----------------- | ------------------------- |
| configuration.yml | configuration.sample.yml  |
| Dockerfile        | Dockerfile.sample         |
