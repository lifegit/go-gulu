/**
* @Author: TheLife
* @Date: 2020-3-27 6:59 下午
* 统计代码行数
 */
package dashboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
)

var (
	linesum int
	mutex   *sync.Mutex = new(sync.Mutex)
)

var (
	// the dir where souce file stored
	rootPath string = "/Users/yxs/Project/CloudDream/control/server/v2/dream-control/often"
	// exclude these sub dirs
	nodirs [5]string = [...]string{"/node_modules", "/.umi", "/goplayer", "/uniqush", "/code.google.com"}
	// the suffix name you care
	suffixname string = ".go"
)

func TestName(t *testing.T) {
	argsLen := len(os.Args)
	fmt.Println("argsLen:", argsLen)
	if argsLen == 2 {
		rootPath = os.Args[1]
	} else if argsLen == 3 {
		rootPath = os.Args[1]
		suffixname = os.Args[2]
	}
	// sync chan using for waiting
	done := make(chan bool)
	go codeLineSum(rootPath, done)
	<-done

	fmt.Println("total line:", linesum)
}


// compute souce file line number
func codeLineSum(root string, done chan bool) {
	var goes int              // children goroutines number
	godone := make(chan bool) // sync chan using for waiting all his children goroutines finished
	isDstDir := checkDir(root)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("root: %s, panic:%#v\n", root, pan)
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

	rootfi, err := os.Lstat(root)
	checkerr(err)

	rootdir, err := os.Open(root)
	checkerr(err)
	defer rootdir.Close()

	if rootfi.IsDir() {
		fis, err := rootdir.Readdir(0)
		checkerr(err)
		for _, fi := range fis {
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}
			goes++
			if fi.IsDir() {
				go codeLineSum(root+"/"+fi.Name(), godone)
			} else {
				go readfile(root+"/"+fi.Name(), godone)
			}
		}
	} else {
		goes = 1 // if rootfi is a file, current goroutine has only one child
		go readfile(root, godone)
	}
}

func readfile(filename string, done chan bool) {
	var line int
	isDstFile := strings.HasSuffix(filename, suffixname)
	defer func() {
		if pan := recover(); pan != nil {
			fmt.Printf("filename: %s, panic:%#v\n", filename, pan)
		}
		if isDstFile {
			addLineNum(line)
			fmt.Printf("file %s complete, line = %d\n", filename, line)
		}
		// this goroutine done, notify his parent
		done <- true
	}()
	if !isDstFile {
		return
	}

	file, err := os.Open(filename)
	checkerr(err)
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
func checkDir(dirpath string) bool {
	for _, dir := range nodirs {
		if rootPath+dir == dirpath {
			return false
		}
	}
	return true
}

func addLineNum(num int) {
	mutex.Lock()

	defer mutex.Unlock()
	linesum += num
}

// if error happened, throw a panic, and the panic will be recover in defer function
func checkerr(err error) {
	if err != nil {
		panic(err.Error())
	}
}