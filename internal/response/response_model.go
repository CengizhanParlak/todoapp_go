package response

type Response struct {
	Success  bool        `json:"success"`
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}
