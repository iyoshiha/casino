package rule

import (
	"math/rand"
	"time"

	"golang.org/x/xerrors"
)

type Result int

const (
	Tie Result  = 0 
	Player Result = 1
	Banker Result = 2
)

func Baccarrat() (Result, error){
    rand.Seed(time.Now().UnixNano())

    // 乱数を生成し、それに基づいて確率的な出力を行う
    random := rand.Float64() * 100

    if random <= 1 {
		return Tie, nil
    } else if random <= 50{
		return Player, nil
    } else if random <= 100 {
		return Banker, nil
    } else {
        // このブロックは通常は実行されませんが、余分な保険として含めています
		return 0, xerrors.Errorf("not 0 ~ 99:actual: %f", random)
    }
}
