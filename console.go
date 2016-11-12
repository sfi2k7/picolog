package picolog

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	fileName string
	file     *os.File
}

func (cl *ConsoleLogger) Println(data ...interface{}) {
	fmt.Fprintln(cl.file, data...)
}

func (cl *ConsoleLogger) Print(data ...interface{}) {
	fmt.Fprint(cl.file, data...)
}

func (cl *ConsoleLogger) Close() {
	if cl.file != nil {
		cl.file.Close()
	}
}

func NewConsole(prefix string) *ConsoleLogger {
	cl := &ConsoleLogger{}
	path := getLogPath(prefix)
	createIfNotExist(path)
	cl.fileName = path + prefix + ".log"
	f, err := os.Create(cl.fileName)
	if err != nil {
		fmt.Println(err)
	}
	cl.file = f
	return cl
}
