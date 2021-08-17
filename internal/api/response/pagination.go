package response

type Pagination struct {
    Page int64 `json:"page"`
    PerPage int64 `json:"per_page"`
    ToTalCount int64 `json:"total_count"`
    Lists []interface{} `json:"lists"`
}
