package picolog

import (
	"os"
)

func getLogPath(prefix string) string {
	p := "c:\\db\\apps\\" + prefix + "\\logs\\"
	if _, err := os.Stat(p); err != nil {
		os.MkdirAll(p, 777)
	}
	return p
}
