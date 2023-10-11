# carve

Go template repository.

```
git grep -l 'carve' | xargs sed -i 's/carve/your_repo/g'
git grep -l 'kijimaD' | xargs sed -i 's/kijimaD/your_name/g'
```

## install

```
$ go install github.com/kijimaD/carve@main
```

## docker run

```
$ docker run -v "$PWD/":/work -w /work --rm -it ghcr.io/kijimad/carve:latest
```
