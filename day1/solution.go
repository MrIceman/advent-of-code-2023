package day1

import (
	_ "embed"
	"log"
)

func CalculateSum(inputs []string) int {
	sum := 0
	for _, s := range inputs {
		sum += extractNumbers(s)
	}
	log.Printf("total sum of all the calibration values: %d", sum)
	return sum
}
