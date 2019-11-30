package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
	"strconv"
)

// CalcMD5Hash calculates hash value of the string input by MD5.
// Remark: MD5 returns 128-bits value
func CalcMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

// DoubleHashing calculates hash value by double-hashing.
// elem: input which is hashed
// i: i-th hash value
// r: the size of hash table, used by calculating modulus
func DoubleHashing(elem string, i int, r uint64) (uint64, error) {
	hash := CalcMD5Hash(elem)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	u64HashA, err := strconv.ParseUint(hashA, 16, 64)
	if err != nil {
		return 0, err
	}
	u64HashB, err := strconv.ParseUint(hashB, 16, 64)
	if err != nil {
		return 0, err
	}

	h := doubleHashing(u64HashA, u64HashB, uint64(i), r)

	return h, nil
}

// doubleHashing calculates hash value by double-hashing.
// h = hashA + n * hashB % r
func doubleHashing(hashA, hashB uint64, n, r uint64) uint64 {
	a := big.NewInt(0).SetUint64(hashA)
	b := big.NewInt(0).SetUint64(hashB)
	mul := big.NewInt(0).SetUint64(n)
	rem := big.NewInt(0).SetUint64(r)

	h := new(big.Int).Mul(mul, b)
	h = new(big.Int).Add(a, h)
	h = new(big.Int).Rem(h, rem)

	return h.Uint64()
}
