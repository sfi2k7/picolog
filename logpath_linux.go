package picolog

func getLogPath(prefix string) string {
	p := "/var/apps/" + prefix + "/logs/"

	return p
}
