package sketch

import (
	"github.com/cipepser/go-sketch/util"
	"math"
	"strconv"
)

const (
	// k is the number of hash functions and tables.
	k = 3

	// N is the size of each table.
	N = 10
)

// TODO: write a comment
type Sketch [k][N]int

// NewSketch constructs a Sketch.
func NewSketch() *Sketch {
	return new(Sketch)
}

// TODO: write a comment
func (s *Sketch) Add(elem string) error {
	hash := util.CalcMD5Hash(elem)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	u64HashA, err := strconv.ParseUint(hashA, 16, 64)
	if err != nil {
		return err
	}
	u64HashB, err := strconv.ParseUint(hashB, 16, 64)
	if err != nil {
		return err
	}

	for i := 0; i < k; i++ {
		s[i][util.DoubleHashing(u64HashA, u64HashB, uint64(i), N)] += 1
	}

	return nil
}

// TODO: write a comment
// TODO: implement Cardinality
// TODO: Addとの共通部分を抜き出して関数化する
func (s *Sketch) Cardinality(elem string) (int, error) {
	min := math.MaxInt64
	hash := util.CalcMD5Hash(elem)
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

	for i := 0; i < k; i++ {
		v := s[i][util.DoubleHashing(u64HashA, u64HashB, uint64(i), N)]
		if v < min {
			min = v
		}
	}

	return min, nil
}
