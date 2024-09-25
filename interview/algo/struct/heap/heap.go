package heap

type MaxHeap struct {
	data []int
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{data: make([]int, 0, cap)}
}

func (h *MaxHeap) Size() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// parent 完全二叉树数组中, 索引的元素的父节点的索引
func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

func (h *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) Add(e int) {
	h.data = append(h.data, e)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) siftUp(k int) {
	for ; k > 0 && h.data[k] > h.data[h.parent(k)]; k = h.parent(k) {
		h.data[k], h.data[h.parent(k)] = h.data[h.parent(k)], h.data[k]
	}
}

func (h *MaxHeap) FindMax() int {
	return h.data[0]
}

func (h *MaxHeap) ExtractMax() int {
	ret := h.FindMax()

	// 交换0和最后一个元素, 并且删除最后一个
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]

	h.siftDown(0)

	return ret
}

func (h *MaxHeap) siftDown(k int) {
	for h.leftChild(k) < h.Size() { // 左孩子是较小索引都比数组size大了说明肯定越界了, 没有左右子节点了
		// j=左子, 然后如果有右子&&右子大, 则j=右子
		j := h.leftChild(k)
		if j+1 < h.Size() && h.data[j+1] > h.data[j] {
			j = h.rightChild(k)
		}

		if h.data[k] >= h.data[j] {
			break
		}

		h.data[k], h.data[j] = h.data[j], h.data[k]
		k = j
	}
}

func (h *MaxHeap) Heapify(arr []int) {
	h.data = arr
	for i := h.parent(len(h.data) - 1); i >= 0; i-- {
		h.siftDown(i)
	}
}
