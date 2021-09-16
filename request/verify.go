package request

var (
	PageGe1     = Ge("1", "页数不能小于1")
	PerPageGe50 = Le("50", "每页数量不能大于50")
)
