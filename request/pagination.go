package request

type Pagination struct {
	Page    uint `json:"page" xml:"page" query:"page" form:"page" protobuf:"page" uri:"page" binding:"gte=1"`
	PerPage uint `json:"per_page" xml:"per_page" query:"per_page" form:"per_page" protobuf:"per_page" uri:"per_page" binding:"gte=1,lte=50"`
}

func (p *Pagination) Default() *Pagination {
	p.Page = 1
	p.PerPage = 20
	return p
}

func (p Pagination) Validate() error {
	return VerifyStruct(p, Rules{
		"Page":    {PageGe1},
		"PerPage": {PageGe1, PerPageGe50},
	})
}
