package page

type PageQueryParam struct {
	PageNo   int `form:"pageNo"`   // 当前页码
	PageSize int `form:"pageSize"` // 每页显示条数
}

func (thiz *PageQueryParam) GetPageNo() int {
	if thiz.PageNo <= 0 {
		return 1
	}
	return thiz.PageNo
}

func (thiz *PageQueryParam) GetPageSize() int {
	if thiz.PageSize <= 0 {
		return 10
	}
	if thiz.PageSize > 5000 {
		return 5000 // 防止查数据太多，拖垮数据库
	}
	return thiz.PageSize
}

func (thiz *PageQueryParam) GetOffset() int {
	return (thiz.GetPageNo() - 1) * thiz.GetPageSize()
}

func (thiz *PageQueryParam) GetLimit() int {
	return thiz.GetPageSize()
}
