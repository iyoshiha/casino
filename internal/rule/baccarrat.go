package rule

import (
	"errors"
	"math/rand"
	"time"
)

type Result int

const (
	Tie Result  = 0 
	Player Result = 1
	Banker Result = 2
)

func baccarrat() (Result, error){
    rand.Seed(time.Now().UnixNano())

    // 乱数を生成し、それに基づいて確率的な出力を行う
    random := rand.Float64() * 100

    if random <= 1 {
		return 0, nil
    } else if random <= 46 {
		return 1, nil
    } else if random <= 90 {
		return 2, nil
    } else {
        // このブロックは通常は実行されませんが、余分な保険として含めています
		return -1, errors.New("error")
    }
}
