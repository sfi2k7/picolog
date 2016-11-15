package picolog

import (
	"fmt"
	"os"
	"time"
)

type ConsoleLogger struct {
	fileName string
	file     *os.File
	IsDev    bool
}

func (cl *ConsoleLogger) Println(data ...interface{}) {
	if cl.IsDev {
		fmt.Print(time.Now().Format(time.Stamp) + ":")
		fmt.Println(data...)
		return
	}
	fmt.Fprint(cl.file, time.Now().Format(time.Stamp)+":")
	fmt.Fprintln(cl.file, data...)
}

func (cl *ConsoleLogger) Print(data ...interface{}) {
	fmt.Fprint(cl.file, time.Now().Format(time.Stamp)+":")
	fmt.Fprint(cl.file, data...)
}

func (cl *ConsoleLogger) Close() {
	if cl.file != nil {
		cl.file.Close()
	}
}

func (cl *ConsoleLogger) SetDev() {
	cl.IsDev = true
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
