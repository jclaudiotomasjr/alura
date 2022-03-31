package main

func bubbleSort(n []int) []int {
	for i, _ := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}

func main() {
	unSorted := []int{1, 99, 83, 53, 9, 14, 8, 3, 4, 20}
	//fmt.Println(bubbleSort(unSorted))
	bubbleSort(unSorted)
}
