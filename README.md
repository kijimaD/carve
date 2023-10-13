# carve

carve is a command that rewrites the tag text in a file based on the git tag.

## worked example

```shell
$ git describe --tags --abbrev=0
v0.0.1
$ git tag v0.0.2 # add new tag
$ carve . pkg/version.go sample.yml # rewrite file versions
v0.0.1 -> v0.0.2
```

```diff
$ git diff

diff --git a/.versions b/.versions
index 95e94cd..90ab6e9 100644
--- a/.versions
+++ b/.versions
@@ -1 +1 @@
-v0.0.1
\ No newline at end of file
+v0.0.2
\ No newline at end of file
diff --git a/pkg/version.go b/pkg/version.go
index 06d51ad..6b5d02c 100644
--- a/pkg/version.go
+++ b/pkg/version.go
@@ -1,3 +1,3 @@
 package carve

-const Version = "v0.0.1"
+const Version = "v0.0.2"
diff --git a/sample.yml b/sample.yml
index 19979e6..3910b1e 100644
--- a/sample.yml
+++ b/sample.yml
@@ -2,7 +2,7 @@

 info:
   description: carve
-  version: v0.0.1
+  version: v0.0.2
   title: API Docs
   contact:
     name: kijimad
```

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
