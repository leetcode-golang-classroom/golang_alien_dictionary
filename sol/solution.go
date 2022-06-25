package sol

import "container/heap"

type charMinHeap []byte

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (h *charMinHeap) Len() int {
	return len(*h)
}
func (h *charMinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *charMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *charMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *charMinHeap) Push(val interface{}) {
	*h = append(*h, val.(byte))
}

/**
 * @param words: a list of words
 * @return: a string which is correct order
 */
func AlienOrder(words []string) string {
	wordsLen := len(words)
	if wordsLen == 0 || words == nil {
		return ""
	}
	adjacencyMap := make(map[byte]charMinHeap)
	inDegree := make(map[byte]int)
	// init adjacencyMap
	for _, word := range words {
		wordLen := len(word)
		for idx := 0; idx < wordLen; idx++ {
			inDegree[word[idx]] = 0
			adjacencyMap[word[idx]] = []byte{}
		}
	}
	// compute direct edge
	for idx := 0; idx < wordsLen-1; idx++ {
		word1 := words[idx]
		word2 := words[idx+1]
		word1Len := len(word1)
		word2Len := len(word2)
		minLen := min(word1Len, word2Len)
		if word1Len > word2Len && word1[:minLen] == word2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				// add word2[j] to word1[j] path
				adjacencyMap[word1[j]] = append(adjacencyMap[word1[j]], word2[j])
				// increase word2[j] in_degree
				inDegree[word2[j]] += 1
				break
			}
		}
	}
	// add in degree 0 vertex
	priorityQueue := &charMinHeap{}
	heap.Init(priorityQueue)
	for key, value := range inDegree {
		if value == 0 {
			heap.Push(priorityQueue, key)
		}
	}
	result := []byte{}
	for priorityQueue.Len() != 0 {
		first := heap.Pop(priorityQueue).(byte)
		result = append(result, first)
		children := adjacencyMap[first]
		for _, child := range children {
			// choose child
			inDegree[child] -= 1
			if inDegree[child] == 0 {
				heap.Push(priorityQueue, child)
			}
		}
	}
	if len(result) != len(inDegree) {
		return ""
	}
	return string(result)
}
