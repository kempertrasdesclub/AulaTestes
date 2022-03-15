package restful

type Metadata struct {
	Success  bool     `json:"success"`
	Error    []string `json:"error"`
	Total    int      `json:"total"`
	Actual   int      `json:"actual"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
}
