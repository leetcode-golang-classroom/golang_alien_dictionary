# golang_alien_dictionary

There is a new alien language which uses the latin alphabet. However, the order among letters are unknown to you. You receive a list of **non-empty** words from the dictionary, where words are **sorted lexicographically by the rules of this new language**. Derive the order of letters in this language.

*Contact me on wechat to get **Amazon、Google** requent Interview questions . (wechat id : **jiuzhang0607**)*

## Examples

**Example 1:**

```
Input：["wrt","wrf","er","ett","rftt"]
Output："wertf"
Explanation：
from "wrt"and"wrf" ,we can get 't'<'f'
from "wrt"and"er" ,we can get 'w'<'e'
from "er"and"ett" ,we can get 'r'<'t'
from "ett"and"rftt" ,we can get 'e'<'r'
So return "wertf"

```

**Example 2:**

```
Input：["z","x"]
Output："zx"
Explanation：
from "z" and "x"，we can get 'z' < 'x'
So return "zx"

```

## 解析

題目給定一個 字串陣列 words,

其中假定這些 words 中的字串是造著某一種字典序來判斷

也就是對兩個 word[i], word[j] 如果 i < j

必須是以下的可能

1. len(word[i]) < len(word[j])
2. minLen = min(len(word[i]), len(word[j])), 則在 index 介於 0 到 minLen 第一個 word[i][index] ≠ word[j][index] → word[i][index] < word[j][index]

要求要透過給定的 words 來找出這些字元的字典序排列

假設存在字典序的話，以最小字典序來回傳

然後把這些字元順序組成一個字串回傳

假設沒有則回傳 “”

這題的關鍵一樣要找出 adjacencyList

運用到 [topological sort](https://zh.wikipedia.org/wiki/%E6%8B%93%E6%92%B2%E6%8E%92%E5%BA%8F)

首先透過 每兩個 words 之間第一個不同的字元來把紀錄

能夠找出有向的邊在 adjacencyList 也就是每個字元可以到達哪一些字元

並且紀錄這個有向邊個數 到 inDegree

這樣每次先把 inDegree = 0 的字元 照字典序先拿出來

然後從這些 inDegree = 0 的字元 透過 adjacencyList 依序去找下一個可以排的字元

然後把這些字元蒐集起來

當發現蒐集起來的長度與所有不同字元長度不同 代表有 cycle 所以回傳 “”

解法如下圖

![](https://i.imgur.com/U67do66.png)

## 程式碼
```go
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

```
## 困難點

1. 需要理解如何透過 words 找出字元關係做出 adjacencyList
2. 需要理解 inDegree 與 adjacencyList 的關係
3. 需要知道字串排序的原理

## Solve Point

- [x]  需要理解如何透過 words 找出字元關係做出 adjacencyList
- [x]  需要透過 inDegree 來標示該字元有幾個相依
- [x]  需要以 priorityQueue 來把輸出字元做排序