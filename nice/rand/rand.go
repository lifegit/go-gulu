/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package rand

import (
	"math/rand"
	"time"
)

// 返回一个具有指定的from、to和size的随机整数数组
func Ints(from, to, size int) []int {
	if to-from < size {
		size = to - from
	}

	var slice []int
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(slice))
		ret = append(ret, slice[idx])
		slice = append(slice[:idx], slice[idx+1:]...)
	}

	return ret
}

const (
	SeedNum          = "0123456789"
	SeedEnglishUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SeedEnglishLower = "abcdefghijklmnopqrstuvwxyz"
	SeedAll          = SeedEnglishUpper + SeedEnglishLower + SeedEnglishLower
)

// 返回指定长度的随机字符串['a'，'z']
func String(seed string, length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	letter := []rune(seed)
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// 返回范围[min，max]内的随机整数。
func Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)

	return min + rand.Intn(max-min)
}

// 返回范围[min，max]内的随机浮点数。
func Float32(min float32, max float32) float32 {
	if max <= min {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)

	return rand.Float32()*(max-min) + min
}
