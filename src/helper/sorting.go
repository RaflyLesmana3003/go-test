package helper

// ----------------------------------------------------------------
// Get Seconds largest number of list
// ----------------------------------------------------------------
// 1. init list of integers.
// 2. call Sorting function to sort list.
// 3. checking whether array size is only having a data then return the array.
// 4. iterating through array and checking whether the current index value is smaller than the next index value.
// 5. if the current index value is smaller than the next value, then relocate current index position to next index position.
// 6. after iterating through array, the get second position of the array.
// ----------------------------------------------------------------

func InitializeData() int {
	// 1. init list of integers.
	var n = []int{12, 5, 7, 17, 8, 0, -1, 16, 7}

	// 2. call Sorting function to sort list.
	result := Sorting(n, len(n))

	// 6. after iterating through array, the get second position of the array.
	/*
		=================
		Find Second Largest =================
		=================
		16
	*/
	return result[1]
}

func Sorting(data []int, size int) []int {
	// 3. checking whether array size is only having a data then return the array.
	if size == 1 {
		return data
	}

	// 4. iterating through array and checking whether the current index value is smaller than the next index value.
	var i = 0
	for i < size-1 {
		// 5. if the current index value (data[i]) is smaller than the next value (data[i+1]), then relocate current index position to next index position.
		/*
			if CurrentIndexValue is smaller than NextIndexValue:
			    CurrentIndexValue will be replaced NextIndexValue and NextIndexValue will replace the CurrentIndexValue
		*/
		if data[i] < data[i+1] {
			data[i], data[i+1] = data[i+1], data[i]
		}
		i++
	}

	// call recursive loop
	Sorting(data, size-1)
	return data
}
