package adjust

// Int returns adjust if value is 0
func Int(value, adjust int) int {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxInt return max int
func MaxInt(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinInt return min value
func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

// Int8 returns adjust if value is 0
func Int8(value, adjust int8) int8 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxInt8 return max int
func MaxInt8(v1, v2 int8) int8 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinInt8 return min value
func MinInt8(v1, v2 int8) int8 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// Int16 returns adjust if value is 0
func Int16(value, adjust int16) int16 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxInt16 return max int
func MaxInt16(v1, v2 int16) int16 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinInt16 return min value
func MinInt16(v1, v2 int16) int16 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// Int32 returns adjust if value is 0
func Int32(value, adjust int32) int32 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxInt32 return max int
func MaxInt32(v1, v2 int32) int32 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinInt32 return min value
func MinInt32(v1, v2 int32) int32 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// Int64 returns adjust if value is 0
func Int64(value, adjust int64) int64 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxInt64 return max int
func MaxInt64(v1, v2 int64) int64 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinInt64 return min value
func MinInt64(v1, v2 int64) int64 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// UInt returns adjust if value is 0
func UInt(value, adjust uint) uint {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxUInt return max uint
func MaxUInt(v1, v2 uint) uint {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinUInt return min value
func MinUInt(v1, v2 uint) uint {
	if v1 < v2 {
		return v1
	}

	return v2
}

// UInt8 returns adjust if value is 0
func UInt8(value, adjust uint8) uint8 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxUInt8 return max uint
func MaxUInt8(v1, v2 uint8) uint8 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinUInt8 return min value
func MinUInt8(v1, v2 uint8) uint8 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// UInt16 returns adjust if value is 0
func UInt16(value, adjust uint16) uint16 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxUInt16 return max uint
func MaxUInt16(v1, v2 uint16) uint16 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinUInt16 return min value
func MinUInt16(v1, v2 uint16) uint16 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// UInt32 returns adjust if value is 0
func UInt32(value, adjust uint32) uint32 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxUInt32 return max uint
func MaxUInt32(v1, v2 uint32) uint32 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinUInt32 return min value
func MinUInt32(v1, v2 uint32) uint32 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// UInt64 returns adjust if value is 0
func UInt64(value, adjust uint64) uint64 {
	if value == 0 {
		return adjust
	}

	return value
}

// MaxUInt64 return max uint
func MaxUInt64(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MinUInt64 return min value
func MinUInt64(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v1
	}

	return v2
}

// String returns adjust if value is nil
func String(value, adjust string) string {
	if value == "" {
		return adjust
	}

	return value
}
