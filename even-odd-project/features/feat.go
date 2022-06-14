package features

import "fmt"

type NumberArray []int

func CreateSliceFromOneTo(size int) NumberArray {
	var numbers NumberArray
	for i := 1; i <= size; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func (na NumberArray) PrintOddOrEven() {
	for _, number := range na {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
