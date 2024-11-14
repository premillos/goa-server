package utils

import (
	"errors"
	"math/rand"
)

// GenerateRandomNumbers 是一个生成随机数的通用函数
func GenerateRandomInteger(count int) ([]any, error) {
	// 校验输入
	if count <= 0 {
		return nil, errors.New("count must be greater than 0")
	}

	// 存储随机数的切片
	results := make([]any, count)

	for i := 0; i < count; i++ {
		// 生成整数类型的随机数
		results[i] = int(rand.Intn(9))
	}

	return results, nil
}
