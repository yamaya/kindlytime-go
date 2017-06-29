kindlytime go library
====

Human kindly date-time library for go.

## Description

`kindlytime` is date-time library for go.

## Usage

```
import (
    "github.com/yamaya/kindlytime"
    "fmt"
)

jst, _ := time.LoadLocation("Asia/Tokyo")
origin := time.Date(2016, 12, 30, 0, 0, 0, 0, jst)

result := kindlytime.Parse("2 days ago", origin)
// 2016-12-28 00:00:00
```

## Requirement

go 1.8.3

## Install

```
$ go get github.com/yamaya/kindlytime
```

## TODO

- [ ] Convert time.Time object to kindlytime string
- [ ] Localization
- [ ] Test! Test! Test!

## Licence

* [MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

* [Masayuki Yamaya](https://github.com/yamaya)

## Contributors

* [Ryosuke Igarashi](https://github.com/cm-igarashi-ryosuke)
