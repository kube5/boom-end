package model

type KindPage struct {
	Page
	Kind int64 `form:"kind" json:"kind"`
}

type Page struct {
	PageNo   int64 `form:"page_no" json:"page_no" binding:"required,gt=0"`
	PageSize int64 `form:"page_size" json:"page_size" binding:"required,gte=5,max=20"`
	Kind     int64 `form:"kind" json:"kind"`
}

func (req Page) ErrMessages() map[string]string {
	return map[string]string{
		"PageNo.required":   "page_no is empty",
		"PageNo.gt":         "required page_no > 0",
		"PageSize.gte":      "required page_size >= 5",
		"PageSize.required": "page_size is empty",
		"PageSize.max":      "maximum page_size length is 20",
	}

}
