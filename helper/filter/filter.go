package filterHelper

func NilAndCreate(s []string) []string {
	res := make([]string, 0)
	for _, v := range s {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}
