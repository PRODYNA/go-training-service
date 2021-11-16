package data

type Result struct {
	Args    Args        `json:"args"`
	Data    string      `json:"data"`
	Files   Files       `json:"files"`
	Form    Form        `json:"form"`
	Headers Headers     `json:"headers"`
	JSON    interface{} `json:"json"`
	Origin  string      `json:"origin"`
	URL     string      `json:"url"`
}
type Args struct {
}
type Files struct {
}
type Form struct {
	OkTrue string `json:"{"ok": true}"`
}
type Headers struct {
	Accept        string `json:"Accept"`
	ContentLength string `json:"Content-Length"`
	ContentType   string `json:"Content-Type"`
	Host          string `json:"Host"`
	UserAgent     string `json:"User-Agent"`
	XAmznTraceID  string `json:"X-Amzn-Trace-Id"`
}