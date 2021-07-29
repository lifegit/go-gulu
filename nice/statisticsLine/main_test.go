/**
* @Author: TheLife
* @Date: 2021/6/17 下午5:15
 */
package main_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/statisticsLine/app"
	"testing"
)
//go:generate ./main -p ./ -e /node_modules,/.umi -s .go
func TestStatisticsLine(t *testing.T) {
	line := app.StatisticsLine{
		RootPath:    "./",
		ExcludeDirs: []string{"/node_modules", "/.umi"},
		SuffixName:  ".js",
	}

	lineCount := line.Run()
	fmt.Println(lineCount)
}
