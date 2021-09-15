package response

type Pagination struct {
	Page       uint        `json:"page"`
	PerPage    uint        `json:"per_page"`
	ToTalCount uint64      `json:"total_count"`
	Lists      interface{} `json:"lists"`
}
