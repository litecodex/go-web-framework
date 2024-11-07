package page

type PageResult[T any] struct {
	Total    int64 `json:"total"`
	PageData []T   `json:"pageData"`
}

func NewPageResult[T any](total int64, pageData []T) *PageResult[T] {
	if pageData == nil {
		pageData = []T{}
	}
	return &PageResult[T]{
		Total:    total,
		PageData: pageData,
	}
}
