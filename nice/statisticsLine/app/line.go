/**
* @Author: TheLife
* @Date: 2020-3-27 6:59 下午
* 统计代码行数
 */
package app

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
)

type StatisticsLine struct {
	// the dir where souce file stored
	RootPath string
	// exclude these sub dirs
	ExcludeDirs []string
	// the suffix name you care
	SuffixName string

	lineSum int
	mutex   sync.Mutex
}

func (s *StatisticsLine) Run() int {
	done := make(chan bool)
	go s.codeLineSum(s.RootPath, done)
	<-done

	return s.lineSum
}

// compute souce file line number
func (s *StatisticsLine) codeLineSum(root string, done chan bool) {
	var goes int              // children goroutines number
	godone := make(chan bool) // sync chan using for waiting all his children goroutines finished
	isDstDir := s.checkDir(root)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("path: %s, panic:%#v\n", root, pan)
		}

		// waiting for his children done
		for i := 0; i < goes; i++ {
			<-godone
		}

		// this goroutine done, notify his parent
		done <- true
	}()
	if !isDstDir {
		return
	}

	rootFi, err := os.Lstat(root)
	s.checkErr(err)

	rootDir, err := os.Open(root)
	s.checkErr(err)
	defer rootDir.Close()

	if rootFi.IsDir() {
		fis, err := rootDir.Readdir(0)
		s.checkErr(err)
		for _, fi := range fis {
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			goes++
			if fi.IsDir() {
				go s.codeLineSum(path.Join(root, fi.Name()), godone)
			} else {
				go s.readFile(path.Join(root, fi.Name()), godone)
			}
		}
	} else {
		goes = 1 // if rootFi is a file, current goroutine has only one child
		go s.readFile(root, godone)
	}
}

func (s *StatisticsLine) readFile(filename string, done chan bool) {
	var line int
	isDstFile := strings.HasSuffix(filename, s.SuffixName)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("filename: %s, panic:%#v\n", filename, pan)
		}
		if isDstFile {
			s.addLineNum(line)
			fmt.Printf("file %s complete, line = %d\n", filename, line)
		}
		// this goroutine done, notify his parent
		done <- true
	}()
	if !isDstFile {
		return
	}

	file, err := os.Open(filename)
	s.checkErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			line++
		}
	}
}

// check whether this dir is the dest dir
func (s *StatisticsLine) checkDir(dirPath string) bool {
	for _, dir := range s.ExcludeDirs {
		if path.Join(s.RootPath, dir) == dirPath {
			return false
		}
	}
	return true
}

func (s *StatisticsLine) addLineNum(num int) {
	s.mutex.Lock()

	defer s.mutex.Unlock()
	s.lineSum += num
}

// if error happened, throw a panic, and the panic will be recover in defer function
func (s *StatisticsLine) checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
