package picolog

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

type LogRotator struct {
	filename string
	file     *os.File
	count    int64
	lock     *sync.Mutex
	prefix   string
	logPath  string
}

func (lr *LogRotator) Log(line []byte) {
	lr.lock.Lock()
	defer lr.lock.Unlock()
	lr.count += int64(len(line))
	if lr.count > 50000000 {
		lr.rotate()
	}
	fmt.Fprintln(lr.file, string(line))
}

func (lr *LogRotator) LogString(logItem string) {
	lr.Log([]byte(logItem))
}

func (lr *LogRotator) fileName() {
	t := time.Now()
	lr.filename = lr.prefix + "_log_" + t.Format("15_04_05") + ".log"
}

func (lr *LogRotator) rotate() {
	fmt.Println("Rotating")
	if lr.file != nil {
		lr.file.Close()
	}

	lr.fileName()
	lr.open()
}

func (lr *LogRotator) open() {
	files, err := ioutil.ReadDir(lr.logPath)

	if err == nil && len(files) > 0 {
		zipFile, err := os.Create(lr.logPath + "backup" + time.Now().Format("15_04_05") + ".gzip")
		if err != nil {
			fmt.Println(err)
		}
		gzipped := gzip.NewWriter(zipFile)
		for _, f := range files {
			if strings.Index(f.Name(), ".gzip") > 3 || strings.Index(f.Name(), lr.prefix+"_") != 0 {
				continue
			}
			s, _ := os.Open(lr.logPath + f.Name())
			io.Copy(gzipped, s)
			s.Close()
			_ = os.Remove(lr.logPath + f.Name())
		}
		gzipped.Flush()
		gzipped.Close()
		zipFile.Close()
	}

	if lr.filename == "" {
		lr.fileName()
	}
	_, err = os.Stat(lr.logPath + lr.filename)
	if err != nil {
		lr.file, err = os.Create(lr.logPath + lr.filename)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		lr.file, err = os.Open(lr.logPath + lr.filename)
		if err != nil {
			fmt.Println(err)
		}
	}
	stat, err := os.Stat(lr.logPath + lr.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	lr.count = stat.Size()
}

func New(prefix string) *LogRotator {
	lr := &LogRotator{
		prefix:  prefix,
		logPath: getLogPath(prefix),
		lock:    &sync.Mutex{},
	}
	createIfNotExist(lr.logPath)
	lr.open()
	return lr
}

func (lr *LogRotator) Close() {
	if lr.file != nil {
		lr.file.Close()
	}
}

func createIfNotExist(path string) {
	if _, err := os.Stat(path); err != nil {
		os.MkdirAll(path, 777)
	}
}
