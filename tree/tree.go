package tree

import "strings"

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

func Split(message string, arbor Arbor) Phrase {
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
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "substantivo"
				if GetVerb(unit) {
					spell.Sentence = "predicado"
				} else {
					spell.Sentence = "sujeito"
				}
			}
		}
		for _, value := range arbor.Verb {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "verbo"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "artigo"
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "pronome"
				spell.Sentence = "sujeito"
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "adjetivo"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Adverb {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "adverbio"
				spell.Sentence = "predicado"
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "preposição"
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "artigo"
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) {
				spell.Class = "conjunção"
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Interjection {
			if value.Name == strings.ToLower(spell.Term) {
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
