# gorar
Extract rar/zip files in Go.


# Install

```
go install https://github.com/154pinkchairs/gorar-ng@latest
```

# Usage

## Extract RAR
```go
RarExtractor("Unrarme.rar", "./")
```
**Multi Archive RAR is supported.**(Do not itirate,only pass first file)


## Extract Zip


```go
ZipExtractor("Unzipme.zip","./")
```

---

### Credits

[mholt/archiver](https://github.com/mholt/archiver) - error handling.

[nwaples/rardecode](https://github.com/nwaples/rardecode) - `rar` decoding library.

[jkotra/gorar](https://github.com/jkotra/gorar) - the original project, currently unmaintained


