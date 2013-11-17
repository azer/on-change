## on-change

A simpler interface to get notified on changes in given file/directory

```go
import (
  . "github.com/azer/on-change"
)

OnChange("foobar.txt", func () {
  // foobar.txt has been updated
})
```

## Install

```bash
$ go get github.com/azer/on-change
```
