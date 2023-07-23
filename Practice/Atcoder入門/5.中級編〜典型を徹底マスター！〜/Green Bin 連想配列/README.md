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

- ABC 164 C - gacha
- ABC 154 C - Distinct or Not
- ABC 118 B - Foods Loved by Everyone
- ABC 081 C - Not so Diverse
- 競プロ典型 90 問 027 Sign Up Requests (# 2)
- ABC 155 C - Poll
- ABC 073 C - Write and Erase
- ABC 082 C - Good Sequence
- ABC 058C - 怪文書
- ABC 206 C - Swappable
- ABC 200 C - Ringo's Favorite Numbers 2
- 競プロ典型 90 間 046 - I Love 46 (\* 3)
