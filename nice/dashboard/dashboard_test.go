/**
* @Author: TheLife
* @Date: 2021/6/17 下午5:15
 */
package dashboard

import (
	"fmt"
	"testing"
)

func TestStatisticsLine(t *testing.T) {
	s := StatisticsLine{
		RootPath:    "/Users/yxs/Project/CloudDream/control/v3/server/services",
		ExcludeDirs: []string{"/node_modules", "/.umi", "/goplayer", "/uniqush", "/code.google.com"},
		SuffixName:  ".go",
	}

	line := s.Run()
	fmt.Println(line)
}
