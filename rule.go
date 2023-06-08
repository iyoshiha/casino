```golang
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    // 乱数を生成し、それに基づいて確率的な出力を行う
    random := rand.Float64() * 100

    if random <= 1 {
        fmt.Println(0)
    } else if random <= 46 {
        fmt.Println(1)
    } else if random <= 90 {
        fmt.Println(2)
    } else {
        // このブロックは通常は実行されませんが、余分な保険として含めています
        fmt.Println("Invalid probability range")
    }
}
```