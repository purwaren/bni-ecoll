package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	TIME_DIFF_LIMIT = 300 // in seconds
)

func HashData(data interface{}, cid, secret string) string {
	t := int(time.Now().Unix())
	dataStr, _ := json.Marshal(data)

	return doubleEncrypt(
		reverse(strconv.Itoa(t))+"."+string(dataStr),
		cid,
		secret)
}

func ParseData(input, cid, secret string) ([]byte, error) {
	parsedString := doubleDecrypt(input, cid, secret)
	idxDot := strings.Index(parsedString, ".")

	if idxDot == -1 {
		return nil, fmt.Errorf("data is invalid")
	}

	t := parsedString[:idxDot]
	d := parsedString[idxDot+1 : len(parsedString)]

	startTime, _ := strconv.Atoi(reverse(t))
	elapsed := int(time.Now().Unix()) - startTime
	if elapsed > TIME_DIFF_LIMIT {
		return nil, fmt.Errorf("time limit exceeded")
	}
	return []byte(d), nil
}

func doubleEncrypt(input, cid, secret string) string {
	var replacer = strings.NewReplacer("+", "-", "/", "_")
	var result string

	result = encrypt(input, cid)
	result = encrypt(result, secret)

	return replacer.Replace(
		strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(result)), "="))
}

func encrypt(input, key string) string {
	var result = ""
	strls := len(input)
	strlk := len(key)
	idx := 0

	for i := 0; i < strls; i++ {
		char := input[i]
		idx = i%strlk - 1
		if idx < 0 {
			idx = strlk - 1
		}
		keyChar := key[idx]

		r1 := int(char)
		r2 := int(keyChar)

		r3 := (r1 + r2) % 128
		result += string(r3)
	}

	return result
}

func doubleDecrypt(input, cid, secret string) string {
	var replacer = strings.NewReplacer("-", "+", "_", "/")
	var str = padRight(input, int(math.Ceil(float64(len(input))/4)*4), "=")

	decoded, _ := base64.StdEncoding.DecodeString(replacer.Replace(str))

	result := decrypt(string(decoded), cid)
	result = decrypt(result, secret)

	return result
}

func decrypt(input, key string) string {
	var result string
	strls := len(input)
	strlk := len(key)
	idx := 0

	for i := 0; i < strls; i++ {
		char := input[i]
		idx = i%strlk - 1
		if idx < 0 {
			idx = strlk - 1
		}
		keyChar := key[idx]

		r1 := int(char)
		r2 := int(keyChar)

		r3 := ((r1 - r2) + 256) % 128

		result += string(r3)
	}

	return result
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func padRight(str string, length int, filler string) string {
	for {
		str += filler
		if len(str) > length {
			return str[0:length]
		}
	}
}
