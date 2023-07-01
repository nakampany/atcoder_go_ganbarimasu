# 振り返り


今までは、decimal型、可変ビットだったから気にしなかったけど
Golangは、固定長だから、ビットを考えないといけない


マイナスに吹っ飛んだ
uint型を採用！！！

```Go
  var numbers []int

	var ans float64
	ans = 0

	for i, v := range numbers {
		ans += float64(v) * math.Pow(2, float64(i))
	}

	fmt.Println(int(ans))
```

```Go
  var numbers []int

	var ans uint64
	ans = 0

	for i, v := range numbers {
		ans += uint64(v) * uint64(math.Pow(2, float64(i)))
	}

	fmt.Println(uint64(ans))
```
