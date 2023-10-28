# atcoder_go_ganbarimasu

Atcoder を Golang で書いていくよ

# 標準入力

## fmt.Scan(&a[i])

```go
var n int
fmt.Scan(&n)
```

### slice

```go
a := make([]string, n)
for i := 0; i < n; i++ {
	fmt.Scan(&a[i])
}
```

```go
a := make([][]int, m)
for i := 0; i < m; i++ {
	a[i] = make([]int, n)
	for j := 0; j < n; j++ {
		fmt.Scan(&a[i][j])
	}
}
```

```go
a := make([][]int, n)
for i := 0; i < n; i++ {
	a[i] = make([]int, n)
	var s string
	fmt.Scan(&s)
	for j := 0; j < n; j++ {
		a[i][j], _ = strconv.Atoi(string(s[j])) // 文字列の特定の文字を整数に変換する操作
	}
}
```

### map

```go
person := make(map[string]int)

for i := 0; i < n; i++ {
	var name string
	var age int
	fmt.Scan(&name, &age)
	person[name] = age
}
```

```go
type Person struct {
	name string
	age int
}
person := []Person{}

for i := 0; i < n; i++ {
	var name string
	var age int
	fmt.Scan(&name, &age)
	person[name] = age
	// person = append(person, Person{x, y})
}
```

## scanner := bufio.NewScanner(os.Stdin)

```go
// int型の入力
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
inputInt, _ := strconv.Atoi(scanner.Text())

// fmt.Println(inputInt)

// string型の入力
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
inputString := scanner.Text()

// fmt.Println(inputString)
```

### slice

inputSlice := strings.Fields(scanner.Text())
引数に与えられた文字列をスペースで分割し、結果の各部分を要素とする文字列のスライスを返します。この場合、scanner.Text()は Scanner が最後に読み込んだ行（つまり、ユーザーからの入力）を返します。

整数の空のスライスを作成します。

var numbers []int
numbers := make([]int, 0)
numbers := []int{}

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
inputSlice := strings.Fields(scanner.Text())
var numbers []int
for _, v := range inputSlice {
	num, _ := strconv.Atoi(v)
	numbers = append(numbers, num)
}
fmt.Println(numbers)
```

### map

```go
n := 3
type Point struct {
	x int
	y int
}
points := []Point{}
scanner := bufio.NewScanner(os.Stdin)
for i := 0; i < n; i++ {
	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	points = append(points, Point{x, y})
}

fmt.Println(points)
```
