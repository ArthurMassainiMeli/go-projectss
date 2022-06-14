package main

import f "even-odd-project/features"

func main() {
	numbers := f.CreateSliceFromOneTo(15)

	numbers.PrintOddOrEven()
}
