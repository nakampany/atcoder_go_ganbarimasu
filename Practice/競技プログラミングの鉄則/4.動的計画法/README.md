# 動的計画法（DP）

- ナップサック問題
- 迷路などの最短路問題
- 区間スケジューリング問題
- 音声認識パターンマッチング問題
- レーベンシュタイン距離
- 発電計画問題
- 分かち書き
- 隠れマルコフモデル

## 大きな問題を小さな部分問題に分解して解くためのアルゴリズム設計手法

1. 動的計画法の特徴:
   大きな問題を小さな部分問題に分けて考える。
   各部分問題の答えをメモ化（記録）しておき、同じ問題を 2 回以上解かないようにする。

2. 簡単な例: フィボナッチ数列

   - フィボナッチ数列とは 前の２つの和が次の数字になる数列。
     1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233,

   フィボナッチ数列は、次のように定義されます。
   F(0)=0,F(1)=1,F(n)=F(n−1)+F(n−2)

   再帰的にフィボナッチ数列を求めるアルゴリズムは次のようになります。

   ```python
   def fibonacci(n):
      if n <= 1:
         return n
      return fibonacci(n-1) + fibonacci(n-2)
   ```

   しかし、このアルゴリズムは非常に効率が悪いです。なぜなら、同じ計算を何度も繰り返しているからです。
   しかしこの方法では計算量が多すぎて、まともに結果が返ってくるのはせいぜい４０番目ぐらいまで。

   動的計画法を使うと、以前に計算した結果を再利用して、計算時間を大幅に削減することができます。

3. 動的計画法を使ったフィボナッチ数列:

   - トップダウン：メモ化再帰

   ```python
   f = []

   def fib2(n):
      if n <= 1:
         return 1

      // n 番目のフィボナッチ数がメモ化されているかをチェック
      if len(f) > n and f[n] is not None:
         return f[n]

      // n 番目のフィボナッチ数を計算
      while len(f) <= n:
         f.append(None)
      f[n] = fib2(n-1) + fib2(n-2)

      return f[n]

   // 使用例
   print(fib2(10)) # 89 を出力
   ```

   go

   ```go
   package main

   import "fmt"

   var memo map[int]int = make(map[int]int)

   func fib2(n int) int {
      if n <= 1 {
         return 1
      }

      // n 番目のフィボナッチ数がメモ化されているかをチェック
      if result, found := memo[n]; found {
         return result
      }

      // n 番目のフィボナッチ数を計算
      memo[n] = fib2(n-1) + fib2(n-2)

      return memo[n]
   }

   func main() {
      fmt.Println(fib2(10))  // 89 を出力
   }
   ```

   - ボトムアップ: 分割統治法、漸化式ループ

   ```python
   def fib3(n):
    if n == 0:
        return 1
    else:
        fib1, fib2, fib3 = 1, 1, 1
        for _ in range(n-1):
            fib3 = fib1 + fib2
            fib1, fib2 = fib2, fib3
        return fib2

   // 使用例
   print(fib3(10))  # 89 を出力
   ```

このアルゴリズムは、memo というリストを使用して、一度計算した結果を保存しています。これにより、同じ計算を繰り返すことなく、効率的に答えを得ることができます。

### 参考

- https://jabba.cloud/20161020172918

```

```
