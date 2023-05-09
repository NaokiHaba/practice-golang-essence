package ch03

import "fmt"

func PrintSlice() {
	var s []int
	fmt.Println(s) // []

	ss := make([]int, 3)
	ss[0] = 1
	ss[1] = 2
	ss[2] = 3
	fmt.Println(ss) // [1 2 3]

	// 多次元配列
	arr1 := [2][2]int{{1, 2}, {3, 4}}
	arr2 := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println(arr1) // [[1 2] [3 4]]
	fmt.Println(arr2) // [[0 1 2] [3 4 5]]

	// 要素を追加
	s = append(s, 1)
	fmt.Println(s) // [1]

	// スライスの部分参照
	fmt.Println(ss[1:3]) // [2 3]

	// スライスssから、インデックスnの要素を削除する
	n := 1
	ss = ss[:n+copy(ss[n:], ss[n+1:])]
	fmt.Println(ss) // [1 3]

	// 固定長の配列からスライスを作成
	var b [4]byte
	b2 := b[:]
	fmt.Println(b2) // [0 0 0 0]
}
