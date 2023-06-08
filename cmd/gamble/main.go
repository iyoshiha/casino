package main

import (
	"fmt"
	"log"

	"github.com/casino/internal/rule"
	"golang.org/x/xerrors"
)


type result struct{
	wins int
	looses int
	ties int
}

func main(){
	res, err := start(rule.Baccarrat, 100)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(res)
}

func start(game func()(rule.Result, error), times int) (*result, error){
	var res = &result{}
	for i := 0; i < times; i++{
		ans, err := game()
		if err != nil {
			return res, xerrors.Errorf("start failed %w",  err)
		}

		switch ans {
		case rule.Player:
			res.won()
		case rule.Banker:
			res.lost()
		case rule.Tie:
			res.tie()
		}
	}

	return res, nil
}

func (r *result) won(){
	r.wins++
}
func (r *result) lost(){
	r.looses++
}
func (r *result) tie(){
	r.ties++
}
