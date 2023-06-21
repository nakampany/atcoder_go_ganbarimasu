https://atcoder.jp/contests/abs/tasks

- PracticeA Welcome to AtCoder
- ABC086A Product
```
文字列の初めから見ていき、１をカウントする
```
- ABC081A Placing Marbles
```
a*bが偶数だったら、Even
a*bが奇数だったら、Odd

a*b%2 == 0
```
- ABC081B Shift only
```
```
- ABC087B Coins
```
for文を3回回して、全通り試す。
for i := 0; i <= a; i++ {
	for j := 0; j <= b; j++ {
		for k := 0; k <= c; k++ {
			if 500*i+100*j+50*k == x {
				ans++
			}
		}
	}
}
```
- ABC083B Some Sums
```
```
- ABC088B Card Game for Two
```
```
- ABC085B Kagami Mochi
```
```
- ABC085C Otoshidama
- ABC049C 白昼夢
- ABC086C Traveling

a1,a2,a3...aN

```
	card := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&card[i])
	}
```
