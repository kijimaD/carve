# carve

carve is a command that rewrites the tag text in a file based on the git tag.

## install

```
$ go install github.com/kijimaD/carve@main
```

## usage

```shell
carve . pkg/version.go sample.yml
```

carve [repo] [oldtext] [target files...]

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/carve:latest
```
