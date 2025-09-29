package model

type UrlMap struct {
	Id       int    `json:"id"`
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}

type UrlMapPayload struct {
	LongUrl string
}
