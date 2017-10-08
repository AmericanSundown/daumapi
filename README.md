# daumapi

A simple wrapper around Daum "Search" REST API.


## Installation
```bash
go get -u github.com/hyunchel/daumapi
```

## Usage
```go
import (
    "fmt"

    "github.com/hyunchel/daumapi"
)

const API_KEY = "some api key from Kakao"

keyword := "볼빨간사춘기"

daumapi.Web(API_KEY, keyword)
daumapi.Vclip(API_KEY, keyword)
daumapi.Image(API_KEY, keyword)
daumapi.Blog(API_KEY, keyword)
daumapi.Tip(API_KEY, keyword)
daumapi.Book(API_KEY, keyword)
daumapi.Cafe(API_KEY, keyword)
```

For the logs, use `PrintLog` function.
```go
// Runs daumapi.Cafe and print all captured logs.
fmt.Println(daumapi.PrintLog(daumapi.Cafe, API_KEY, keyword))
```

## Reference
* [Daum Kakao REST API][daum-kakao-rest-api]


[daum-kakao-rest-api]: https://developers.kakao.com/docs/restapi/search