package sketch

import (
	"github.com/cipepser/go-sketch/util"
	"math"
)

const (
	// k is the number of hash functions and tables.
	k = 3

	// N is the size of each table.
	N = 10
)

// Sketch represents Count-min Sketch, which is a probabilistic data structure to store counts of each element.
type Sketch [k][N]int

// NewSketch constructs a Sketch.
func NewSketch() *Sketch {
	return new(Sketch)
}

// Add stores `elem` to Sketch.
func (s *Sketch) Add(elem string) error {
	for i := 0; i < k; i++ {
		h, err := util.DoubleHashing(elem, i, N)
		if err != nil {
			return err
		}
		s[i][h] += 1
	}

	return nil
}

// Cardinality returns a count the element being stored in Sketch.
func (s *Sketch) Cardinality(elem string) (int, error) {
	min := math.MaxInt64

	for i := 0; i < k; i++ {
		h, err := util.DoubleHashing(elem, i, N)
		if err != nil {
			return 0, err
		}

		v := s[i][h]
		if v < min {
			min = v
		}
	}

	return min, nil
}
