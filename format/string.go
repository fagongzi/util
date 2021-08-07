package format

import (
	"log"
	"runtime"
	"strconv"

	"github.com/fagongzi/util/hack"
)

var (
	boolBytes = []rune{'t', 'r', 'u', 'e', 'T', 'R', 'U', 'E'}
	boolTrue  = []byte{'t', 'r', 'u', 'e'}
	boolFalse = []byte{'f', 'a', 'l', 's', 'e'}
)

// Uint32ToStringBytes uint32 to str
func Uint32ToStringBytes(v uint32) []byte {
	return strconv.AppendUint(nil, uint64(v), 10)
}

// Uint32ToString uint32 to str
func Uint32ToString(v uint32) string {
	return hack.SliceToString(Uint32ToStringBytes(v))
}

// ParseStringUint32 str -> uint32
func ParseStringUint32(data string) (uint32, error) {
	ret, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(ret), nil
}

// ParseStringUint32Bytes str -> uint64
func ParseStringUint32Bytes(data []byte) (uint32, error) {
	return ParseStringUint32(hack.SliceToString(data))
}

// MustParseStringUint32 str -> uint32
func MustParseStringUint32(data string) uint32 {
	value, err := ParseStringUint32(data)
	if err != nil {
		logPanic("uint32", data, err)
	}

	return value
}

// Uint64ToStringBytes uint64 to str
func Uint64ToStringBytes(v uint64) []byte {
	return strconv.AppendUint(nil, v, 10)
}

// Uint64ToString uint64 to str
func Uint64ToString(v uint64) string {
	return hack.SliceToString(Uint64ToStringBytes(v))
}

// ParseStringUint64 str -> uint64
func ParseStringUint64(data string) (uint64, error) {
	return strconv.ParseUint(data, 10, 64)
}

// ParseStringUint64Bytes str -> uint64
func ParseStringUint64Bytes(data []byte) (uint64, error) {
	return ParseStringUint64(hack.SliceToString(data))
}

// MustParseStringUint64 str -> uint64
func MustParseStringUint64(data string) uint64 {
	value, err := ParseStringUint64(data)
	if err != nil {
		logPanic("uint64", data, err)
	}

	return value
}

// ParseStringInt str -> int
func ParseStringInt(data string) (int, error) {
	v, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

// ParseStringIntBytes str -> int
func ParseStringIntBytes(data []byte) (int, error) {
	return ParseStringInt(hack.SliceToString(data))
}

// MustParseStringInt str -> int
func MustParseStringInt(data string) int {
	value, err := ParseStringInt(data)
	if err != nil {
		logPanic("int", data, err)
	}

	return value
}

// Int64ToStringBytes int64 to str
func Int64ToStringBytes(v int64) []byte {
	return strconv.AppendInt(nil, v, 10)
}

// Int64ToString int64 to str
func Int64ToString(v int64) string {
	return hack.SliceToString(Int64ToStringBytes(v))
}

// ParseStringInt64 str -> int64
func ParseStringInt64(data string) (int64, error) {
	return strconv.ParseInt(data, 10, 64)
}

// ParseStringInt64Bytes str -> int
func ParseStringInt64Bytes(data []byte) (int64, error) {
	return ParseStringInt64(hack.SliceToString(data))
}

// MustParseStringInt64 str -> int64
func MustParseStringInt64(data string) int64 {
	value, err := ParseStringInt64(data)
	if err != nil {
		logPanic("int64", data, err)
	}

	return value
}

// Float64ToStringBytes float64 to str
func Float64ToStringBytes(v float64) []byte {
	return strconv.AppendFloat(nil, v, 'f', -1, 64)
}

// Float64ToString float64 to str
func Float64ToString(v float64) string {
	return hack.SliceToString(Float64ToStringBytes(v))
}

// ParseStringFloat64 str -> float64
func ParseStringFloat64(data string) (float64, error) {
	return strconv.ParseFloat(data, 64)
}

// ParseStringFloat64Bytes str -> float64
func ParseStringFloat64Bytes(data []byte) (float64, error) {
	return ParseStringFloat64(hack.SliceToString(data))
}

// MustParseStringFloat64 str -> float64
func MustParseStringFloat64(data string) float64 {
	value, err := ParseStringFloat64(data)
	if err != nil {
		logPanic("float64", data, err)
	}

	return value
}

// ParseStringIntSlice parse []string -> []int
func ParseStringIntSlice(data []string) ([]int, error) {
	var target []int

	for _, str := range data {
		id, err := ParseStringInt(str)
		if err != nil {
			return nil, err
		}

		target = append(target, id)
	}

	return target, nil
}

// ParseStringInt64Slice parse []string -> []int64
func ParseStringInt64Slice(data []string) ([]int64, error) {
	var target []int64

	for _, str := range data {
		id, err := ParseStringInt64(str)
		if err != nil {
			return nil, err
		}

		target = append(target, id)
	}

	return target, nil
}

// ParseStringUint64Slice parse []string -> []uint64
func ParseStringUint64Slice(data []string) ([]uint64, error) {
	var target []uint64

	for _, str := range data {
		id, err := ParseStringUint64(str)
		if err != nil {
			return nil, err
		}

		target = append(target, id)
	}

	return target, nil
}

// BoolToString bool ->  str
func BoolToString(v bool) string {
	if v {
		return "true"
	}

	return "false"
}

// BoolToStringBytes bool ->  str
func BoolToStringBytes(v bool) []byte {
	if v {
		return boolTrue
	}

	return boolFalse
}

// ParseStringBool parse string -> bool
func ParseStringBool(data string) (bool, error) {
	if len(data) != 4 {
		return false, nil
	}

	for idx, c := range data {
		if c != boolBytes[idx] && c != boolBytes[idx+4] {
			return false, nil
		}
	}

	return true, nil
}

func logPanic(parseType string, data string, err error) {
	buf := make([]byte, 4096)
	runtime.Stack(buf, true)
	log.Fatalf("parse to %s failed, data=<%s> errors:\n %+v \n %s",
		parseType,
		data,
		err,
		buf)
}
