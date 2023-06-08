package main

import (
	"github.com/casino/rule"
	
)

type result struct{
	player_wins int
	banker_wins int
	tie_numbers int
}

func main(){
	
}

func start(game func()(int, error), times int) result{
	for i := 0; i < times; i++{
		a, err := game()
		if err != nil {
			return nil
		}

		switch a {
		case rule.
		}
	}

	return result{}
}

