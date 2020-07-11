# Samplr

**Samplr** is a language-agnostic command that will generate updated commitable sample files, while you keep original ones git-ignored.

<a target="_blank" href="https://youtu.be/g85wHOZSdxU">
  See it in action (1m video)<br/>
  <img src="https://img.youtube.com/vi/g85wHOZSdxU/maxresdefault.jpg" alt="Samplr" width="400"/>
</a>

## Installation

Follow instructions [here](./INSTALL.md).

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

Setup your `.samplr.yml` configuration file at the root of your project. It uses regex matches for file paths. Some example configs [here](./examples/.samplr.sample.yml)

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

#### Keywords combinations

You can use different keywords in the same file. If you use different keywords in the same line, it will use this priority:

1. `#ssamplr#`
1. `#hsamplr#`
1. `#samplr#`

And if many same keywords in the same line, it will use the first ocurrence.

### 3. Samplr command

Run `samplr` at the root of your project. When run, it will scan all of your directory files, and will generate samples if:

- file path matches with the [configuration file](#1-samplr-config) settings.
- it includes at least one [samplr keyword](#2-setup-your-original-files).

### 4. Check the samples

Generated files does not include the `.sample` extension at the end of the file , but before the real extension (if exists), this helps editor linting and coloring.

| Original filename | Generated sample filename |
| ----------------- | ------------------------- |
| configuration.yml | configuration.sample.yml  |
| Dockerfile        | Dockerfile.sample         |

### 5. Set up a githook (optional)

Great, now you have autogenerated sample files, but you need to remember to run `samplr` everytime? Let's set up a githook, so in every commit this happens automatically. Create a file `.git/hooks/pre-commit` (or append this to it).

```sh
#!/bin/sh

set -e

# Run samplr command to generate sample files
samplr
# List all changed and not-ignored files, with a filename that matches with ".sample", and add it to the commit
git ls-files -mo --exclude-standard | grep "\.sample" | xargs git add
```

## Examples

Check examples for many type files [here](./examples).
