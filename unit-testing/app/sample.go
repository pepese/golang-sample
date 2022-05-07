package app

import "fmt"

type SampleInterfacer interface {
	SampleFunc(in int) int
}

func Execute(s SampleInterfacer) {
	result := s.SampleFunc(1)
	fmt.Printf("result = %d\n", result)
}
