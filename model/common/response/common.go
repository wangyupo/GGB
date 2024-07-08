package response

type PageResult struct {
	List  interface{} `json:"list" swaggertype:"array,object"`
	Total int64       `json:"total"`
}
