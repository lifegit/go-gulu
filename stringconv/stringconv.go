/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package stringconv

import (
	"strings"
)

// 取文件中间
func Match(tracer, startStr, endStr string) string {
	start := strings.Index(tracer, startStr)
	if start == -1 {
		return ""
	}
	start = start + len(startStr)

	end := strings.Index(tracer[start:], endStr)
	if end == -1 {
		return ""
	}

	return tracer[start : start+end]
}
