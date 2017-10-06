// Declaration of response types from Daum search REST API.
// Refer to the following URL for a full list of types and examples.
// https://developers.kakao.com/docs/restapi/search
package daumcrawler

// All "Response" structs are composed of Meta and Documents where
// only Documents differ in types. Meta is used on all "Response" types.
type Meta struct {
	TotalCount int
	PageableCount int
	IsEnd bool
}

type WebDocument struct {
	Title string
	Contents string
	URL string
	Datetime string
}

type WebResponse struct {
	Meta Meta
	Documents []WebDocument
}

type VclipDocument struct {
	Title string
	URL string
	Datetime string
	Playtime int
	Thumbnail string
	Author string
}

type VclipResponse struct {
	Meta Meta
	Documents []VclipDocument
}

type ImageDocument struct {
	Collection string
	ThumbnailURL string
	ImageURL string
	Width int
	Height int
	DisplaySitename string
	DocURL string
	Datetime string
}

type ImageResponse struct {
	Meta Meta
	Documents []ImageDocument
}

type BlogDocument struct {
	Title string
	Contents string
	URL string
	Blogname string
	Thumbnail string
	Datetime string
}

type BlogResponse struct {
	Meta Meta
	Documents []BlogDocument
}

type TipDocument struct {
	Title string
	Contents string
	QURL string
	AURL string
	Thumbnails []string
	Type string
	Datetime string
}

type TipResponse struct {
	Meta Meta
	Documents []TipDocument
}

type BookDocument struct {
	Title string
	Contents string
	URL string
	ISBN string
	Datetime string
	Authors []string
	Publisher string
	Translators []string
	Price string
	SalePrice string
	SaleYN string
	Category string
	Thumbnail string
	Barcode string
	EbookBarcode string
	Status string
}

type BookResponse struct {
	Meta Meta
	Documents []BookDocument
}

type CafeDocument struct {
	Title string
	Contents string
	URL string
	Cafename string
	Thumbnail string
	Datetime string
}

type CafeResponse struct {
	Meta Meta
	Documents []CafeDocument
}
