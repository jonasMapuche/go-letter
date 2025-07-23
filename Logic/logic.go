package logic

import (
	"letter.go/grammar"
)

type Range struct {
	Term    string
	Logical string
	Rule    string
	Value   string
}

type Sense struct {
	Sentence string
	Range    []Range
}

func Math(phrase grammar.Phrase, sentence string) Sense {

	var unit []Range
	var spell Range
	spell.Term = ""
	spell.Logical = ""
	spell.Rule = ""
	spell.Value = ""
	var order bool = false
	var write bool = false
	for _, item := range phrase.Word {
		if item.Sentence == "predicado" && item.Class == "substantivo" {
			spell.Term = item.Term
		}
		if item.Sentence == "predicado" && (item.Class == "adverbio" || item.Class == "adjetivo") {
			spell.Rule = item.Term
		}
		if item.Sentence == "predicado" && item.Class == "conjunção" {
			if order {
				spell.Logical = item.Term
			}
		}
		if item.Sentence == "predicado" && item.Class == "numeral" {
			spell.Value = item.Term
			if !order {
				order = true
				spell.Logical = "and"
			}
			if !write {
				write = true
			}
		}
		if write {
			unit = append(unit, spell)
			write = false
		}
	}
	logic := Sense{
		Sentence: sentence,
		Range:    unit,
	}
	return logic
}
