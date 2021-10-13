package operate

func HideMobile(mobile string) string {
	return mobile[0:3] + "****" + mobile[7:]
}
