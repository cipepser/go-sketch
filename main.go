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
		"f", "f", "f", "f", "f", "f",
		"g", "g", "g", "g", "g", "g", "g",
		"h", "h", "h", "h", "h", "h", "h", "h",
	}
	uniques := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for _, e := range elements {
		err := s.Add(e)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(s)

	for _, e := range uniques {
		cnt, err := s.Count(e)
		if err != nil {
			panic(err)
		}
		fmt.Println(e, ":", cnt)
	}
}
