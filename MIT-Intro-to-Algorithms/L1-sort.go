package GoAlgorithmsMIT

// Interface borrowed from sort package
type Interface interface {
		// Len is the number of elements in the collection.
		Len() int
		// Less reports whether the element with
		// index i should sort before the element with index j.
		Less(i, j int) bool
		// Swap swaps the elements with indexes i and j.
		Swap(i, j int)
}

/* Pseudocode from lecture
Insertion-Sort(A, n)
	for j := 2 to n
	do	key := A[j]
		i := j - 1
		while i > 0 and A[i] > key
		do	A[i+1] = A[i]
			i = i - 1
		A[i+1] = key
*/
// A, is the array (uses Interface)
// n, is the last indice to be sorted of the array
func InsertionSort (A Interface, a, b int) {
	for j := a + 1;
		j <= b;
		j++ {
		for i := j - 1;
			i >= a && A.Less(j, i);
			i-- {
			A.Swap(i + 1, i)
		}
	}
}
