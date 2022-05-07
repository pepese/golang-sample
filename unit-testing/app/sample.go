package app

import "fmt"

type SampleInterfacer interface {
	SampleFunc(in int) int
}

func Execute(s SampleInterfacer) bool {
	result := s.SampleFunc(1)
	fmt.Printf("result = %d\n", result)
	return result == 101
}
