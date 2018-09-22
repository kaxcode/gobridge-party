package main

func Sum(x int, y int) int {
	return x + y
}

// Fix Multiply Function
func Multiply(a, b int) int {
	return a *  b
}

// Math Challenge #1
func subtraction(a, b int) int{
	return a - b
}


// Math Challenge #2
// Write func for Sum of multiples of 3 *OR* 5
func multiple3or5(n int) int {
	sum := 0;
  for i := 1; i < n; i++ {
    if (i % 3 == 0) || (i % 5 == 0) {
      sum += i
    }
  }
  return sum;
}
