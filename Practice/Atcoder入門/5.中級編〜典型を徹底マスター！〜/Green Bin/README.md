# バケット、連想配列

### strings.Split(s, ""): "hello" を ["h", "e", "l", "l", "o"]

### sort.Strings(slice): ["h", "e", "l", "l", "o"] -> ["e", "h", "l", "l", "o"] にソート

### strings.Join(slice, ""): ソートされた文字列のスライスを一つの文字列に結合("ehllo")

```
func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}
```

### N 個の文字列をグループ化して、その個数を数える

### [abbemnoptu, acinnorstt acinnorstt(同じ)　]

```
total += count * (count - 1) / 2
```

```
n(n-1)/2
nC2
```
