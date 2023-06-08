package main

import (
	"github.com/casino/internal/rule"
)

type result struct{
	wins int
	looses int
	ties int
}

func main(){
	
}

func start(game func()(rule.Result, error), times int) result, error{
	for i := 0; i < times; i++{
		a, err := game()
		if err != nil {
			return a, nil
		}

		switch a {
		case rule.Player:
		case rule.Banker:
		case rule.Tie:
		}
	}

	return result{}
}

