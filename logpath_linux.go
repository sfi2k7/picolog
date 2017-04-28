package picolog

func getLogPath(prefix string) string {
	p := "/var/apps/logs/" + prefix + "/"

	return p
}
