# carve

carve is a command that rewrites the tag text in a file based on the git tag.

## use

```shell
carve . v0.0.1 constants.go openapi.yml
```

carve [repo] [oldtext] [target files...]

## install

```
$ go install github.com/kijimaD/carve@main
```

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/carve:latest
```
