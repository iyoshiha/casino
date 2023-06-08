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

func	baccarrat() (Result, error){
    rand.Seed(time.Now().UnixNano())

    // 乱数を生成し、それに基づいて確率的な出力を行う
    random := rand.Float64() * 100

    if random <= 1 {
		return Tie, nil
    } else if random <= 49{
		return Player, nil
    } else if random <= 99 {
		return banker, nil
    } else {
        // このブロックは通常は実行されませんが、余分な保険として含めています
		return 0, errors.New("error")
    }
}
