# gorar
Extract rar/zip files in Go.


# Install

```
go get -v https://github.com/jagadeesh-kotra/gorar/
```

# Usage

## Extract RAR
```go
RarExtractor("Unrarme.rar", "./")
```

## Extract Zip


```go
ZipExtractor("Unzipme.zip","./")
```

## Credits

mholt/archiver - Error handling (Thx!)
nwaples/rardecode - rar decoding library (Thx!)



