package ch03

import "fmt"

func PrintMap() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m) // map[a:1 b:2]

	// cap を指定してマップを作成
	mCap := make(map[string]int, 10)
	mCap["a"] = 1
	fmt.Println(mCap) // map[a:1]

	// リテラルでマップを作成
	rMap := map[string]int{"a": 1, "b": 2}
	fmt.Println(rMap) // map[a:1 b:2]

	// 要素を削除
	delete(rMap, "a")
	fmt.Println(rMap) // map[b:2]
}
