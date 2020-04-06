package array

func Sum(numbers []int) int  {
	sum := 0
	// for i:= 0; i<5; i++{
	// 	sum += numbers[i]
	// }
	for _,num := range numbers{
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int,lengthOfNumbers)

	for _,numbers := range numbersToSum{
		sums = append(sums,Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int){
	for _,numbers := range numbersToSum{
		// lengthOfNums := len(numbers)
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		}else{
			sums = append(sums,0)
		}
	}
	return
}