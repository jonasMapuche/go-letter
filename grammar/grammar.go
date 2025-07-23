package grammar

import (
	"strconv"
	"strings"
)

type Verb struct {
	Name     string
	Language string
}

type Noun struct {
	Name     string
	Language string
}

type Pronoun struct {
	Name     string
	Language string
}

type Adjective struct {
	Name     string
	Language string
}

type Adverb struct {
	Name     string
	Language string
}

type Preposition struct {
	Name     string
	Language string
}

type Article struct {
	Name     string
	Language string
}

type Numeral struct {
	Initial  int
	Name     string
	Language string
}

type Conjunction struct {
	Name     string
	Language string
}

type Interjection struct {
	Name     string
	Language string
}

type Arbor struct {
	Noun         []Noun
	Verb         []Verb
	Pronoun      []Pronoun
	Adjective    []Adjective
	Adverb       []Adverb
	Preposition  []Preposition
	Article      []Article
	Numeral      []Numeral
	Interjection []Interjection
	Conjunction  []Conjunction
}

type Word struct {
	Term     string
	Class    string
	Sentence string
}

type Phrase struct {
	Kind string
	Word []Word
}

func GetVerb(word []Word) bool {
	for _, term := range word {
		if term.Class == "verbo" {
			return true
		}
	}
	return false
}

func GetNumeral(numeral []Numeral, word string, language string) bool {
	var amount int = 0
	var vector []string
	for i := 0; i < len(word); i++ {
		vector = append(vector, word[i:i+1])
		for _, term := range numeral {
			if term.Language == language {
				if vector[i] == strconv.Itoa(term.Initial) {
					amount++
				}
			}
		}
	}
	if len(word) == amount {
		return true
	} else {
		return false
	}
}

func Split(message string, arbor Arbor, language string) Phrase {
	var errand string = strings.ToLower(message)
	errand = strings.ReplaceAll(errand, ".", "")
	errand = strings.ReplaceAll(errand, "!", " ")
	errand = strings.ReplaceAll(errand, "?", " ")
	errand = strings.ReplaceAll(errand, ",", " ")

	var word []string = strings.Split(errand, " ")
	var unit []Word

	for _, term := range word {
		var spell Word
		spell.Term = term
		spell.Class = ""
		spell.Sentence = ""
		for _, value := range arbor.Noun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "substantivo"
				if GetVerb(unit) {
					spell.Sentence = "predicado"
				} else {
					spell.Sentence = "sujeito"
				}
			}
		}
		for _, value := range arbor.Verb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "verbo"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "artigo"
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "pronome"
				spell.Sentence = "sujeito"
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "adjetivo"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Adverb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "adverbio"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "preposição"
				if GetVerb(unit) {
					spell.Sentence = "predicado"
				} else {
					spell.Sentence = "sujeito"
				}
			}
		}
		if GetNumeral(arbor.Numeral, spell.Term, language) {
			spell.Class = "numeral"
			if GetVerb(unit) {
				spell.Sentence = "predicado"
			} else {
				spell.Sentence = "sujeito"
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "artigo"
				if GetVerb(unit) {
					spell.Sentence = "predicado"
				} else {
					spell.Sentence = "sujeito"
				}
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "conjunção"
				if GetVerb(unit) {
					spell.Sentence = "predicado"
				} else {
					spell.Sentence = "sujeito"
				}
			}
		}
		for _, value := range arbor.Interjection {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = "interjeição"
				spell.Sentence = ""
			}
		}
		unit = append(unit, spell)
	}

	var class string = Type(message)

	locution := Phrase{
		Kind: class,
		Word: unit,
	}

	return locution
}

func Type(message string) string {
	var class string = "inconsistente"
	if strings.Contains(message, ".") {
		class = "declarativa" // or imperativa (uma ordem)
	} else {
		if strings.Contains(message, "?") {
			class = "interrogativa"
		} else {
			if strings.Contains(message, "!") {
				class = "exclamativa" // or optativa
			}
		}
	}
	return class
}

func Agree(phrase Phrase) Phrase {
	var third bool = false
	var consent bool = false
	for _, value := range phrase.Word {
		if value.Class == "verbo" {
			var vector []string
			var term string = value.Term
			for i := 0; i < len(term); i++ {
				if i == 0 {
					vector = append(vector, "_"+term[i:i+1])
					continue
				}
				if i == len(term)-1 {
					vector = append(vector, term[i:i+1]+"_")
					continue
				}
				vector = append(vector, term[i:i+2])
			}
			for _, charter := range vector {
				if charter == "s_" {
					third = true
					break
				}
				if charter == "es" {
					third = true
					break
				}
			}
		}
	}
	for _, value := range phrase.Word {
		if value.Class == "pronome" {
			if third {
				if strings.ToLower(value.Term) == "he" || strings.ToLower(value.Term) == "she" {
					consent = true
					break
				}
			} else {
				if !(strings.ToLower(value.Term) == "he" || strings.ToLower(value.Term) == "she") {
					consent = true
					break
				}
			}
		}
	}
	if !consent {
		for _, value := range phrase.Word {
			if value.Class == "verbo" {
				var unit []Word
				var spell Word
				spell.Term = value.Term
				spell.Class = value.Class
				spell.Sentence = value.Sentence
				unit = append(unit, spell)
				locution := Phrase{
					Kind: phrase.Kind,
					Word: unit,
				}
				return locution
			}
		}
	}
	return phrase
}
