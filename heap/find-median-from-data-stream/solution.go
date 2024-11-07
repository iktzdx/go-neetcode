package findmedianfromdatastream

import "container/heap"

// threshold defines the allowed difference between the sizes of the left and right halves of the data stream.
//
// The data stream is split approximately in the middle. If the total number of elements is odd,
// one half will contain exactly one more element than the other.
const threshold = 1

// The data stream sorted in ascending order.
type MedianFinder struct {
	left  *maxHeap // Max heap containing the left half of the data stream.
	right *minHeap // Min heap containing the right half of the data stream.
}

// Constructor initializes and returns a new instance of MedianFinder.
func Constructor() MedianFinder {
	return MedianFinder{
		left:  new(maxHeap),
		right: new(minHeap),
	}
}

// AddNum adds a provided integer to the MedianFinder.
//
// The number is inserted into the appropriate heap to maintain the order of the halves.
// If the number is less than or equal to the maximum of the left half, it goes to the left half;
// otherwise, it goes to the right half.
//
// Time complexity: O(log n).
func (f *MedianFinder) AddNum(num int) {
	// Determine which heap to add the number to based on its value.
	if f.left.Len() == 0 || num < f.left.Top() {
		heap.Push(f.left, num) // Add to the left half (max heap).
	} else {
		heap.Push(f.right, num) // Add to the right half (min heap).
	}

	// Balance the heaps if the size difference exceeds the threshold.
	if f.left.Len() > f.right.Len()+threshold {
		// Move the maximum element from the left half to the right half.
		heap.Push(f.right, heap.Pop(f.left))
	} else if f.right.Len() > f.left.Len() {
		// Move the minimum element from the right half to the left half.
		heap.Push(f.left, heap.Pop(f.right))
	}
}

// FindMedian returns the median of all elements added so far.
//
// If both halves have equal sizes, it calculates the median as the average of the two top elements.
// Otherwise, it returns the top element of the left half as the median.
//
// Time complexity: O(1).
func (f *MedianFinder) FindMedian() float64 {
	// Check if both halves have equal sizes.
	if f.left.Len() == f.right.Len() {
		// Calculate the median for an even number of elements.
		return float64((f.left.Top() + f.right.Top())) / 2.0
	}

	// Return the top element of the left half for an odd number of elements.
	return float64(f.left.Top())
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Top() int {
	return m[0]
}

func (m *maxHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return v
}

func (m *maxHeap) Push(v interface{}) {
	num, _ := v.(int)

	*m = append(*m, num)
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minHeap) Top() int {
	return m[0]
}

func (m *minHeap) Pop() interface{} {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return v
}

func (m *minHeap) Push(v interface{}) {
	num, _ := v.(int)

	*m = append(*m, num)
}
