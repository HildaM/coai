package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCode(length int) string {
	var code string
	for i := 0; i < length; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}

func GenerateChar(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func ConvertTime(t []uint8) *time.Time {
	val, err := time.Parse("2006-01-02 15:04:05", string(t))
	if err != nil {
		return nil
	}
	return &val
}

func Unmarshal[T interface{}](data []byte) (form T, err error) {
	err = json.Unmarshal(data, &form)
	return form, err
}

func UnmarshalForm[T interface{}](data string) *T {
	form, err := Unmarshal[T]([]byte(data))
	if err != nil {
		return nil
	}
	return &form
}

func Marshal[T interface{}](data T) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

func MapToStruct[T any](data interface{}) *T {
	val := Marshal(data)
	if form, err := Unmarshal[T]([]byte(val)); err == nil {
		return &form
	} else {
		return nil
	}
}

func ToInt(value string) int {
	if res, err := strconv.Atoi(value); err == nil {
		return res
	} else {
		return 0
	}
}

func ConvertSqlTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetImageMarkdown(url string) string {
	return fmt.Sprintf("![image](%s)", url)
}
