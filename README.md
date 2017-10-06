# daumcrawler

A crawler for Daum "Search" REST API.


## Installation
```bash
go get -u github.com/hyunchel/daumcrawler
```

## Usage
```go
import (
    "fmt"

    "github.com/hyunchel/daumcrawler"
)

keyword := "볼빨간사춘기"

fmt.Println(daumcrawler.Web(keyword))
fmt.Println(daumcrawler.Vclip(keyword))
fmt.Println(daumcrawler.Image(keyword))
fmt.Println(daumcrawler.Blog(keyword))
fmt.Println(daumcrawler.Tip(keyword))
fmt.Println(daumcrawler.Book(keyword))
fmt.Println(daumcrawler.Cafe(keyword))
```

## Reference
* [Daum Kakao REST API][daum-kakao-rest-api]


[daum-kakao-rest-api]: https://developers.kakao.com/docs/restapi/search