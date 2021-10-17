package request

import (
    "skeleton/validation"
)

type Pagination struct {
    Page    uint `json:"page" xml:"page" query:"page" form:"page" protobuf:"page" uri:"page"`
    PerPage uint `json:"per_page" xml:"per_page" query:"per_page" form:"per_page" protobuf:"per_page" uri:"per_page"`
}

func (p Pagination) Validate() error {
    return validation.Verify(
        validation.Field(p.Page).
            Rule(validation.Required.SetMessage("页数不能为空")),
        validation.Field(p.PerPage).
            Rule(validation.Required.SetMessage("每页数量不能为空")).
            Rule(validation.Max(50).SetMessage("每页数量不能超过50")),
    )
}
