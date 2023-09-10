package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
	// 默认情况下无参数生成8位包含大小写字母和数字的随机字符串
	fmt.Println(generateRandomString(8))
	} else if len(os.Args) == 2 { //只有1个参数的情况
		input := os.Args[1] //参数值赋值为input
		if isValidSpecialChars(input) {
			// 如果参数是特殊字符串，则添加特殊字符并生成指定长度包含特殊字符的随机字符串
			if hasDuplicateSpecialChars(input) { //有重复字符则报错
				fmt.Println("Error: Special characters contain duplicates.")
				os.Exit(1)
			}
			length := 8 //length赋值8
			if len(input) > length { //参数长度大于8的话
				length = len(input) //length长度赋值为参数长度
			}
			fmt.Println(generateRandomStringWithSpecialChars(input, length)) //生成含有特殊字符的字符串
		} else {
			length, err := strconv.Atoi(input)
			if err != nil || length < 4 || length > 32 {
				fmt.Println("Error: Invalid length. Please provide a length between 4 and 32.")
				os.Exit(1)
			}
			// 如果参数是整数，则生成指定长度的随机字符串
			fmt.Println(generateRandomString(length))
		}
	} else if len(os.Args) == 3 { //只有2个参数的情况
		lengthStr := os.Args[1]
		specialChars := os.Args[2]

		length, err := strconv.Atoi(lengthStr)
		if err != nil || length < 4 || length > 32 {
			fmt.Println("Error: Invalid length. Please provide a length between 4 and 32.")
			os.Exit(1)
		}

		if !isValidSpecialChars(specialChars) {
			fmt.Println("Error: Invalid special characters. Please provide valid special characters.")
			os.Exit(1)
		}

		if hasDuplicateSpecialChars(specialChars) {
			fmt.Println("Error: Special characters contain duplicates.")
			os.Exit(1)
		}

		// 如果参数是整数且第二个参数是特殊字符串，则生成指定长度的包含特殊字符的随机字符串
		fmt.Println(generateRandomStringWithSpecialChars(specialChars, length))
	} else {
		fmt.Println("Error: Too many arguments.")
		os.Exit(1)
	}
}

func isValidSpecialChars(input string) bool {
	allowedChars := "!@#$%^&*"
	for _, char := range input {
		if !strings.ContainsRune(allowedChars, char) {
			return false
		}
	}
	return true
}

func hasDuplicateSpecialChars(input string) bool {
	charCount := make(map[rune]int)
	for _, char := range input {
		if strings.ContainsRune("!@#$%^&*", char) {
			charCount[char]++
			if charCount[char] > 1 {
				return true
			}
		}
	}
	return false
}

func generateRandomStringWithSpecialChars(specialChars string, length int) string {
	if len(specialChars) == 1 {
		if length >= 4 && length <= 8 {
			// 长度在4到8之间，生成的随机字符串必须包含一个特殊字符，剩下的位数为普通字符
			return generateRandomStringWithOneSpecialChar(specialChars, length)
		} else if length > 8 && length <= 32 {
			// 长度在9到32之间，生成的随机字符串必须包含一个特殊字符，剩下的位数为普通字符与这个特殊字符一并随机生成
			return generateRandomStringWithSpecialCharAppended(specialChars, length)
		}
	} else if len(specialChars) == 2 {
		if length >= 4 && length <= 5 {
			// 长度在4到5之间，从2个特殊字符中随机选择1个特殊字符，生成的随机字符串必须包含这1个特殊字符，剩下的位数为普通字符
			return generateRandomStringWithOneRandomSpecialChar(specialChars, length)
		} else if length >= 6 && length <= 8 {
			// 长度在6到8之间，生成的随机字符串必须包含两个特殊字符，剩下的位数为普通字符
			return generateRandomStringWithTwoSpecialChars(specialChars, length)
		} else if length > 8 && length <= 32 {
			// 长度在9到32之间，生成的随机字符串必须包含两个特殊字符，剩下的位数为普通字符与这两个特殊字符一并随机生成
			return generateRandomStringWithSpecialCharsAppended(specialChars, length)
		}
	} else if len(specialChars) >= 3 {
		if length >= 4 && length <= 5 {
			// 长度在4到5之间，从3个及以上特殊字符中随机选择1个特殊字符，生成的随机字符串必须包含这1个特殊字符，剩下的位数为普通字符
			return generateRandomStringWithOneRandomSpecialChar(specialChars, length)
		} else if length >= 6 && length <= 8 {
			// 长度在6到8之间，从3个及以上特殊字符中随机选择2个特殊字符，生成的随机字符串必须包含这2个特殊字符，剩下的位数为普通字符
			return generateRandomStringWithTwoRandomSpecialChars(specialChars, length)
		} else if length > 8 && length <= 32 {
			// 长度在9到32之间，从3个及以上特殊字符中随机选择2个特殊字符，生成的随机字符串必须包含这2个特殊字符，剩下的位数为普通字符与这两个特殊字符一并随机生成
			return generateRandomStringWithMultipleSpecialCharsAppended(specialChars, length)
		}
	}

	// 如果未满足任何条件，生成指定长度的随机字符串
	return generateRandomString(length)
}

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func generateRandomStringWithOneSpecialChar(specialChar string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length-1)

	// 生成普通字符
	for i := 0; i < length-1; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 插入一个特殊字符
	result = append(result, specialChar[0])

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithSpecialCharAppended(specialChar string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + specialChar

	// 创建一个用于存储随机字符的切片
	result := make([]byte, length)

	// 随机生成 length-1 个普通字符
	for i := 0; i < length-1; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 随机生成一个特殊字符的索引
	rand.Seed(time.Now().UnixNano())
	specialCharIndex := rand.Intn(length)

	// 在特殊字符索引处插入特殊字符
	result[specialCharIndex] = specialChar[0]

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithTwoSpecialChars(specialChars string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length-2)

	// 生成普通字符
	for i := 0; i < length-2; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 插入两个特殊字符
	result = append(result, specialChars[0])
	result = append(result, specialChars[1])

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithSpecialCharsAppended(specialChars string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + specialChars
	// 创建一个用于存储随机字符的切片，初始长度为 length
	result := make([]byte, length)

	// 随机生成 length-2 个普通字符
	for i := 0; i < length-2; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 随机生成两个特殊字符的索引
	rand.Seed(time.Now().UnixNano())
	specialCharIndex1 := rand.Intn(length - 1)
	specialCharIndex2 := rand.Intn(length - 1)

	// 在特殊字符索引处插入两个特殊字符
	result[specialCharIndex1] = specialChars[0]
	result[specialCharIndex2] = specialChars[1]

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithMultipleSpecialCharsAppended(specialChars string, length int) string {
	// 创建一个用于存储随机字符的切片，初始长度为 length
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + specialChars
	result := make([]byte, length)

	// 随机生成 length-2 个普通字符
	for i := 0; i < length-2; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 从特殊字符中随机生成长度为2的字符串
	rand.Seed(time.Now().UnixNano())
	specialCharsIndex := rand.Intn(len(specialChars) - 1)
	twoSpecialChars := specialChars[specialCharsIndex: specialCharsIndex+2]

	// 将长度为2的特殊字符插入到生成的字符串中
	result[length-2] = twoSpecialChars[0]
	result[length-1] = twoSpecialChars[1]

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithOneRandomSpecialChar(specialChars string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length-1)

	// 生成普通字符
	for i := 0; i < length-1; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 从3个及以上特殊字符中随机选择1个特殊字符，插入到结果中
	selectedSpecialChar := specialChars[rand.Intn(len(specialChars))]
	result = append(result, selectedSpecialChar)

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func generateRandomStringWithTwoRandomSpecialChars(specialChars string, length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length-2)

	// 生成普通字符
	for i := 0; i < length-2; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// 从3个及以上特殊字符中随机选择2个特殊字符，插入到结果中
	selectedSpecialChars := selectRandomSpecialChars(specialChars, 2)
	result = append(result, selectedSpecialChars[0])
	result = append(result, selectedSpecialChars[1])

	// 打乱字符顺序
	rand.Shuffle(length, func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}

func selectRandomSpecialChars(specialChars string, count int) []byte {
	if count >= len(specialChars) {
		// 如果要选择的特殊字符数量大于等于特殊字符参数的数量，直接返回全部特殊字符
		return []byte(specialChars)
	}

	perm := rand.Perm(len(specialChars))
	selectedSpecialChars := make([]byte, count)
	for i := 0; i < count; i++ {
		selectedSpecialChars[i] = specialChars[perm[i]]
	}
	return selectedSpecialChars
}

