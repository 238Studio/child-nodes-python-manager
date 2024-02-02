package py

import "container/heap"

//实现对指令的小根堆排序

// RequestHeap 请求堆 优先权值0～255的小根堆
type RequestHeap []*PythonScriptRequest

func (pq *RequestHeap) Len() int { return len(*pq) }

func (pq *RequestHeap) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *RequestHeap) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	//交换索引
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *RequestHeap) Push(x interface{}) {
	// 获取元素下标
	n := len(*pq)
	item := x.(*PythonScriptRequest)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *RequestHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	x.index = -1 //安全
	*pq = old[0 : n-1]
	return x
}

// Remove 移除堆中的请求对象
func (pq *RequestHeap) Remove(i int) {
	heap.Remove(pq, i)
}
