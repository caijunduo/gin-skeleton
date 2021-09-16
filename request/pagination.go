package request

type Pagination struct {
	Page    uint `json:"page" xml:"page" query:"page" form:"page" protobuf:"page" uri:"page" binding:"gte=1"`
	PerPage uint `json:"per_page" xml:"per_page" query:"per_page" form:"per_page" protobuf:"per_page" uri:"per_page" binding:"gte=1,lte=50"`
}

func (r *Pagination) Default() {
	r.Page = 1
	r.PerPage = 20
}

func (r Pagination) Validate() error {
	return VerifyStruct(r, Rules{
		"Page":    {PageGe1},
		"PerPage": {PageGe1, PerPageGe50},
	})
}
