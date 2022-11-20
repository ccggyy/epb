# EPB

> EPB is short for **E Package Builder**

## Install

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/ccggyy/epb/HEAD/install.sh)"
```

Or manual download binary from [release](https://github.com/ccggyy/epb/releases).

## Using

- `-v` needs increase part of version, options: `leading`, `center`, `trailing`.
- `-f` config filename
- `-p` bundle path
- `-d` dist path
- `-o` output path

```shell
epb \
-v trailing \
-f package.json \
-p /path/to/bundle/ \
-d /path/to/dist/ \
-o /path/to/output/ 
```

## Build snapshot with GoReleaser

Before run following command, make sure you have been installed `goreleaser` on your machine.

```shell
goreleaser build --snapshot --rm-dist
```

## Release with GoReleaser

```shell
goreleaser release --rm-dist
```
