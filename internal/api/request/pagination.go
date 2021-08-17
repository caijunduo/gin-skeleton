package request

type Pagination struct {
    Page    int64 `json:"page" xml:"page" query:"page" form:"page" protobuf:"page" uri:"page" binding:"gte=1"`
    PerPage int64 `json:"per_page" xml:"per_page" query:"per_page" form:"per_page" protobuf:"per_page" uri:"per_page" binding:"gte=1,lte=50"`
}

func (r *Pagination) DefaultPagination() {
    r.Page = 1
    r.PerPage = 20
}