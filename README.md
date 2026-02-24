# tossinvest

Go client for wts.tossinvest.com

This library requires http.Client that can handle authentication. Check 
[xtdlib/chttp](https://github.com/xtdlib/chttp).

API can change without notice. Please check the code for details.

## Example

```go
package main

import (
    "context"
    "log"

    "github.com/investing-kr/tossinvest"
    "github.com/davecgh/go-spew/spew"
    "github.com/xtdlib/chttp"
)

func main() {
    wts := tossinvest.NewClient(chttp.NewHTTPClient("ws://localhost:9222"))

    resp, _ := wts.X.Buy("AAPL", 50, 1)
    spew.Dump(resp)
    // (*tossinvest.OrderCreateDirectResponse)(0xc00058c930)({
    //  Message: (string) (len=27) "애플 구매 주문 완료",
    //  OrderDate: (string) (len=10) "2026-02-25",
    //  OrderNo: (int) 21,
    //  IsReserved: (bool) false
    // })
}
```

## wts

Simple CLI to demonstrate the usage of the library. 

### Installation

```bash
go install github.com/investing-kr/tossinvest/wts@latest
```

### Usage

```bash
# login with chrome
google-chrome --remote-debugging-port=9222 --user-data-dir=/tmp/chrome https://tossinvest.com

# AAPL 1주 매수
wts buy AAPL 240.9 1

# 삼성전자(A005930) 1주 매도
wts sell A005930 150000 1
```


