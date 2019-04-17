package main

import (
	"fmt"
)

/*
You have a 2-dimensional rectangular crate of size X by Y, and a
bunch of boxes, each of size x by y. The dimensions are all positive
integers.

Given X, Y, x, and y, determine how many boxes can fit into a single
crate if they have to be placed so that the x-axis of the boxes is
aligned with the x-axis of the crate, and the y-axis of the boxes is
aligned with the y-axis of the crate. That is, you can't rotate the
boxes. The best you can do is to build a rectangle of boxes as large
as possible in each dimension.

For instance, if the crate is size X = 25 by Y = 18, and the boxes are
size x = 6 by y = 5, then the answer is 12. You can fit 4 boxes along
the x-axis (because 6*4 <= 25), and 3 boxes along the y-axis
(because 5*3 <= 18), so in total you can fit 4*3 = 12 boxes in a rectangle.

MAIN CHALLENGE: DONE
BONUS #1: DONE
BONUS #2: DONE
BONUS #3: TODO

*/

func main() {

	// x, y, a, b := 12345, 678910, 1112, 1314
	x, y, z, a, b, c := 1234567, 89101112, 13141516, 171819, 202122, 232425
	numFit := fit3(x, y, z, a, b, c)
	fmt.Println(numFit)

}

// TODO: solve with recursive backtracking solution
func bonus3() {
	tester := []int{2, 3, 4, 5}
	// var storage [][]int
	for i := 0; i < len(tester); i++ {
		t := tester[i]
		newTest := append(tester[:0:0], tester...)
		rest := remove(newTest, i)
		for p := 0; p < len(rest); p++ {
			fmt.Println(insertValue(rest, t, p))
		}
		//newStore := insertValue(rest)
		//fmt.Println(i, tester)
		//fmt.Printf("extracted: %d -- rest: %v\n", t, rest)
		// for p := 1; p < len(rest); p++ {
		// 	newRow := insertValue(rest, t, p)
		// 	//storage = append(storage, insertValue(rest, t, p))
		// 	fmt.Println(rest, t, p, newRow)
		// }
	}
}

func insertValue(list []int, x int, p int) []int {
	var returnArray []int
	if p == 0 { // insert into front of slice
		returnArray = []int{x}
		returnArray = append(returnArray, list[:]...)
	} else if p > 0 { // insert into position(p)
		returnArray = append(list, 0)
		copy(returnArray[p:], returnArray[p-1:])
		returnArray[p-1] = x
	}
	return returnArray
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func fit4(crateD []int, boxD []int) int {
	ans := 0
	if len(crateD) == len(boxD) {

	}
	return ans
}

func findPermutations(boxD []int) [][]int {
	var outPermutations [][]int
	if len(boxD) > 1 {
		//var newPerm []int
		for i := 0; i < len(boxD); i++ {
			x := boxD[i]            // get ith number
			rest := remove(boxD, i) // send all rest to another list for recursive call
			fmt.Println(x, rest)
		}
	}

	return outPermutations
}

// Optional Bonus #3
func fit3(x, y, z, a, b, c int) int {
	flip := func(x, y, z, a, b, c int) int {
		if x/a > 0 && y/b > 0 && z/c > 0 {
			return (x / a) * (y / b) * (z / c)
		}
		return 0
	}
	ans1 := flip(x, y, z, a, b, c)
	ans2 := flip(x, y, z, a, c, b)
	ans3 := flip(x, y, z, b, a, c)
	ans4 := flip(x, y, z, c, a, b)
	ans5 := flip(x, y, z, c, b, a)
	ans6 := flip(x, y, z, b, c, a)

	highest := 0
	for _, p := range []int{ans1, ans2, ans3, ans4, ans5, ans6} {
		if p > highest {
			highest = p
		}
	}

	return highest
}

// Optional Bonus #1
func fit2(x, y, a, b int) int {
	try1 := fit1(x, y, a, b)
	try2 := fit1(x, y, b, a) // 90deg box turn
	if try1 > try2 {
		return try1
	} else if try2 > try1 {
		return try2
	}
	return try1
}

// Main challenge
func fit1(x, y, a, b int) int {
	if x/a > 0 && y/b > 0 {
		return (x / a) * (y / b)
	}
	return 0
}
