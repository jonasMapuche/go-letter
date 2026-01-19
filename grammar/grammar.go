package grammar

import (
	"regexp"
	"sort"
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

type Dictionary struct {
	First  string
	Second string
}

type Order struct {
	Key   Dictionary
	Value int
}

type Interpret struct {
	Word  string
	Grade []string
}

const (
	ENGLISH       = "english"
	PREDICATIVE   = "predicado"
	SUBJECT       = "sujeito"
	VERB          = "verbo"
	NOUN          = "substantivo"
	PRONOUN       = "pronome"
	ADVERB        = "adverbio"
	ADJECTIVE     = "adjetivo"
	PREPOSITION   = "preposição"
	NUMERAL       = "numeral"
	ARTICLE       = "artigo"
	CONJUNCTION   = "conjunção"
	INTERJECTION  = "interjeição"
	NOT           = "not"
	DO            = "do"
	TO            = "to"
	ESPECIAL      = "especial"
	DO_NOT        = "do not"
	DECLARATIVE   = "declarativa"
	EXCLAMATORY   = "exclamativa"
	INTERROGATIVE = "interrogativa"
	INCONSISTENT  = "inconsistente"
)

func GetVerb(word []Word) bool {
	for _, term := range word {
		if term.Class == VERB {
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

func GetSpesh(word string) bool {
	var special = `[^a-zA-ZáàâãéèêíïóôõöúçñÁÀÂÃÉÈÊÍÏÓÔÕÖÚÇÑ0-9\s]` //`[^a-zA-Z0-9\s]+`
	if match, _ := regexp.MatchString(special, word); match {
		return true
	}
	return false
}

func GetBefore(word string) bool {
	var special = `[\"]` //\¿\¡]`
	if match, _ := regexp.MatchString(special, word); match {
		return true
	}
	special = `[¿¡]`
	if match, _ := regexp.MatchString(special, word); match {
		return true
	}
	return false
}

func SetDictionary(word []string) map[Dictionary]int {
	var embendding = make(map[Dictionary]int)
	var count int = 0
	var length = len(word)
	for count < length {
		if count+1 == length {
			embendding[Dictionary{word[count], ""}] = count
		} else {
			embendding[Dictionary{word[count], word[count+1]}] = count
		}
		count++
	}
	return embendding
}

func OrderDictionary(embendding map[Dictionary]int) []Order {
	var pares []Order
	for key, value := range embendding {
		pares = append(pares, Order{Key: key, Value: value})
	}
	sort.Slice(pares, func(i, j int) bool {
		return pares[i].Value < pares[j].Value
	})
	return pares
}

func SplitEnglish(word []string, arbor Arbor, language string) []string {
	var dictionary = make(map[Dictionary]int)
	var vocable []string
	var pares []Order

	var preposition string = TO
	var adverb string = DO_NOT
	var salt bool = false

	dictionary = SetDictionary(word)
	pares = OrderDictionary(dictionary)
	for _, order := range pares {
		var term = order.Key.First + " " + order.Key.Second
		if salt {
			salt = false
			continue
		}
		for _, value := range arbor.Verb {
			if value.Name == strings.ToLower(term) && value.Language == language {
				vocable = append(vocable, term)
				salt = true
				break
			}
		}
		if order.Key.First == preposition {
			for _, value := range arbor.Verb {
				if value.Name == strings.ToLower(order.Key.Second) && value.Language == language {
					vocable = append(vocable, term)
					salt = true
					break
				}
			}
		}
		for _, value := range arbor.Adverb {
			if value.Name == strings.ToLower(order.Key.Second) && value.Language == language {
				if strings.ToLower(term) == adverb {
					vocable = append(vocable, term)
					salt = true
					break
				}
			}
		}
		if !salt {
			vocable = append(vocable, order.Key.First)
		}
	}
	return vocable
}

func SplitVerb(word string) string {
	var vocable []string = strings.Split(word, " ")
	var preposition string = TO
	if len(vocable) > 1 {
		if vocable[0] == preposition {
			return vocable[1]
		}
	}
	return word
}

func SplitAdverb(word string) string {
	var vocable []string = strings.Split(word, " ")
	var adverb string = NOT
	var verb string = DO
	if len(vocable) > 1 {
		if vocable[0] == verb && vocable[1] == adverb {
			return vocable[1]
		}
	}
	return word
}

func PrepositionVerb(term string) bool {
	var vocable []string = strings.Split(term, " ")
	var preposition string = TO
	if len(vocable) > 1 {
		if vocable[0] == preposition {
			return true
		}
	}
	return false
}

func SecondVerb(term string, unit []Word) bool {
	var vocable []string = strings.Split(term, " ")
	if len(vocable) == 1 {
		if GetVerb(unit) {
			return true
		} else {
			return false
		}
	}
	return false
}

func DoubleVerb(term string) bool {
	var vocable []string = strings.Split(term, " ")
	if len(vocable) > 1 {
		return true
	}
	return false
}

func AdjunctVerb(word string, arbor Arbor, language string) bool {
	var term string = word
	for _, value := range arbor.Verb {
		var preposition bool = false
		if language == ENGLISH {
			preposition = PrepositionVerb(word)
			if preposition {
				term = SplitVerb(word)
			}
		}
		if value.Name == strings.ToLower(term) && value.Language == language {
			return true
		}
	}
	return false
}

func AdverbialPredicative(term string, arbor Arbor, language string) bool {
	var verb bool = false
	verb = AdjunctVerb(term, arbor, language)
	if verb {
		return true
	}
	return false
}

func Snap(message string, arbor Arbor, language string) []Phrase {
	var errand string = strings.ToLower(message)
	errand = strings.ReplaceAll(errand, ".", " . ")
	errand = strings.ReplaceAll(errand, "!", " ! ")
	errand = strings.ReplaceAll(errand, "?", " ? ")
	errand = strings.ReplaceAll(errand, ",", " , ")
	errand = strings.ReplaceAll(errand, "¿", " ¿ ")
	errand = strings.ReplaceAll(errand, "¡", " ¡ ")
	errand = strings.ReplaceAll(errand, ":", " : ")
	errand = strings.ReplaceAll(errand, ";", " ; ")
	errand = strings.ReplaceAll(errand, "\"", " \" ")

	var word []string = strings.Split(errand, " ")
	var vocable []string
	for _, term := range word {
		if term != "" {
			vocable = append(vocable, term)
		}
	}
	if language == ENGLISH {
		vocable = SplitEnglish(vocable, arbor, language)
	}
	var unit []Word
	var phrase []Phrase
	for _, term := range vocable {
		var spell Word
		spell.Term = term
		spell.Class = ""
		spell.Sentence = ""
		unit = append(unit, spell)
		if GetSpesh(spell.Term) {
			if GetBefore(spell.Term) {
				continue
			}
			var locution []Word = Slash(unit, arbor, language)
			var kind string = Type(spell.Term)

			var clause Phrase = Phrase{
				Kind: kind,
				Word: locution,
			}
			phrase = append(phrase, clause)
			unit = nil
		}
	}
	return phrase
}

func Slash(word []Word, arbor Arbor, language string) []Word {
	var unit []Word
	var locution []Word
	locution = word
	for order, term := range word {
		var spell Word
		spell.Term = term.Term
		spell.Class = term.Class
		spell.Sentence = term.Sentence
		if GetSpesh(spell.Term) {
			spell.Class = ESPECIAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATIVE
			} else {
				spell.Sentence = SUBJECT
			}
		}
		for _, value := range arbor.Verb {
			var term string = spell.Term
			var preposition bool = false
			var second bool = false
			var double bool = false
			if language == ENGLISH {
				preposition = PrepositionVerb(spell.Term)
				if preposition {
					term = SplitVerb(spell.Term)
				} else {
					second = SecondVerb(term, unit)
					double = DoubleVerb(term)
				}
			}
			if value.Name == strings.ToLower(term) && value.Language == language && spell.Sentence == "" {
				if preposition || !second || double {
					spell.Class = VERB
					spell.Sentence = PREDICATIVE
				}
			}
		}
		for _, value := range arbor.Noun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = NOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = PRONOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = ADJECTIVE
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Adverb {
			var term string = spell.Term
			if language == ENGLISH {
				term = SplitAdverb(spell.Term)
			}
			if value.Name == strings.ToLower(term) && value.Language == language && spell.Sentence == "" {
				spell.Class = ADVERB
				var predicative bool = AdverbialPredicative(locution[order+1].Term, arbor, language)
				if GetVerb(unit) || predicative {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = PREPOSITION
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		if GetNumeral(arbor.Numeral, spell.Term, language) && spell.Sentence == "" {
			spell.Class = NUMERAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATIVE
			} else {
				spell.Sentence = SUBJECT
			}
		}
		for _, value := range arbor.Numeral {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = NUMERAL
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = ARTICLE
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = CONJUNCTION
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Interjection {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = INTERJECTION
				spell.Sentence = ""
			}
		}
		unit = append(unit, spell)
	}
	return unit
}

func Translate(message string, arbor Arbor, language string) []Interpret {
	var word []string = strings.Split(message, " ")
	var vocable []string
	for _, term := range word {
		if term != "" {
			vocable = append(vocable, term)
		}
	}
	if language == ENGLISH {
		vocable = SplitEnglish(vocable, arbor, language)
	}
	var unit []Interpret
	for _, term := range vocable {
		var spell Interpret
		spell.Word = term
		if GetSpesh(spell.Word) {
			spell.Grade = append(spell.Grade, ESPECIAL)
		}
		for _, value := range arbor.Verb {
			var term string = spell.Word
			var preposition bool = false
			if language == ENGLISH {
				preposition = PrepositionVerb(spell.Word)
				if preposition {
					term = SplitVerb(spell.Word)
				}
			}
			if value.Name == strings.ToLower(term) && value.Language == language {
				spell.Grade = append(spell.Grade, VERB)
			}
		}
		for _, value := range arbor.Noun {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, NOUN)
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, PRONOUN)
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, ADJECTIVE)
			}
		}
		for _, value := range arbor.Adverb {
			var term string = spell.Word
			if language == "english" {
				term = SplitAdverb(spell.Word)
			}
			if value.Name == strings.ToLower(term) && value.Language == language {
				spell.Grade = append(spell.Grade, ADVERB)
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, PREPOSITION)
			}
		}
		if GetNumeral(arbor.Numeral, spell.Word, language) {
			spell.Grade = append(spell.Grade, NUMERAL)
		}
		for _, value := range arbor.Numeral {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, NUMERAL)
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, ARTICLE)
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, CONJUNCTION)
			}
		}
		for _, value := range arbor.Interjection {
			if value.Name == strings.ToLower(spell.Word) && value.Language == language {
				spell.Grade = append(spell.Grade, INTERJECTION)
			}
		}
		unit = append(unit, spell)
	}
	return unit
}

func Split(message string, arbor Arbor, language string) Phrase {
	var errand string = strings.ToLower(message)
	errand = strings.ReplaceAll(errand, ".", " . ")
	errand = strings.ReplaceAll(errand, "!", " ! ")
	errand = strings.ReplaceAll(errand, "?", " ? ")
	errand = strings.ReplaceAll(errand, ",", " , ")
	errand = strings.ReplaceAll(errand, "¿", " ¿ ")
	errand = strings.ReplaceAll(errand, "¡", " ¡ ")
	errand = strings.ReplaceAll(errand, ":", " : ")
	errand = strings.ReplaceAll(errand, ";", " ; ")
	errand = strings.ReplaceAll(errand, "\"", " \" ")

	var word []string = strings.Split(errand, " ")
	var unit []Word

	for _, term := range word {
		var spell Word
		spell.Term = term
		spell.Class = ""
		spell.Sentence = ""

		if GetSpesh(spell.Term) {
			spell.Class = ESPECIAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATIVE
			} else {
				spell.Sentence = SUBJECT
			}
		}

		for _, value := range arbor.Noun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = NOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Verb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = VERB
				spell.Sentence = PREDICATIVE
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ARTICLE
				spell.Sentence = ""
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = PRONOUN
				spell.Sentence = SUBJECT
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ADJECTIVE
				spell.Sentence = PREDICATIVE
			}
		}
		for _, value := range arbor.Adverb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ADVERB
				spell.Sentence = PREDICATIVE
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = PREPOSITION
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		if GetNumeral(arbor.Numeral, spell.Term, language) {
			spell.Class = NUMERAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATIVE
			} else {
				spell.Sentence = SUBJECT
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ARTICLE
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = CONJUNCTION
				if GetVerb(unit) {
					spell.Sentence = PREDICATIVE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Interjection {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = INTERJECTION
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
	var class string = INCONSISTENT
	if strings.Contains(message, ".") ||
		strings.Contains(message, ":") ||
		strings.Contains(message, ",") {
		class = DECLARATIVE // or imperativa (uma ordem)
	} else {
		if strings.Contains(message, "?") {
			class = INTERROGATIVE
		} else {
			if strings.Contains(message, "!") {
				class = EXCLAMATORY // or optativa
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
