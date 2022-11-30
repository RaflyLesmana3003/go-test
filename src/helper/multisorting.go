package helper

import "fmt"

// ----------------------------------------------------------------
// Mutliple Sort on list of data points
// ----------------------------------------------------------------
// 1. init map class which is contain person data (Name, Grade, Point).
// 2. call Sorting function to sort list.
// 3. checking whether array size is only having a data then return the array.
// 4. iterating through array and checking whether the current index value is smaller than the next index value.
// 5. check if the current index Name is smaller than the next index Name, then relocate current index position to next index position.
// 6. check if the current index Point is larger than the next index Point, then relocate current index position to next index position.
// 7. check if the current index Grade is smaller than the next index Grade, then relocate current index position to next index position.
// 8. reserve array position to match the expected output.
// 9. print the ordered result Names based on Grade order asc then Point order desc.
// ----------------------------------------------------------------

type Person struct {
	Name  string
	Point int32
	Grade string
}

func InitializeDataMulti() {
	//1. init map class which is contain person data (Name, Grade, Point).
	person := make(map[int]Person)
	person[0] = Person{"evan", 50000, "D"}
	person[1] = Person{"jefry", 70000, "C"}
	person[2] = Person{"rizky", 30000, "D"}
	person[3] = Person{"hanson", 10000, "B"}
	person[4] = Person{"candra", 30000, "A"}
	person[5] = Person{"goklas", 20000, "A"}
	person[6] = Person{"hendra", 20000, "B"}
	person[7] = Person{"surya", 30000, "A"}

	// 2. call Sorting function to sort list.
	calculate := SortingFilter(person, len(person))

	// 8. reserve array position to match the expected output.
	for i, j := 0, len(calculate)-1; i < j; i, j = i+1, j-1 {
		calculate[i], calculate[j] = calculate[j], calculate[i]
	}

	// 9. print the ordered result Names based on Grade order asc then Point order desc.

	/*
		=================
		Result multiple sorting =================
		=================
		candra
		surya
		goklas
		hendra
		hanson
		jefry
		evan
		rizky
	*/
	for i := 0; i < len(calculate); i++ {
		fmt.Println(calculate[i].Name)
	}

}

func SortingFilter(arr map[int]Person, size int) map[int]Person {
	// 3. checking whether array size is only having a data then return the array.
	if size == 1 {
		return arr
	}

	// 4. iterating through array and checking whether the current index value is smaller than the next index
	var i = 0
	for i < size-1 {
		// 5. check if the current index Name is smaller than the next index Name, then relocate current index position to next index position.
		if arr[i].Name[0] < arr[i+1].Name[0] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		// 6. check if the current index Point is larger than the next index Point, then relocate current index position to next index position.
		if arr[i].Point > arr[i+1].Point {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		// 7. check if the current index Grade is smaller than the next index Grade, then relocate current index position to next index position.
		if arr[i].Grade < arr[i+1].Grade {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		i++
	}

	// recursive loop
	SortingFilter(arr, size-1)
	return arr
}
