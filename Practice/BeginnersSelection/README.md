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
奇数が見つかるまで、for文で回す

	for {
		for i := 0; i < n; i++ {
			if a[i]%2 == 1 {
				return
			}
			a[i] /= 2
		}
		count++
	}

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
桁の足し算

		for num > 0 {
			total += num % 10
			num /= 10
		}

```
- ABC088B Card Game for Two
```
降順でソートして、Aliceが最初に引く

	sort.Slice(card, func(i, j int) bool { return card[i] > card[j] })

	fmt.Println(card)

	var alice, bob int

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("alice")
			fmt.Println(card[i])
			alice += card[i]
		} else {
			fmt.Println("bob")
			fmt.Println(card[i])
			bob += card[i]
		}
	}
```
- ABC085B Kagami Mochi
```
配列内の同じものを削除する
新しい配列に格納
	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
 
	se := make(map[int]bool)
	for i := 0; i < n; i++ {
		se[t[i]] = true
	}
	ans := len(se)
	fmt.Println(ans)
```
- ABC085C Otoshidama
- ABC049C 白昼夢
- ABC086C Traveling

a1,a2,a3...aN

```
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
```
