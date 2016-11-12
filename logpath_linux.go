package picolog

func getLogPath(prefix string) string {
	p := "/var/" + prefix + "/logs/"

	return p
}
