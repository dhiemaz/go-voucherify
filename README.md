# go-voucherify

[![Go Report Card](https://goreportcard.com/badge/github.com/dhiemaz/go-voucherify)](https://goreportcard.com/report/github.com/dhiemaz/go-voucherify)

Generate unique, coupon and voucher codes. Write purely in Go

Use cases: promo codes, loyalty coupons, gift vouchers, in-app purchases

This library originates from Voucherify.

# Motivation

There is no library as far as I have search for generating voucher or promo codes in Go.

# Inspiration

I took a lot of ideas from the [voucherify]http://generator.voucherify.io/#).

# Installation

This package is a go getable package.

``$ go get github.com/dhiemaz/go-voucherify``

# Example

```go
package main

import (
        "github.com/dhiemaz/go-voucherify"
	"fmt"
)

func main() {

     result := voucher.Generate("ABDS12")

     // print code generated //
     fmt.Println(result)
}
```

# LICENSE

The MIT License (MIT)

Copyright (c) 2016 Claudemiro

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


