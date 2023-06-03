5/20はB問題解けなかった
5/27はB解けた！

今回
A, C問題解けなかった
B問題だけできた！

aについて
年齢が一番小さい人を見つけて、その人基準に出力するプログラムが書けなかった
```Go
  fmt.Scan(&n)
	s := make([]string, n)
	a := make([]int, n)
 
	min := int(2e9)
	minIdx := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i], &a[i])
		if a[i] < min {
			min = a[i]
			minIdx = i
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(s[(i+minIdx)%n])
	}
```