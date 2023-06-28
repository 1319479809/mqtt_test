package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// MD5 计算MD5
func MD5(text string) string {
	has := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", has)
}

// RandomStr  随机数生成
// size 长度
// Random   0: 纯数字; 1: 小写字母 ;2: 大写字母 ;3: 数字、大小写字母
func RandomStr(size int, Random int) string {
	iRandom, Randoms, result := Random, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	iAll := Random > 2 || Random < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if iAll { // random ikind
			iRandom = rand.Intn(3)
		}
		scope, base := Randoms[iRandom][0], Randoms[iRandom][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// RandomStr  随机数生成
// size 长度
// Random   0: 纯数字; 1: 小写字母 ;2: 大写字母 ;3: 数字、大小写字母
// stype 1 纯数字 2 数字、小写字母 3 数字、大小写字母
func RandomStr2(size int, Random int, stype int) string {
	iRandom, Randoms, result := Random, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	iAll := Random > 2 || Random < 0
	iAll = stype > 3 || stype < 1
	if iAll {
		stype = 3
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		iRandom = rand.Intn(stype) // random ikind
		scope, base := Randoms[iRandom][0], Randoms[iRandom][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
