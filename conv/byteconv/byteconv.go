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
	start = start + len(startByte)

	end := bytes.Index(tracer[start:], endByte)
	if end == -1 {
		return
	}

	return tracer[start : start+end]
}

// 取字节集最后
func MatchLast(tracer, startByte, endByte []byte) (b []byte) {
	start := bytes.Index(tracer, startByte)
	if start == -1 {
		return
	}
	start = start + len(startByte)

	end := bytes.LastIndex(tracer, endByte)
	if end == -1 {
		return
	}

	return tracer[start:end]
}
