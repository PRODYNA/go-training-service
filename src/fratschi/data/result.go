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
	Accept        string `json:"Accept"`
	CacheControl  string `json:"Cache-Control"`
	ContentLength string `json:"Content-Length"`
	ContentType   string `json:"Content-Type"`
	Host          string `json:"Host"`
	UserAgent     string `json:"User-Agent"`
	XAmznTraceID  string `json:"X-Amzn-Trace-Id"`
}
type JSON struct {
	Test bool `json:"test"`
}

