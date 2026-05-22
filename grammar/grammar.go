package grammar

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"letter.go/brand"
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

type Glossary struct {
	First  Word
	Second Word
	Order  int
}

type Vocabulary struct {
	Jargon []Glossary
}

type Talk struct {
	Term     string
	Etiology []string
	Pattern  []string
	Order    int
}

type Recite struct {
	Kind string
	Talk []Talk
}

type Poll struct {
	Word []Word
}

const (
	ENGLISH       = "english"
	PREDICATE     = "predicado"
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
	MISSING       = "inexistente"
	SAMPLE        = "simples"
	COMPOUND      = "composto"
	UNKNOW        = "indefinido"
	RELATIVE      = "relativo"
	ADNOMINAL     = "adjunto adnominal"
	ADVERBIAL     = "adjunto adverbial"
	PREDICATIVE   = "predicativo"
	DIRECT        = "objeto direto"
	INDIRECT      = "objeto indireto"
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

func SplitNoun(word string) string {
	var vocable []string = strings.Split(word, " ")
	var noun string = ""
	if len(vocable) > 1 {
		noun = vocable[1]
	} else {
		noun = vocable[0]
	}
	return noun
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
	var predicate bool = false
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
			if TypeEspecial(spell.Term) {
				predicate = true
			}
			var locution []Word = Slash(unit, arbor, language, predicate)
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

func Slash(word []Word, arbor Arbor, language string, predicate bool) []Word {
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
				spell.Sentence = PREDICATE
			} else {
				if predicate {
					spell.Sentence = PREDICATE
				} else {
					spell.Sentence = SUBJECT
				}
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
					spell.Sentence = PREDICATE
				}
			}
		}
		for _, value := range arbor.Noun {
			var term string = SplitNoun(value.Name)
			if term == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = NOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		for _, value := range arbor.Pronoun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = PRONOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		for _, value := range arbor.Adjective {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = ADJECTIVE
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
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
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = PREPOSITION
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		if GetNumeral(arbor.Numeral, spell.Term, language) && spell.Sentence == "" {
			spell.Class = NUMERAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATE
			} else {
				if predicate {
					spell.Sentence = PREDICATE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Numeral {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = NUMERAL
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = ARTICLE
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
				}
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language && spell.Sentence == "" {
				spell.Class = CONJUNCTION
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					if predicate {
						spell.Sentence = PREDICATE
					} else {
						spell.Sentence = SUBJECT
					}
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
			var term string = SplitNoun(value.Name)
			if term == strings.ToLower(spell.Word) && value.Language == language {
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
				spell.Sentence = PREDICATE
			} else {
				spell.Sentence = SUBJECT
			}
		}

		for _, value := range arbor.Noun {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = NOUN
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Verb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = VERB
				spell.Sentence = PREDICATE
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
				spell.Sentence = PREDICATE
			}
		}
		for _, value := range arbor.Adverb {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ADVERB
				spell.Sentence = PREDICATE
			}
		}
		for _, value := range arbor.Preposition {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = PREPOSITION
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		if GetNumeral(arbor.Numeral, spell.Term, language) {
			spell.Class = NUMERAL
			if GetVerb(unit) {
				spell.Sentence = PREDICATE
			} else {
				spell.Sentence = SUBJECT
			}
		}
		for _, value := range arbor.Article {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = ARTICLE
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
				} else {
					spell.Sentence = SUBJECT
				}
			}
		}
		for _, value := range arbor.Conjunction {
			if value.Name == strings.ToLower(spell.Term) && value.Language == language {
				spell.Class = CONJUNCTION
				if GetVerb(unit) {
					spell.Sentence = PREDICATE
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

func OrderFirst(glossaries []Glossary, word Word) int {
	var minus int = 0
	for _, noun := range glossaries {
		if noun.First == word {
			var num = noun.Order
			if num < minus {
				minus = num
			}
		}
	}
	return minus
}

func OrderLast(glossaries []Glossary, word Word) int {
	var maximus int = 0
	for _, noun := range glossaries {
		if noun.First == word {
			var max = noun.Order
			if max > maximus {
				maximus = max
			}
		}
	}
	return maximus
}

func FilterGlossary(glossaries []Glossary, word Word) []Glossary {
	var lexicons []Glossary
	for _, glossary := range glossaries {
		if glossary.First == word {
			var lexicon Glossary
			lexicon.First = glossary.First
			lexicon.Second = glossary.Second
			lexicon.Order = glossary.Order
			lexicons = append(lexicons, lexicon)
		}
	}
	return lexicons
}

func MountVocabulary(glossaries []Glossary, crescent bool) []Vocabulary {
	if glossaries == nil {
		return nil
	}
	sort.Slice(glossaries, func(i, j int) bool {
		return glossaries[i].First.Term < glossaries[j].First.Term
	})
	var vocabularies []Vocabulary
	var before string = ""
	var vocabulary Vocabulary
	for _, association := range glossaries {
		if association.Order <= 0 {
			if before == "" {
				vocabulary.Jargon = append(vocabulary.Jargon, association)
				before = association.First.Term
			} else {
				if before != association.First.Term {
					if crescent {
						sort.Slice(vocabulary.Jargon, func(i, j int) bool {
							return vocabulary.Jargon[i].Order > vocabulary.Jargon[j].Order
						})
					} else {
						sort.Slice(vocabulary.Jargon, func(i, j int) bool {
							return vocabulary.Jargon[i].Order < vocabulary.Jargon[j].Order
						})
					}
					vocabularies = append(vocabularies, vocabulary)
					vocabulary.Jargon = nil
					before = association.First.Term
				} else {
					vocabulary.Jargon = append(vocabulary.Jargon, association)
				}
			}
		}
	}
	if vocabulary.Jargon != nil {
		if crescent {
			sort.Slice(vocabulary.Jargon, func(i, j int) bool {
				return vocabulary.Jargon[i].Order > vocabulary.Jargon[j].Order
			})
		} else {
			sort.Slice(vocabulary.Jargon, func(i, j int) bool {
				return vocabulary.Jargon[i].Order < vocabulary.Jargon[j].Order
			})
		}
		vocabularies = append(vocabularies, vocabulary)
	}
	return vocabularies
}

func MountPreposition(noun Word, prepositions []Glossary) []Word {
	var link Word
	for _, preposition := range prepositions {
		if preposition.Order == 0 {
			if preposition.Second == noun {
				link = preposition.First
			}
		}
	}
	var terms []Word
	if link.Term != "" {
		terms = append(terms, link)
		terms = append(terms, noun)
	} else {
		terms = append(terms, noun)
	}
	return terms
}

func MountNoun(nouns []Glossary, prepositions []Glossary) []Talk {
	if nouns == nil {
		return nil
	}

	var talks []Talk
	var vocabularies = MountVocabulary(nouns, true)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var noun Word = glossaries[0].First
		var article Word
		var pronome Word
		var numeral Word
		var adverb Word
		var adverb_adverb Word
		var adjetivo Word
		var preposition Word

		var quantity int = len(glossaries)
		var filters []Glossary
		for count := 0; count < quantity; count++ {
			if glossaries[count].Order <= 0 {
				filters = append(filters, glossaries[count])
			}
		}
		sort.Slice(filters, func(i, j int) bool {
			return filters[i].Order > filters[j].Order
		})
		var size int = len(filters)
		for count := 0; count < size; count++ {
			var second Word = filters[count].Second
			var order int = filters[count].Order

			var terms []Word = MountPreposition(second, prepositions)
			var lenght = len(terms)
			if lenght > 1 {
				preposition = terms[0]
			}

			if second.Class == NOUN && order != 0 {
				break
			}
			if second.Class == VERB && order != 0 {
				break
			}
			if second.Class == CONJUNCTION && order != 0 {
				break
			}
			if second.Class == ARTICLE && order != 0 {
				article = second
				break
			}
			if second.Class == ADVERB {
				adverb = second
			}
			if second.Class == ADVERB && adverb.Term != "" {
				adverb_adverb = second
			}
			if second.Class == ADJECTIVE {
				adjetivo = second
			}
			if second.Class == NUMERAL {
				numeral = second
			}
			if second.Class == PRONOUN {
				pronome = second
			}
		}

		var words []Word
		if preposition.Term != "" {
			words = append(words, preposition)
		}
		if article.Term != "" {
			words = append(words, article)
		}
		if adjetivo.Term != "" {
			words = append(words, adjetivo)
		}
		if adverb.Term != "" {
			words = append(words, adverb)
		}
		if adverb_adverb.Term != "" {
			words = append(words, adverb_adverb)
		}
		if numeral.Term != "" {
			words = append(words, numeral)
		}
		if pronome.Term != "" {
			words = append(words, pronome)
		}
		if noun.Term != "" {
			words = append(words, noun)
		}

		for _, word := range words {
			var talk Talk
			talk.Term = word.Term
			talk.Etiology = append(talk.Etiology, word.Class)
			talk.Pattern = append(talk.Pattern, word.Sentence)
			talk.Pattern = append(talk.Pattern, ADNOMINAL)
			if word.Sentence == PREDICATE {
				if preposition.Term != "" {
					talk.Pattern = append(talk.Pattern, INDIRECT)
				} else {
					talk.Pattern = append(talk.Pattern, DIRECT)
				}
			}
			talks = append(talks, talk)
		}
	}
	return talks
}

func MountVerb(verbs []Glossary) []Talk {
	if verbs == nil {
		return nil
	}

	var talks []Talk
	var vocabularies = MountVocabulary(verbs, true)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var verb Word = glossaries[0].First
		var adverb Word
		var adverb_adverb Word

		var quantity int = len(glossaries)
		var filters []Glossary
		for count := 0; count < quantity; count++ {
			if glossaries[count].Order >= 0 {
				filters = append(filters, glossaries[count])
			}
		}
		sort.Slice(filters, func(i, j int) bool {
			return filters[i].Order < filters[j].Order
		})
		var size int = len(filters)
		for count := 0; count < size; count++ {
			var second Word = filters[count].Second
			var order int = filters[count].Order
			if second.Class == NOUN {
				break
			}
			if second.Class == NUMERAL {
				break
			}
			if second.Class == ADJECTIVE {
				break
			}
			if second.Class == VERB && order != 0 {
				break
			}
			if second.Class == ADVERB && adverb.Term != "" {
				adverb_adverb = second
				break
			}
			if second.Class == ADVERB {
				adverb = second
			}
		}

		var words []Word
		if adverb.Term != "" {
			words = append(words, adverb)
		}
		if adverb_adverb.Term != "" {
			words = append(words, adverb_adverb)
		}
		if verb.Term != "" {
			words = append(words, verb)
		}

		for _, word := range words {
			var talk Talk
			talk.Term = word.Term
			talk.Etiology = append(talk.Etiology, word.Class)
			talk.Pattern = append(talk.Pattern, word.Sentence)
			talk.Pattern = append(talk.Pattern, ADVERBIAL)
			talks = append(talks, talk)
		}
	}
	return talks
}

func MountAdverb(glossaries []Glossary, adverbs []Glossary) []Talk {
	if glossaries == nil {
		return nil
	}

	var filter Word
	for _, adverb := range adverbs {
		if adverb.Second.Class == ADJECTIVE {
			filter = adverb.Second
			break
		}
	}
	var adjectives []Glossary = FilterGlossary(glossaries, filter)

	var talks []Talk
	var vocabularies = MountVocabulary(adjectives, true)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var adjective Word = glossaries[0].First
		var adverb Word
		var adverb_adverb Word

		var quantity int = len(glossaries)
		var filters []Glossary
		for count := 0; count < quantity; count++ {
			if glossaries[count].Order <= 0 {
				filters = append(filters, glossaries[count])
			}
		}
		sort.Slice(filters, func(i, j int) bool {
			return filters[i].Order < filters[j].Order
		})
		var size int = len(filters)
		for count := 0; count < size; count++ {
			var second Word = filters[count].Second
			if second.Class == ADJECTIVE {
				break
			}
			if second.Class == NOUN {
				break
			}
			if second.Class == VERB {
				break
			}
			if second.Class == NUMERAL {
				break
			}
			if second.Class == ARTICLE {
				break
			}
			if second.Class == ADVERB && adverb.Term != "" {
				adverb_adverb = second
				break
			}
			if second.Class == ADVERB {
				adverb = second
			}
		}

		var words []Word
		if adverb.Term != "" {
			words = append(words, adverb)
		}
		if adverb_adverb.Term != "" {
			words = append(words, adverb_adverb)
		}
		if adjective.Term != "" {
			words = append(words, adjective)
		}

		for _, word := range words {
			var talk Talk
			talk.Term = word.Term
			talk.Etiology = append(talk.Etiology, word.Class)
			talk.Pattern = append(talk.Pattern, word.Sentence)
			talk.Pattern = append(talk.Pattern, ADVERBIAL)
			talks = append(talks, talk)
		}
	}
	return talks
}

func TypeEspecial(value string) bool {
	if value == "," || value == ";" || value == ":" {
		return true
	}
	return false
}

func MountConnection(connections []Glossary, syntax string, dome brand.Arbor, language string) []Vocabulary {
	var associations []Glossary
	for _, association := range connections {
		if syntax == SUBJECT {
			if association.First.Class == CONJUNCTION {
				var lexicon Glossary
				lexicon.First = association.First
				lexicon.Second = association.Second
				lexicon.Order = association.Order
				associations = append(associations, lexicon)
			}
		} else {
			if association.First.Class == CONJUNCTION || TypeEspecial(association.First.Term) {
				var lexicon Glossary
				lexicon.First = association.First
				lexicon.Second = association.Second
				lexicon.Order = association.Order
				associations = append(associations, lexicon)
			}
			if association.First.Class == PRONOUN {
				for _, value := range dome.Pronoun {
					if value.Name == strings.ToLower(association.First.Term) && value.Type == RELATIVE &&
						value.Language == language {
						var first Word
						first.Sentence = association.First.Sentence
						first.Term = association.First.Term
						first.Class = RELATIVE
						var lexicon Glossary
						lexicon.First = first
						lexicon.Second = association.Second
						lexicon.Order = association.Order
						associations = append(associations, lexicon)
					}
					if value.Name == strings.ToLower(association.First.Term) && value.Type == UNKNOW &&
						value.Language == language {
						var first Word
						first.Sentence = association.First.Sentence
						first.Term = association.First.Term
						first.Class = UNKNOW
						var lexicon Glossary
						lexicon.First = first
						lexicon.Second = association.Second
						lexicon.Order = association.Order
						associations = append(associations, lexicon)
					}
				}
			}
		}
	}

	var vocabularies []Vocabulary = MountVocabulary(associations, true)

	if vocabularies == nil {
		return nil
	}
	if syntax == SUBJECT {
		return vocabularies
	}
	if syntax == PREDICATE {
		var specials []Word
		var before string = ""
		for _, special := range connections {
			if TypeEspecial(special.First.Term) {
				if before == "" {
					before = special.First.Term
					specials = append(specials, special.First)
				} else {
					if before != special.First.Term {
						specials = append(specials, special.First)
						before = special.First.Term
					} else {
						specials = append(specials, special.First)
					}
				}
			}
		}
		var prepositions []Word
		var ahead string = ""
		for _, preposition := range connections {
			if preposition.First.Class == PREPOSITION {
				if before == "" {
					ahead = preposition.First.Term
					prepositions = append(prepositions, preposition.First)
				} else {
					if ahead == preposition.First.Term {
						prepositions = append(prepositions, preposition.First)
						ahead = preposition.First.Term
					} else {
						prepositions = append(prepositions, preposition.First)
					}
				}
			}
		}
		var dictionaries []Vocabulary
		for _, vocabulary := range vocabularies {
			var dictionary Vocabulary
			dictionary.Jargon = append(dictionary.Jargon, vocabulary.Jargon...)
			var contact bool = false
			var sequence int = -1
			for _, glossary := range vocabulary.Jargon {
				if glossary.First.Class == CONJUNCTION {
					for _, preposition := range prepositions {
						if glossary.Order == sequence && glossary.Second == preposition {
							var filters = FilterGlossary(connections, preposition)
							dictionary.Jargon = append(dictionary.Jargon, filters...)
							contact = true
							break
						}
					}
				}
				if glossary.First.Class == UNKNOW {
					for _, preposition := range prepositions {
						if glossary.Order == sequence && glossary.Second == preposition {
							var filters = FilterGlossary(connections, preposition)
							dictionary.Jargon = append(dictionary.Jargon, filters...)
							contact = true
							break
						}
					}
				}
				if glossary.First.Class == RELATIVE {
					for _, preposition := range prepositions {
						if glossary.Order == sequence && glossary.Second == preposition {
							var filters = FilterGlossary(connections, preposition)
							dictionary.Jargon = append(dictionary.Jargon, filters...)
							contact = true
							break
						}
					}
				}
			}
			for _, glossary := range vocabulary.Jargon {
				if glossary.First.Class == CONJUNCTION {
					if contact == true {
						sequence--
					}
					for _, special := range specials {
						if glossary.Order == sequence && glossary.Second == special {
							var filters = FilterGlossary(connections, special)
							dictionary.Jargon = append(dictionary.Jargon, filters...)
							break
						}
					}
				}
			}
			dictionaries = append(dictionaries, dictionary)
		}
		return dictionaries
	}
	return nil
}

func TypeSubject(nouns []Glossary, connections []Vocabulary) string {
	if nouns == nil {
		return MISSING
	}
	if connections != nil {
		return COMPOUND
	}
	return SAMPLE
}

func NounCompound(terms []Word, connection Word, captions []Talk) []Talk {
	var expressions []Talk
	var lenght int = 0
	if terms != nil {
		lenght = len(terms)
		for _, term := range terms {
			var talk Talk
			talk.Term = term.Term
			talk.Etiology = append(talk.Etiology, term.Class)
			talk.Pattern = append(talk.Pattern, term.Sentence)
			talk.Pattern = append(talk.Pattern, ADNOMINAL)
			if term.Sentence == PREDICATE {
				if lenght > 1 {
					talk.Pattern = append(talk.Pattern, INDIRECT)
				} else {
					talk.Pattern = append(talk.Pattern, DIRECT)
				}
			}
			expressions = append(expressions, talk)
		}
	}
	expressions = append(expressions, captions...)
	var indirect bool = false
	if lenght > 1 {
		indirect = true
	} else {
		for _, caption := range captions[0].Pattern {
			if caption == INDIRECT {
				indirect = true
			}
		}
	}
	var preposition Talk
	preposition.Term = connection.Term
	preposition.Etiology = append(preposition.Etiology, connection.Class)
	preposition.Pattern = append(preposition.Pattern, connection.Sentence)
	preposition.Pattern = append(preposition.Pattern, ADNOMINAL)
	if connection.Sentence == PREDICATE {
		if indirect {
			preposition.Pattern = append(preposition.Pattern, INDIRECT)
		} else {
			preposition.Pattern = append(preposition.Pattern, DIRECT)
		}
	}
	expressions = append(expressions, preposition)
	return expressions
}

func MountAdjective(adjectives []Glossary) []Talk {
	if adjectives == nil {
		return nil
	}
	var adjective Word = adjectives[0].First
	var adverb Word
	var adverb_adverb Word

	var quantity int = len(adjectives)
	var filters []Glossary
	for count := 0; count < quantity; count++ {
		if adjectives[count].Order >= 0 {
			filters = append(filters, adjectives[count])
		}
	}
	sort.Slice(filters, func(i, j int) bool {
		return filters[i].Order < filters[j].Order
	})
	var size int = len(filters)
	for count := 0; count < size; count++ {
		var second Word = filters[count].Second
		if second.Class == ADJECTIVE {
			break
		}
		if second.Class == NOUN {
			break
		}
		if second.Class == VERB {
			break
		}
		if second.Class == ADVERB && adverb.Term != "" {
			adverb_adverb = adjectives[count].Second
			break
		}
		if second.Class == ADVERB {
			adverb = adjectives[count].Second
		}
	}

	var words []Word
	if adverb.Term != "" {
		words = append(words, adverb)
	}
	if adverb_adverb.Term != "" {
		words = append(words, adverb_adverb)
	}
	if adjective.Term != "" {
		words = append(words, adjective)
	}

	var talks []Talk
	for _, word := range words {
		var talk Talk
		talk.Term = word.Term
		talk.Etiology = append(talk.Etiology, word.Class)
		talk.Pattern = append(talk.Pattern, word.Sentence)
		talk.Pattern = append(talk.Pattern, ADVERBIAL)
		talks = append(talks, talk)
	}
	return talks
}

func TypePeriod(talks []Talk) string {
	var quantity int = 0
	for _, verb := range talks {
		for _, word := range verb.Etiology {
			if word == VERB {
				quantity++
			}
		}
	}
	var brand string = ""
	if quantity > 1 {
		brand = COMPOUND
	} else {
		brand = SAMPLE
	}
	return brand
}

func TypeConnection(vocabularies []Vocabulary) Word {
	var connection Word
	for _, vocabulary := range vocabularies {
		for _, jargon := range vocabulary.Jargon {
			if jargon.Order == 0 {
				connection = jargon.First
			}
		}
	}
	return connection
}

func KindConnection(vocabularies []Vocabulary) Word {
	var kind Word
	for _, vocabulary := range vocabularies {
		for _, jargon := range vocabulary.Jargon {
			if jargon.Order == -1 {
				kind = jargon.Second
			}
		}
	}
	return kind
}

func SyntaxSubject(orations []Phrase, dome brand.Arbor, language string, rate int) Recite {
	var lexicons []Glossary = Oration(orations, SUBJECT, rate)

	var nouns []Glossary
	for _, subject := range lexicons {
		if subject.First.Class == NOUN {
			var noun Glossary
			noun.First = subject.First
			noun.Second = subject.Second
			noun.Order = subject.Order
			nouns = append(nouns, noun)
		}
	}
	var prepositions []Glossary
	for _, predicate := range lexicons {
		if predicate.First.Class == PREPOSITION {
			var preposition Glossary
			preposition.First = predicate.First
			preposition.Second = predicate.Second
			preposition.Order = predicate.Order
			prepositions = append(prepositions, preposition)
		}
	}
	var vocabularies []Vocabulary
	vocabularies = MountConnection(lexicons, SUBJECT, dome, language)

	var expressions []Talk
	if vocabularies != nil {
		var kind Word = KindConnection(vocabularies)
		var terms []Word = MountPreposition(kind, prepositions)
		var connection Word = TypeConnection(vocabularies)
		if kind.Class == NUMERAL {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(terms, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == ADJECTIVE {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(terms, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == NOUN {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(nil, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == ADVERB {
			var dictions []Glossary = FilterGlossary(lexicons, kind)
			var adjectives []Talk = MountAdverb(lexicons, dictions)
			var captions []Talk = MountNoun(nouns, prepositions)
			adjectives = append(adjectives, captions...)
			var compounds []Talk = NounCompound(terms, connection, adjectives)
			expressions = append(expressions, compounds...)
		}
	} else {
		var lexicons = MountNoun(nouns, prepositions)
		expressions = append(expressions, lexicons...)
	}

	var brand string = TypeSubject(nouns, vocabularies)
	var clause Recite = Recite{
		Kind: brand,
		Talk: expressions,
	}
	return clause
}

func SyntaxPredicate(orations []Phrase, dome brand.Arbor, language string, rate int) Recite {
	var lexicons []Glossary = Oration(orations, PREDICATE, rate)

	var verbs []Glossary
	for _, predicate := range lexicons {
		if predicate.First.Class == VERB {
			var verb Glossary
			verb.First = predicate.First
			verb.Second = predicate.Second
			verb.Order = predicate.Order
			verbs = append(verbs, verb)
		}
	}
	var nouns []Glossary
	for _, predicate := range lexicons {
		if predicate.First.Class == NOUN {
			var noun Glossary
			noun.First = predicate.First
			noun.Second = predicate.Second
			noun.Order = predicate.Order
			nouns = append(nouns, noun)
		}
	}
	var adjectives []Glossary
	for _, predicate := range lexicons {
		if predicate.First.Class == ADJECTIVE {
			var adjective Glossary
			adjective.First = predicate.First
			adjective.Second = predicate.Second
			adjective.Order = predicate.Order
			adjectives = append(adjectives, adjective)
		}
	}
	var prepositions []Glossary
	for _, predicate := range lexicons {
		if predicate.First.Class == PREPOSITION {
			var preposition Glossary
			preposition.First = predicate.First
			preposition.Second = predicate.Second
			preposition.Order = predicate.Order
			prepositions = append(prepositions, preposition)
		}
	}
	var specials []Glossary
	for _, predicate := range lexicons {
		if TypeEspecial(predicate.First.Term) {
			var special Glossary
			special.First = predicate.First
			special.Second = predicate.Second
			special.Order = predicate.Order
			specials = append(specials, special)
		}
	}
	var vocabularies []Vocabulary
	vocabularies = MountConnection(lexicons, PREDICATE, dome, language)

	var expressions []Talk
	if vocabularies != nil {
		var adverbial_verb []Talk = MountVerb(verbs)
		expressions = append(expressions, adverbial_verb...)
		var kind Word = KindConnection(vocabularies)
		var terms []Word = MountPreposition(kind, prepositions)
		var connection Word = TypeConnection(vocabularies)
		if kind.Class == NUMERAL {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(terms, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == ADJECTIVE {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(terms, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == NOUN {
			var captions []Talk = MountNoun(nouns, prepositions)
			var compounds []Talk = NounCompound(nil, connection, captions)
			expressions = append(expressions, compounds...)
		}
		if kind.Class == ADVERB {
			var dictions []Glossary = FilterGlossary(lexicons, kind)
			var adjectives []Talk = MountAdverb(lexicons, dictions)
			var captions []Talk = MountNoun(nouns, prepositions)
			adjectives = append(adjectives, captions...)
			var compounds []Talk = NounCompound(terms, connection, adjectives)
			expressions = append(expressions, compounds...)
		}
	} else {
		var adverbial_verb []Talk = MountVerb(verbs)
		expressions = append(expressions, adverbial_verb...)
		var adnominal_adjunt []Talk = MountNoun(nouns, prepositions)
		expressions = append(expressions, adnominal_adjunt...)
		var adverbial_adjective []Talk = MountAdjective(adjectives)
		expressions = append(expressions, adverbial_adjective...)
	}

	var brand string = TypePeriod(expressions)
	var clause Recite = Recite{
		Kind: brand,
		Talk: expressions,
	}
	return clause
}

func Syntax(orations []Phrase, dome brand.Arbor, language string) []Recite {
	var rate int = 2

	var subject Recite = SyntaxSubject(orations, dome, language, rate)
	var predicates Recite = SyntaxPredicate(orations, dome, language, rate)

	var phrases []Recite
	phrases = append(phrases, subject)
	phrases = append(phrases, predicates)
	return phrases
}

func Oration(orations []Phrase, sentence string, rate int) []Glossary {
	var words []Word
	for _, oration := range orations {
		for _, word := range oration.Word {
			words = append(words, word)
		}
	}
	var filters []Word
	for _, subject := range words {
		if subject.Sentence == sentence {
			filters = append(filters, subject)
		}
	}

	var locutions []Phrase
	var locution Phrase
	locution.Word = filters
	locutions = append(locutions, locution)
	var glossaries []Glossary = SetGlossary(locutions)
	var lexicons []Glossary = MountGlossary(glossaries, rate)
	return lexicons
}

func Right(glossaries []Glossary, order int, quantity int) []Glossary {
	var lenght int = len(glossaries)
	var sequence int = order + quantity
	if sequence >= lenght {
		sequence = lenght - 1
	}
	var lexicon []Glossary
	var next int = 1
	for count := order; count < sequence; count++ {
		var first Word = glossaries[order].First
		var second Word = glossaries[order+next].Second
		var dictionary Glossary
		dictionary.First = first
		dictionary.Second = second
		dictionary.Order = next
		lexicon = append(lexicon, dictionary)
		next++
	}
	return lexicon
}

func Left(glossaries []Glossary, order int, quantity int) []Glossary {
	var sequence int = order - quantity
	if sequence < 0 {
		sequence = 0
	}
	var lexicon []Glossary
	var next int = 1
	for count := order; count > sequence; count-- {
		var first Word = glossaries[order].First
		var second Word = glossaries[order-next].First
		var dictionary Glossary
		dictionary.First = first
		dictionary.Second = second
		dictionary.Order = next * -1
		lexicon = append(lexicon, dictionary)
		next++
	}
	return lexicon
}

func MountGlossary(glossaries []Glossary, rate int) []Glossary {
	var lexicons []Glossary = glossaries
	var length = len(glossaries)

	for count := 0; count < length; count++ {
		var left []Glossary = Left(glossaries, count, rate)
		for _, item := range left {
			var lexicon Glossary
			lexicon.First = item.First
			lexicon.Second = item.Second
			lexicon.Order = item.Order
			lexicons = append(lexicons, lexicon)
		}
		var right []Glossary = Right(glossaries, count, rate)
		for _, item := range right {
			var lexicon Glossary
			lexicon.First = item.First
			lexicon.Second = item.Second
			lexicon.Order = item.Order
			lexicons = append(lexicons, lexicon)
		}
	}
	return lexicons
}

func SetGlossary(word []Phrase) []Glossary {
	var embendding []Glossary

	var count int = 0
	var length_phrase = len(word)
	for count < length_phrase {
		var spell Phrase = word[count]
		var length_word = len(spell.Word)
		var caption []Word = spell.Word
		for index := range spell.Word {
			if index+1 == length_word {
				var locution Glossary
				locution.First = caption[index]
				embendding = append(embendding, locution)
			} else {
				var locution Glossary
				locution.First = caption[index]
				locution.Second = caption[index+1]
				embendding = append(embendding, locution)
			}
		}
		count++
	}
	return embendding
}
