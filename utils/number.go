package utils

import (
	"crypto/md5"
	"sync/atomic"
	"time"
)

// ID生成规则：32B secondtime + 6B sid + 6B flag + 20B autoinc, 每秒最大调用量为1048576
var autoIncSig uint64 = 0

func IsNumber(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != '-' && (str[i] < '0' || str[i] > '9') {
			return false
		}
	}
	return true
}

func GetBigID(sid int, flag int) uint64 {
	localSig := atomic.AddUint64(&autoIncSig, 1)
	bigID := uint64(time.Now().Unix())
	bigID = bigID<<20 | (localSig % (1 << 20)) //避免产生连续ID
	bigID = bigID<<6 | (uint64(sid) % (1 << 6))
	bigID = bigID<<6 | (uint64(flag) % (1 << 6))
	return bigID
}

func ParseBigID(bigID uint64) (int, int, int64) {
	flag := int(bigID % (1 << 6))
	sid := int(bigID >> 6 % (1 << 6))
	timestamp := int64(bigID >> 32)
	return sid, flag, timestamp
}

// Hash64 取低8个字节作为digest的值
func Hash64(s string) uint64 {
	sum := md5.Sum([]byte(s))
	digest := uint64(0)
	for i := 0; i < 8; i++ {
		idx := uint64(sum[i])
		digest += idx << uint64(i*8)
	}
	return digest
}

// Hash56 取低7个字节作为digest的值
func Hash56(s string) uint64 {
	sum := md5.Sum([]byte(s))
	digest := uint64(0)
	for i := 0; i < 7; i++ {
		idx := uint64(sum[i])
		digest += idx << uint64(i*8)
	}
	return digest
}

// Hash32 取低4个字节作为digest的值
func Hash32(s string) uint32 {
	sum := md5.Sum([]byte(s))
	digest := uint32(0)
	for i := 0; i < 4; i++ {
		idx := uint32(sum[i])
		digest += idx << uint64(i*8)
	}
	return digest
}
