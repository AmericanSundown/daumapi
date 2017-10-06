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

fmt.Println(daumapi.Web(API_KEY, keyword))
fmt.Println(daumapi.Vclip(API_KEY, keyword))
fmt.Println(daumapi.Image(API_KEY, keyword))
fmt.Println(daumapi.Blog(API_KEY, keyword))
fmt.Println(daumapi.Tip(API_KEY, keyword))
fmt.Println(daumapi.Book(API_KEY, keyword))
fmt.Println(daumapi.Cafe(API_KEY, keyword))
```

## Reference
* [Daum Kakao REST API][daum-kakao-rest-api]


[daum-kakao-rest-api]: https://developers.kakao.com/docs/restapi/search