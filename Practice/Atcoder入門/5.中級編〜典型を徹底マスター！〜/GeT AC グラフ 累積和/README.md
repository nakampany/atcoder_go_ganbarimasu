# 累積和

適切な前処理をしておくことで、配列上の区間の総和を求めるクエリを爆速で処理できるようになる手法

###　普通にやると TLE になる

今回は、区間に関するクエリだから、

### TLE コード

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	var s string
	fmt.Scan(&s)

	for i := 0; i < q; i++ {
		var l1, l2 int
		fmt.Scan(&l1, &l2)

		word := s[l1-1 : l2]
		count := strings.Count(word, "AC")

		fmt.Println(count)
	}

}

```

https://qiita.com/drken/items/56a6b68edef8fc605821

- 競プロ典型 90 問 010 - Score Sum Queries (\* 2)
- 全国統一プログラミング 王決定戦本戦 A - Abundant
- 第 6 回 PASTD-K 項足し算
- 第 9 回 日本情報オリンピック本選 A- 旅人
- ABC 098 C - Attention
- ABC 084 D - 2017-like
- ABC 154 D - Dice in Line
- 第 10 回 日本情報オリンピック本選 A-惑星探査 (Planetary Exploration)
- AGC 023 A - Zero-Sum Ranges
- ABC 125 C - GCD on Blackboard
