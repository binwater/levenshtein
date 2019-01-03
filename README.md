## levenshtein

go语言实现的Levenshtein Distance 算法及基于Levenshtein Distance 算法的相似度计算算法

## USAGE

```
package main

import (
	"fmt"

	"github.com/cloudfstrife/levenshtein"
)

func main() {
	fmt.Println(levenshtein.Levenshtein("你好世界", "晚安世界"))
	fmt.Println(levenshtein.SimilarDegree("你好世界", "晚安世界"))
}
```

输出：

```
$ go run main.go
2
0.5
```