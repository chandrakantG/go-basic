package main

import "fmt"

//array is a fixed-size sequence of elements of a same type

func main() {
	// # syntax or declare an array
	// var arrayName [size]Type
	var nums [5]int
	fmt.Println(nums)

	// # initialize an array at time of declaration

	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// # shorthand
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)

	// # zero value
	// If an array is declared but not initialized, its elements take the zero value of the array's type
	var numbers [5]int
	fmt.Println(numbers)

	// # length
	// The len() function will always return the size you declared, irrespective of the number of elements present.
	students := [5]int{1, 2, 3, 4, 5}
	fmt.Println(students)
	fmt.Println("len[no of elements == size ]:", len(students))
	newStudents := [5]int{1, 2, 3, 4}
	fmt.Println("len[no of elements != size ]:", len(newStudents))

	// # accessing array elements using index
	// syntax:
	// element = arrayName[index]
	firstStudent := students[0] // accessing the first element
	fmt.Println(firstStudent)

	// # bounds(size limit) checking
	// trying to access an index outside the valid range will result in a runtime panic.
	// unknown := students[5]

	// # iterate over an array
	// with traditional for loop
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i]) // Prints each number
	}

	// with range
	for k, v := range students {
		fmt.Println(k, v)
	}

	// nested array
	var matrix = [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// #  modify/update/manipulate element
	// direct assignment
	students[0] = 6

	// using for loop
	for i, v := range students {
		if v == 5 {
			students[i] = 6
		}
	}

	// # Partial Updates With Slices
	// Slices, derived from arrays, can also facilitate modifications, especially when targeting a subsection of the array.
	fmt.Println("before fruits update:", fruits)
	slice := fruits[0:2] // Slice containing the first two elements
	fmt.Println("before slice change:", slice)
	slice[1] = "kiwi" // Modifies the second element of the slice and the original array
	fmt.Println("after slice change:", slice)
	fmt.Println("after fruits update:", fruits)
	// Remember, slices share the same underlying data with the original array. Thus, modifying a slice also modifies the corresponding elements in the array.

	// # Runtime Panic On Out-Of-Bounds Access
	// Go performs automatic bounds checking for arrays.
	// If you attempt to access or modify an element using an index outside the valid range, the program will trigger a runtime panic
	// var colors = [3]string{"Red", "Green", "Blue"}
	// fmt.Println(colors[4]) // This will panic

	// Using The Len Function For Safety
	// To avoid out-of-bounds issues, always ensure that indices are within the valid range. One way to achieve this is by using the len function, which returns the size of the array.

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i]) // Safe access of elements
	}

	// # Common Array Operations
	// 1.Determining The Length:
	length := len(fruits)
	fmt.Println("length:", length)

	// 2.Compair Array: you can directly compare arrays of the same type and size using relational operators
	var a = [3]int{1, 2, 3}
	var b = [3]int{1, 2, 3}
	isEqual := a == b // isEqual will be true
	fmt.Println("isEqual:", isEqual)
	// The arrays a and b are identical, so the comparison returns true.

	// 3. Copying Array
	// Copying is achieved through assignment. But remember, this creates a shallow copy.
	var original = [3]int{10, 20, 30}
	var copyArr = original // Creates a copy of the original array
	copyArr[0] = 50        // Modifying copy doesn't affect original
	fmt.Println("copyArr:", copyArr)

	// 4. Finding Elements
	// finding an elements using iterating over loop
	var target = "Banana"
	found := false
	for _, fruit := range fruits {
		if fruit == target {
			found = true
			break
		}
	}
	fmt.Println("found:", found)

	// 5. Filling an array
	var values [5]int
	for i := range values {
		values[i] = -1
	}
	fmt.Println("values:", values)

}
