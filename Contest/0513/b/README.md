## Go言語のスライスまたは文字列で負のインデックスを使用することはできない。
- s[-1]は使えない。
- 文字列の最後の文字を取得するためには、s[len(s)-1]を使用する。


## インデックスを使って文字列から単一の文字を取得するとき、Goはその文字をbyteとして返す

byteとstringを比較すると予期せぬ結果が得られます。
比較を行うときに、byteをstringに変換する必要がある。