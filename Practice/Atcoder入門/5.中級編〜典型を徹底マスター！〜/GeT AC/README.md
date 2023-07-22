# 累積和

適切な前処理をしておくことで、配列上の区間の総和を求めるクエリを爆速で処理できるようになる手法

###　普通にやるとTLEになる

今回は、区間に関するクエリだから、


### TLEコード
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