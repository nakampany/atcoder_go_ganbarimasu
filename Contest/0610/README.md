# 振り返り
C問題解けなかった！

aについて
簡単に書ける書き方があって、学ぶが多すぎた
```Go
fmt.Println(math.Floor((n+2)/5) * 5)
```

# Goのarrayとsliceとmapを理解してないなーと
```go
// スライス 可変長
s3 := []int{1, 2, 3, 4}

make([]T, cap)
make([]T, len, cap)
s4 := make([]int, 5, 10)

// array 固定長
a3 := [3]int{1, 2, 3}


// map
連想配列のようなもの
a := map[string]int{"A": 0, "B": 3, "C": 4, "D": 8, "E": 9, "F": 14, "G": 23}

// make map
var m1 = make(map[int]int)       // 容量を省略してmapをつくるとき(初期値0が入る)
var m2 = make(map[int]string, 5) // 容量を指定してmapをつくるとき
m := make(map[int]int, 5)
```