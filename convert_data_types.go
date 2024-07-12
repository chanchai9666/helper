package helper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ToInt: ฟังก์ชันนี้พยายามแปลงค่า interface{} เป็น int โดยรองรับประเภท int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, และ string. หากแปลงไม่ได้ จะคืนค่าเป็น 0.
func ToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int8, int16, int32, int64:
		return int(v.(int64))
	case uint, uint8, uint16, uint32, uint64:
		return int(v.(uint64))
	case float32, float64:
		return int(v.(float64))
	case string:
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return 0
}

// ToFloat64: ฟังก์ชันนี้พยายามแปลงค่า interface{} เป็น float64 โดยรองรับประเภท float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, และ string. หากแปลงไม่ได้ จะคืนค่าเป็น 0.0.
func ToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	case int, int8, int16, int32, int64:
		return float64(v.(int64))
	case uint, uint8, uint16, uint32, uint64:
		return float64(v.(uint64))
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0.0
}

// ToString: ฟังก์ชันนี้พยายามแปลงค่า interface{} เป็น string โดยรองรับประเภท string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, และ bool. หากแปลงไม่ได้ จะใช้ fmt.Sprintf เพื่อคืนค่าเป็น string.
func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(v.(int64), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case float32, float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// ToBool: ฟังก์ชันนี้พยายามแปลงค่า interface{} เป็น bool โดยรองรับประเภท bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, และ float64. หากแปลงไม่ได้ จะคืนค่าเป็น false.
func ToBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	case int, int8, int16, int32, int64:
		return v.(int64) != 0
	case uint, uint8, uint16, uint32, uint64:
		return v.(uint64) != 0
	case float32, float64:
		return v.(float64) != 0.0
	}
	return false
}

// TrimSpace: ใช้ฟังก์ชัน strings.TrimSpace ในการลบช่องว่างที่อยู่ด้านหน้าและด้านหลังของ string
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// ToLower: ใช้ฟังก์ชัน strings.ToLower ในการแปลง string เป็นตัวอักษรตัวเล็ก
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper: ใช้ฟังก์ชัน strings.ToUpper ในการแปลง string เป็นตัวอักษรตัวใหญ่
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// RemoveExtraSpaces: ใช้ฟังก์ชัน strings.Fields ในการแยก string เป็น slice ของคำโดยตัดช่องว่างเกินออก จากนั้นใช้ strings.Join เพื่อรวมคำเหล่านั้นกลับมาเป็น string โดยใช้ช่องว่างหนึ่งช่องเป็นตัวแบ่ง
func RemoveExtraSpaces(s string) string {
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}

// Contains - ตรวจสอบว่า item อยู่ในสไลซ์หรือไม่
func ContainsSlice(s []interface{}, item interface{}) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// GetRandomInt - สร้างตัวเลขสุ่มระหว่าง min และ max
func GetRandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
