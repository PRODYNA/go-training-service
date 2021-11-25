package data

type Result struct {
	Args    Args    `json:"args"`
	Data    string  `json:"data"`
	Files   Files   `json:"files"`
	Form    Form    `json:"form"`
	Headers Headers `json:"headers"`
	JSON    JSON    `json:"json"`
	Origin  string  `json:"origin"`
	URL     string  `json:"url"`
}

type Args struct {
}

type Files struct {
}

type Form struct {
}

type Headers struct {
	Accept          string `json:"Accept"`
	AcceptEncoding  string `json:"Accept-Encoding"`
	AcceptLanguage  string `json:"Accept-Language"`
	Authorization   string `json:"Authorization"`
	CacheControl    string `json:"Cache-Control"`
	ContentLength   string `json:"Content-Length"`
	ContentType     string `json:"Content-Type"`
	Host            string `json:"Host"`
	Origin          string `json:"Origin"`
	PostmanToken    string `json:"Postman-Token"`
	SecChUa         string `json:"Sec-Ch-Ua"`
	SecChUaMobile   string `json:"Sec-Ch-Ua-Mobile"`
	SecChUaPlatform string `json:"Sec-Ch-Ua-Platform"`
	SecFetchDest    string `json:"Sec-Fetch-Dest"`
	SecFetchMode    string `json:"Sec-Fetch-Mode"`
	SecFetchSite    string `json:"Sec-Fetch-Site"`
	UserAgent       string `json:"User-Agent"`
	XAmznTraceID    string `json:"X-Amzn-Trace-Id"`
}

type JSON struct {
	Key string `json:"key"`
}
