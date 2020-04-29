/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package byteconv

import (
	"bytes"
)

// 取字节集中间
func Match(tracer, startByte, endByte []byte) (b []byte) {
	start := bytes.Index(tracer, startByte)
	if start == -1 {
		return
	}
	
	end := bytes.Index(tracer[start:], endByte)
	if end == -1 {
		return
	}

	return tracer[start+len(startByte) : start+end+len(endByte)]
}


// 取字节集中间
func MatchLast(tracer, startByte, endByte []byte) (b []byte) {
	start := bytes.Index(tracer, startByte)
	if start == -1 {
		return
	}

	end := bytes.LastIndex(tracer, endByte)
	if end == -1 {
		return
	}

	return tracer[start+len(startByte) : end]
}