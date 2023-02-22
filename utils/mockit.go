package utils

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
)

// A global mock value generator for int, float, string, bytes

var intValue int32

func getIntValue() int32 {
	val := atomic.AddInt32(&intValue, 1)
	if val > math.MaxInt32-1000 {
		atomic.StoreInt32(&intValue, 0)
	}
	return val
}

func MockInt() int {
	return int(getIntValue())
}

func MockInt32() int32 {
	return getIntValue()
}

func MockInt64() int64 {
	return int64(getIntValue()) + math.MaxInt32
}

var i64 int64 = math.MaxInt32

func getInt64Value() int64 {
	val := atomic.AddInt64(&i64, 1)
	if val > math.MaxInt64-1000 {
		atomic.StoreInt64(&i64, math.MaxInt32)
	}
	return val
}

func MockByte() []byte {
	val := getInt64Value()
	str := []byte(fmt.Sprintf("%d", val))
	return str
	//res := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
	//base64.StdEncoding.Encode(res, str)
	//return res
}

func MockString() string {
	return string(MockByte())
}

func MockUUID() uuid.UUID {
	return uuid.New()
}

func MockTime() time.Time {
	t := time.Unix(getInt64Value(), 0)
	return t
}

func MockDate() string {
	return MockTime().Format("2006-01-02")
}

func MockDateTime() string {
	return MockTime().Format(time.RFC3339)
}
