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

- ABC 143 C - Slimes
- ABC 019B - 高橋 1300 くんと文字列圧縮当
- ABC 139 C - Lower 300
- ABC 116 C - Grand Garden
- AGC 026 A - Colorful Slimes 2
- AGC 040 A - ><
- AGC 013 A - Sorted Arrays
- AGC 039 A - Connection and Disconnection
- ABC 136 D - Gathering Children
- AGC 016 A - Shrinking
- 競プロ典型 90 問 084 - There are two types of characters (\* 3)
