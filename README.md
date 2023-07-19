# atcoder_go_ganbarimasu
AtcoderをGolangで書いていくよ

# 標準入力
```go
	var n int
	fmt.Scan(&n)
```

```go
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
```
```go
	s := bufio.NewScanner(os.Stdin)
	a := make([]kv, n)
	for i := 0; i < n; i++ {
		s.Scan()
		line := s.Text()
		fmt.Sscanf(line, "%d %d", &a[i].Key, &a[i].Value)
	}
```