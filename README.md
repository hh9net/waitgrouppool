# WaitGroupPool

对`sync.WaitGroup`加以封装，以使得程序可以优雅地限制`goroutine`的数量，应用场景举例：

* 场景1：限制并行发送消息的数量
* 场景2：限制并行TCP连接的数量（因为打开文件描述符有限制，如限制1024个连接）

# 使用方法

使用起来与`sync.WaitGroup`很相近，没有多大的区别。

```go
package main

import (
	"fmt"

	"github.com/scue/waitgrouppool"
)

func main() {
	pool := waitgrouppool.New(2)

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		pool.Add()
		go func(c int) {
			defer pool.Done()
			fmt.Printf("hello %d\n", c)
		}(i)

	}

	// wait until we call JobDone for all jobs
	pool.Wait()
}
```

Output:

```txt
hello 1
hello 2
hello 3
hello 0
hello 5
hello 6
hello 7
hello 8
hello 9
hello 4
```