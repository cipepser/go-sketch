package main

import (
	"fmt"
	"github.com/cipepser/go-sketch/sketch"
)

func main() {
	s := sketch.NewSketch()

	elements := []string{
		"a",
		"b", "b",
		"c", "c", "c",
		"d", "d", "d", "d",
		"e", "e", "e", "e", "e",
	}
	uniques := []string{"a", "b", "c", "d", "e"}

	for _, e := range elements {
		err := s.Add(e)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(s)

	for _, e := range uniques {
		cnt, err := s.Cardinality(e)
		if err != nil {
			panic(err)
		}
		fmt.Println(e, ":", cnt)
	}
}
