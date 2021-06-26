package common

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math"
	"reflect"
)

const (
	TimeFormat  = "2006-01-02 15:04:05"
	AdminTag    = "admin"
	AdminRoleId = 1
)

// MD5 MD5哈希值
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5哈希值
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 SHA1哈希值
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1哈希值
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

//查找数组是否包含某元素
func InArray(find interface{}, source interface{}) bool {
	targetValue := reflect.ValueOf(source)
	switch reflect.TypeOf(source).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == find {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(find)).IsValid() {
			return true
		}
	}
	return false
}

//反射结构体所有字段名
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		fmt.Println(t.Field(i).Type)
		result = append(result, t.Field(i).Name)
	}
	return result
}

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n // TODO +0.5 是为了四舍五入，如果不希望这样去掉这个
}
