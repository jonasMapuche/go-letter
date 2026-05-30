package grammar

import (
	"regexp"
	"slices"
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
	First  Talk
	Second Talk
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
	BOTH          = "ambos"
	COMMA         = ","
	COLON         = ":"
	SEMICOLON     = ";"
	SPOT          = "."
	APPOSITIVE    = "aposto"
	ENUMERATIVE   = "enumerativo"
	CONNECTIVE    = "conectivo"
	RATE          = 2
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
	var words []string = strings.Split(errand, " ")
	var vocables []string
	for _, term := range words {
		if term != "" {
			vocables = append(vocables, term)
		}
	}
	if language == ENGLISH {
		vocables = SplitEnglish(vocables, arbor, language)
	}
	var units []Word
	var phrases []Phrase
	var predicates bool = false
	for _, term := range vocables {
		var spell Word
		spell.Term = term
		spell.Class = ""
		spell.Sentence = ""
		units = append(units, spell)
		if GetSpesh(spell.Term) {
			if GetBefore(spell.Term) {
				continue
			}
			var locution []Word = Slash(units, arbor, language, predicates)
			var kind string = TypePhrase(spell.Term)
			var clause Phrase = Phrase{
				Kind: kind,
				Word: locution,
			}
			phrases = append(phrases, clause)
			units = nil
			if TypeEspecial(spell.Term) {
				predicates = true
			}
		}
	}
	var expressions []Phrase = ChangeAnbuguity(phrases, arbor, language)
	return expressions
}

func ChangeAnbuguity(phrases []Phrase, arbor Arbor, language string) []Phrase {
	var rate int = RATE
	var predicates []Glossary = Oration(phrases, PREDICATE, rate)
	var verbs []Glossary
	for _, predicate := range predicates {
		for _, etiology := range predicate.First.Etiology {
			if etiology == VERB {
				var verb Glossary = predicate
				verbs = append(verbs, verb)
			}
		}
	}
	var actions []Talk
	for _, verb := range verbs {
		for _, noun := range arbor.Noun {
			var term string = SplitNoun(noun.Name)
			if verb.First.Term == term && noun.Language == language && verb.Order == -1 {
				var word Talk = verb.Second
				for _, etiology := range word.Etiology {
					if etiology == ADJECTIVE || etiology == PRONOUN || etiology == NUMERAL {
						var talk Talk
						talk.Term = verb.First.Term
						talk.Etiology = append(talk.Etiology, NOUN)
						talk.Pattern = verb.First.Pattern
						talk.Order = verb.First.Order
						actions = append(actions, talk)
					}
				}
			}
		}
	}
	var expressions []Phrase
	var ratio int = 1
	for _, phrase := range phrases {
		var terms []Word
		for mark, word := range phrase.Word {
			terms = append(terms, word)
			for _, action := range actions {
				if action.Order == ratio {
					var discussion Word
					discussion.Term = action.Term
					discussion.Class = action.Etiology[0]
					discussion.Sentence = action.Pattern[0]
					terms[mark] = discussion
					break
				}
			}
			ratio++
		}
		var diction Phrase = Phrase{
			Kind: phrase.Kind,
			Word: terms,
		}
		expressions = append(expressions, diction)
	}
	return expressions
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
			if language == ENGLISH {
				preposition = PrepositionVerb(spell.Term)
				if preposition {
					term = SplitVerb(spell.Term)
				}
			}
			if value.Name == strings.ToLower(term) && value.Language == language && spell.Sentence == "" {
				spell.Class = VERB
				spell.Sentence = PREDICATE
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

	var class string = TypePhrase(message)

	locution := Phrase{
		Kind: class,
		Word: unit,
	}

	return locution
}

func TypePhrase(message string) string {
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

func OrderFirst(glossaries []Glossary) int {
	var minus int = 0
	for index, noun := range glossaries {
		if index == 0 {
			minus = noun.First.Order
		}
		var num = noun.First.Order
		if num < minus {
			minus = num
		}
	}
	return minus
}

func OrderLast(glossaries []Glossary) int {
	var maximus int = 0
	for _, noun := range glossaries {
		var max = noun.First.Order
		if max > maximus {
			maximus = max
		}
	}
	return maximus
}

func FilterGlossary(glossaries []Glossary, word Talk) []Glossary {
	var lexicons []Glossary
	for _, glossary := range glossaries {
		if slices.Equal(glossary.First.Etiology, word.Etiology) &&
			slices.Equal(glossary.First.Pattern, word.Pattern) &&
			glossary.First.Term == word.Term && glossary.First.Order == word.Order {
			var lexicon Glossary
			lexicon.First = glossary.First
			lexicon.Second = glossary.Second
			lexicon.Order = glossary.Order
			lexicons = append(lexicons, lexicon)
		}
	}
	return lexicons
}

func MountVocabulary(glossaries []Glossary, crescent bool, rear bool) []Vocabulary {
	if glossaries == nil {
		return nil
	}
	sort.Slice(glossaries, func(i, j int) bool {
		if glossaries[i].First.Term != glossaries[j].First.Term {
			return glossaries[i].First.Term < glossaries[j].First.Term
		}
		return glossaries[i].First.Order < glossaries[j].First.Order
	})
	var vocabularies []Vocabulary
	var before string = ""
	var order int = 0
	var vocabulary Vocabulary
	for _, association := range glossaries {
		if !(rear && association.Order <= 0) {
			if !(!rear && association.Order >= 0) {
				continue
			}
		}
		if before == "" {
			vocabulary.Jargon = append(vocabulary.Jargon, association)
			before = association.First.Term
			order = association.First.Order
		} else {
			if !(before == association.First.Term && order == association.First.Order) {
				if crescent {
					sort.Slice(vocabulary.Jargon, func(i, j int) bool {
						if vocabulary.Jargon[i].First.Term != vocabulary.Jargon[j].First.Term {
							return vocabulary.Jargon[i].First.Term < vocabulary.Jargon[j].First.Term
						}
						return vocabulary.Jargon[i].First.Order < vocabulary.Jargon[j].First.Order
					})
				} else {
					sort.Slice(vocabulary.Jargon, func(i, j int) bool {
						if vocabulary.Jargon[i].First.Term != vocabulary.Jargon[j].First.Term {
							return vocabulary.Jargon[i].First.Term > vocabulary.Jargon[j].First.Term
						}
						return vocabulary.Jargon[i].First.Order > vocabulary.Jargon[j].First.Order
					})
				}
				vocabularies = append(vocabularies, vocabulary)
				vocabulary.Jargon = nil
				vocabulary.Jargon = append(vocabulary.Jargon, association)
				before = association.First.Term
				order = association.First.Order
			} else {
				vocabulary.Jargon = append(vocabulary.Jargon, association)
			}
		}
	}
	if vocabulary.Jargon != nil {
		if crescent {
			sort.Slice(vocabulary.Jargon, func(i, j int) bool {
				if vocabulary.Jargon[i].First.Term != vocabulary.Jargon[j].First.Term {
					return vocabulary.Jargon[i].First.Term < vocabulary.Jargon[j].First.Term
				}
				return vocabulary.Jargon[i].First.Order < vocabulary.Jargon[j].First.Order
			})
		} else {
			sort.Slice(vocabulary.Jargon, func(i, j int) bool {
				if vocabulary.Jargon[i].First.Term != vocabulary.Jargon[j].First.Term {
					return vocabulary.Jargon[i].First.Term > vocabulary.Jargon[j].First.Term
				}
				return vocabulary.Jargon[i].First.Order > vocabulary.Jargon[j].First.Order
			})
		}
		vocabularies = append(vocabularies, vocabulary)
	}
	return vocabularies
}

func MountPreposition(noun Talk, prepositions []Glossary) []Talk {
	var link Talk
	for _, preposition := range prepositions {
		if preposition.Order == 0 {
			if CompareTalk(preposition.Second, noun) {
				link = preposition.First
			}
		}
	}
	var terms []Talk
	if link.Term != "" {
		terms = append(terms, link)
		terms = append(terms, noun)
	} else {
		terms = append(terms, noun)
	}
	return terms
}

func TypeEspecial(value string) bool {
	if value == COMMA || value == SEMICOLON || value == COLON {
		return true
	}
	return false
}

func TypeConnection(vocabularies []Vocabulary) Talk {
	var connection Talk
	for _, vocabulary := range vocabularies {
		for _, jargon := range vocabulary.Jargon {
			if jargon.Order == 0 {
				connection = jargon.First
			}
		}
	}
	return connection
}

func TypeCore(talk Talk) string {
	var kind string = NOUN
	for _, brand := range talk.Etiology {
		if brand == PRONOUN {
			kind = PRONOUN
		}
	}
	return kind
}

func KindConnection(vocabularies []Vocabulary) Talk {
	var kind Talk
	for _, vocabulary := range vocabularies {
		for _, jargon := range vocabulary.Jargon {
			if jargon.Order == -1 {
				kind = jargon.Second
			}
		}
	}
	return kind
}

func TypePronoun(pronouns []Glossary) string {
	var pronoun string = ""
	var kinds []string
	var vocabularies []Vocabulary = MountVocabulary(pronouns, true, false)
	for _, vocabulary := range vocabularies {
		var lexicons []Glossary = vocabulary.Jargon
		for _, lexicon := range lexicons {
			if lexicon.Order == 0 {
				var brand bool = false
				for _, etiology := range lexicon.Second.Etiology {
					if etiology == NOUN {
						kinds = append(kinds, ADJECTIVE)
						brand = true
					}
				}
				if !brand {
					kinds = append(kinds, PRONOUN)
				}
			}
		}
	}
	var quantity int = 0
	for _, brand := range kinds {
		if brand == ADJECTIVE {
			quantity++
		}
		if brand == PRONOUN {
			quantity++
		}
	}
	if quantity > 1 {
		pronoun = BOTH
	} else {
		pronoun = kinds[0]
	}
	return pronoun
}

func SplitPronoun(glossaries []Glossary) []Glossary {
	var pronouns []Glossary
	var vocabularies []Vocabulary = MountVocabulary(glossaries, true, false)
	for _, vocabulary := range vocabularies {
		var lexicons []Glossary = vocabulary.Jargon
		var kind string = TypePronoun(lexicons)
		if kind == PRONOUN {
			pronouns = append(pronouns, lexicons...)
		}
	}
	return pronouns
}

func TypeAdjective(pronouns []Glossary) string {
	var pronoun string = ""
	var kinds []string
	var vocabularies []Vocabulary = MountVocabulary(pronouns, true, true)
	for _, vocabulary := range vocabularies {
		var lexicons []Glossary = vocabulary.Jargon
		for _, lexicon := range lexicons {
			if lexicon.Order == 0 {
				var brand bool = false
				for _, etiology := range lexicon.Second.Etiology {
					if etiology == NOUN || etiology == CONJUNCTION {
						kinds = append(kinds, ADJECTIVE)
						brand = true
					}
				}
				if !brand {
					kinds = append(kinds, ADVERBIAL)
				}
			}
		}
	}
	var quantity int = 0
	for _, brand := range kinds {
		if brand == ADJECTIVE {
			quantity++
		}
		if brand == ADVERBIAL {
			quantity++
		}
	}
	if quantity > 1 {
		pronoun = BOTH
	} else {
		pronoun = kinds[0]
	}
	return pronoun
}

func SplitAdjective(glossaries []Glossary) []Glossary {
	var adjectives []Glossary
	var vocabularies []Vocabulary = MountVocabulary(glossaries, true, true)
	for _, vocabulary := range vocabularies {
		var lexicons []Glossary = vocabulary.Jargon
		var kind = TypeAdjective(lexicons)
		if kind == ADVERBIAL {
			adjectives = append(adjectives, lexicons...)
		}
	}
	return adjectives
}

func CompareTalk(first Talk, second Talk) bool {
	var compare bool = false
	if slices.Equal(first.Etiology, second.Etiology) &&
		slices.Equal(first.Pattern, second.Pattern) &&
		first.Term == second.Term &&
		first.Order == second.Order {
		return true
	}
	return compare
}

func MountCompound(noun Talk, prepositions []Glossary, integrant string) Recite {
	var recipe Recite
	var talk Talk
	talk.Term = noun.Term
	talk.Etiology = noun.Etiology
	talk.Pattern = append(noun.Pattern, ADNOMINAL)
	talk.Order = noun.Order
	var link Talk
	for _, preposition := range prepositions {
		if preposition.Order == 0 {
			if CompareTalk(preposition.Second, talk) {
				link = preposition.First
			}
		}
	}
	var terms []Talk
	if link.Term != "" && (integrant == INDIRECT || integrant == BOTH) {
		terms = append(terms, link)
		terms = append(terms, talk)
		recipe.Kind = INDIRECT
		recipe.Talk = terms
	}
	if link.Term == "" && (integrant == DIRECT || integrant == BOTH) {
		terms = append(terms, talk)
		recipe.Kind = DIRECT
		recipe.Talk = terms
	}
	return recipe
}

func MountNoun(nouns []Glossary, prepositions []Glossary, integrant string) []Recite {
	if nouns == nil {
		return nil
	}
	var recipes []Recite
	var vocabularies = MountVocabulary(nouns, true, true)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var noun Talk = glossaries[0].First
		var article Talk
		var pronome Talk
		var numeral Talk
		var adverb Talk
		var adverb_adverb Talk
		var adjetivo Talk
		var preposition Talk
		var size int = len(glossaries)
		var exit bool = false
		for count := 0; count < size; count++ {
			var second Talk = glossaries[count].Second
			var order int = glossaries[count].Order
			var terms []Talk = MountPreposition(second, prepositions)
			var lenght = len(terms)
			if lenght > 1 {
				preposition = terms[0]
			}
			var kind string = TypeCore(noun)
			if kind == NOUN {
				for _, vocable := range second.Etiology {
					if vocable == PREPOSITION && order != 0 {
						exit = true
						break
					}
					if vocable == NOUN && order != 0 {
						exit = true
						break
					}
					if vocable == VERB && order != 0 {
						exit = true
						break
					}
					if vocable == CONJUNCTION && order != 0 {
						exit = true
						break
					}
					if vocable == ARTICLE && order != 0 {
						article = second
						exit = true
						break
					}
					if vocable == ADVERB && order != 0 {
						adverb = second
					}
					if vocable == ADVERB && adverb.Term != "" && order != 0 {
						adverb_adverb = second
					}
					if vocable == ADJECTIVE && order != 0 {
						adjetivo = second
					}
					if vocable == NUMERAL && order != 0 {
						numeral = second
					}
					if vocable == PRONOUN && order != 0 {
						pronome = second
					}
				}
				if exit {
					break
				}
			}
		}
		if preposition.Term == "" {
			var vocables []Talk = MountPreposition(noun, prepositions)
			var lenght = len(vocables)
			if lenght > 1 {
				preposition = vocables[0]
			}
		}
		var words []Talk
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
		if integrant == BOTH ||
			(integrant == DIRECT && preposition.Term == "") ||
			(integrant == INDIRECT && preposition.Term != "") {
			var talks []Talk
			for _, word := range words {
				var talk Talk
				talk.Term = word.Term
				talk.Etiology = append(talk.Etiology, word.Etiology...)
				talk.Pattern = append(talk.Pattern, word.Pattern...)
				talk.Pattern = append(talk.Pattern, ADNOMINAL)
				talk.Order = word.Order
				talks = append(talks, talk)
			}
			var recipe = Recite{
				Kind: integrant,
				Talk: talks,
			}
			recipes = append(recipes, recipe)
		}
	}
	return recipes
}

func MountVerb(verbs []Glossary) []Recite {
	if verbs == nil {
		return nil
	}
	var recites []Recite
	var vocabularies = MountVocabulary(verbs, true, false)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var verb Talk = glossaries[0].First
		var adverb Talk
		var adverb_adverb Talk
		var size int = len(glossaries)
		var exit bool = false
		for count := 0; count < size; count++ {
			var second Talk = glossaries[count].Second
			var order int = glossaries[count].Order
			for _, vocable := range second.Etiology {
				if vocable == NOUN {
					exit = true
					break
				}
				if vocable == NUMERAL {
					exit = true
					break
				}
				if vocable == ADJECTIVE {
					exit = true
					break
				}
				if vocable == VERB && order != 0 {
					exit = true
					break
				}
				if vocable == ADVERB && adverb.Term != "" {
					adverb_adverb = second
					exit = true
					break
				}
				if vocable == ADVERB {
					adverb = second
				}
			}
			if exit {
				break
			}
		}
		var words []Talk
		if adverb.Term != "" {
			words = append(words, adverb)
		}
		if adverb_adverb.Term != "" {
			words = append(words, adverb_adverb)
		}
		if verb.Term != "" {
			words = append(words, verb)
		}
		var talks []Talk
		for _, word := range words {
			var talk Talk
			talk.Term = word.Term
			talk.Etiology = append(talk.Etiology, word.Etiology...)
			talk.Pattern = append(talk.Pattern, word.Pattern...)
			talk.Pattern = append(talk.Pattern, ADVERBIAL)
			talk.Order = word.Order
			talks = append(talks, talk)
		}
		var recipe Recite = Recite{
			Kind: VERB,
			Talk: talks,
		}
		recites = append(recites, recipe)
	}
	return recites
}

func MountAdjective(adjectives []Glossary) []Recite {
	if adjectives == nil {
		return nil
	}
	var recipes []Recite
	var vocabularies = MountVocabulary(adjectives, true, true)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var adjective Talk = glossaries[0].First
		var adverb Talk
		var adverb_adverb Talk
		var size int = len(glossaries)
		var exit bool = false
		for count := 0; count < size; count++ {
			var second Talk = glossaries[count].Second
			for _, vocable := range second.Etiology {
				if vocable == ADJECTIVE {
					exit = true
					break
				}
				if vocable == NOUN {
					exit = true
					break
				}
				if vocable == VERB {
					exit = true
					break
				}
				if vocable == ADVERB && adverb.Term != "" {
					adverb_adverb = second
					exit = true
					break
				}
				if vocable == ADVERB {
					adverb = second
				}
			}
			if exit {
				break
			}
		}
		var words []Talk
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
			talk.Etiology = append(talk.Etiology, word.Etiology...)
			talk.Pattern = append(talk.Pattern, word.Pattern...)
			talk.Pattern = append(talk.Pattern, ADVERBIAL)
			talk.Order = word.Order
			talks = append(talks, talk)
		}
		var recipe = Recite{
			Kind: PREDICATIVE,
			Talk: talks,
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func MountAdverb(glossaries []Glossary, adverbs []Glossary) []Recite {
	if glossaries == nil {
		return nil
	}
	var filter Talk
	for _, adverb := range adverbs {
		for _, term := range adverb.Second.Etiology {
			if term == ADJECTIVE {
				filter = adverb.Second
				break
			}
		}
	}
	var adjectives []Glossary = FilterGlossary(glossaries, filter)
	var recipes []Recite
	var vocabularies = MountVocabulary(adjectives, true, false)
	for _, vocabulary := range vocabularies {
		var glossaries []Glossary = vocabulary.Jargon
		var adjective Talk = glossaries[0].First
		var adverb Talk
		var adverb_adverb Talk
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
		var exit bool = false
		for count := 0; count < size; count++ {
			var second Talk = filters[count].Second
			for _, vocable := range second.Etiology {
				if vocable == ADJECTIVE {
					exit = true
					break
				}
				if vocable == NOUN {
					exit = true
					break
				}
				if vocable == VERB {
					exit = true
					break
				}
				if vocable == NUMERAL {
					exit = true
					break
				}
				if vocable == ARTICLE {
					exit = true
					break
				}
				if vocable == ADVERB && adverb.Term != "" {
					adverb_adverb = second
					exit = true
					break
				}
				if vocable == ADVERB {
					adverb = second
				}
			}
			if exit {
				break
			}
		}
		var words []Talk
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
			talk.Etiology = append(talk.Etiology, word.Etiology...)
			talk.Pattern = append(talk.Pattern, word.Pattern...)
			talk.Pattern = append(talk.Pattern, ADVERBIAL)
			talk.Order = word.Order
			talks = append(talks, talk)
		}
		var recipe = Recite{
			Kind: DIRECT,
			Talk: talks,
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

func MountConnection(connections []Glossary, syntax string, dome brand.Arbor, language string) []Vocabulary {
	var associations []Glossary
	for _, association := range connections {
		if syntax == SUBJECT {
			for _, term := range association.First.Etiology {
				if term == CONJUNCTION {
					var lexicon Glossary
					lexicon.First = association.First
					lexicon.Second = association.Second
					lexicon.Order = association.Order
					associations = append(associations, lexicon)
				}
			}
		} else {
			for _, term := range association.First.Etiology {
				if term == CONJUNCTION || TypeEspecial(association.First.Term) {
					var lexicon Glossary
					lexicon.First = association.First
					lexicon.Second = association.Second
					lexicon.Order = association.Order
					associations = append(associations, lexicon)
				}
			}
			for _, term := range association.First.Etiology {
				if term == PRONOUN {
					for _, value := range dome.Pronoun {
						if value.Name == strings.ToLower(association.First.Term) && value.Type == RELATIVE &&
							value.Language == language {
							var first Talk
							first.Pattern = association.First.Pattern
							first.Etiology = association.First.Etiology
							first.Etiology = append(first.Etiology, RELATIVE)
							first.Term = association.First.Term
							first.Order = association.First.Order
							var lexicon Glossary
							lexicon.First = first
							lexicon.Second = association.Second
							lexicon.Order = association.Order
							associations = append(associations, lexicon)
							break
						}
						if value.Name == strings.ToLower(association.First.Term) && value.Type == UNKNOW &&
							value.Language == language {
							var first Talk
							first.Pattern = association.First.Pattern
							first.Etiology = association.First.Etiology
							first.Etiology = append(first.Etiology, UNKNOW)
							first.Term = association.First.Term
							first.Order = association.First.Order
							var lexicon Glossary
							lexicon.First = first
							lexicon.Second = association.Second
							lexicon.Order = association.Order
							associations = append(associations, lexicon)
							break
						}
					}
				}
			}
		}
	}
	var vocabularies []Vocabulary = MountVocabulary(associations, true, true)
	if vocabularies == nil {
		return nil
	}
	if syntax == SUBJECT {
		return vocabularies
	}
	if syntax == PREDICATE {
		var specials []Glossary
		for _, connection := range connections {
			if TypeEspecial(connection.First.Term) {
				var special Glossary
				special.First = connection.First
				special.Second = connection.Second
				special.Order = connection.Order
				specials = append(specials, special)
			}
		}
		var prepositions []Glossary
		for _, connection := range connections {
			for _, term := range connection.First.Etiology {
				if term == PREPOSITION {
					var preposition Glossary
					preposition.First = connection.First
					preposition.Second = connection.Second
					preposition.Order = connection.Order
					prepositions = append(prepositions, preposition)
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
				for _, term := range glossary.First.Etiology {
					if term == CONJUNCTION || term == UNKNOW || term == RELATIVE {
						for _, preposition := range prepositions {
							if glossary.Order == sequence && CompareTalk(glossary.Second, preposition.First) {
								var filters = FilterGlossary(connections, preposition.First)
								var dependences []Vocabulary = MountVocabulary(filters, true, true)
								var bondages []Glossary = dependences[0].Jargon
								dictionary.Jargon = append(dictionary.Jargon, bondages...)
								contact = true
								break
							}
						}
					}
				}
			}
			for _, glossary := range vocabulary.Jargon {
				for _, term := range glossary.First.Etiology {
					if term == CONJUNCTION {
						if contact == true {
							sequence--
						}
						for _, special := range specials {
							if glossary.Order == sequence && CompareTalk(glossary.Second, special.First) {
								var filters = FilterGlossary(connections, special.First)
								var dependences []Vocabulary = MountVocabulary(filters, true, true)
								var bondages []Glossary = dependences[0].Jargon
								dictionary.Jargon = append(dictionary.Jargon, bondages...)
								break
							}
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

func NounCompound(connection Talk, captions []Recite, kind string) []Recite {
	var recipes []Recite
	var expressions []Talk
	for _, caption := range captions {
		expressions = append(expressions, caption.Talk...)
	}
	var preposition Talk
	preposition.Term = connection.Term
	preposition.Etiology = append(preposition.Etiology, connection.Etiology...)
	preposition.Pattern = append(preposition.Pattern, connection.Pattern...)
	preposition.Pattern = append(preposition.Pattern, ADNOMINAL)
	preposition.Order = connection.Order
	expressions = append(expressions, preposition)
	sort.Slice(expressions, func(i, j int) bool {
		return expressions[i].Order < expressions[j].Order
	})
	var recipe = Recite{
		Kind: kind,
		Talk: expressions,
	}
	recipes = append(recipes, recipe)
	return recipes
}

func PeriodSample(glossaries []Glossary) []Recite {
	var recites []Recite
	var nouns []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == NOUN {
				var noun Glossary
				noun.First = predicate.First
				noun.Second = predicate.Second
				noun.Order = predicate.Order
				nouns = append(nouns, noun)
			}
		}
	}
	var pronouns []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == PRONOUN {
				var pronoun Glossary
				pronoun.First = predicate.First
				pronoun.Second = predicate.Second
				pronoun.Order = predicate.Order
				pronouns = append(pronouns, pronoun)
			}
		}
	}
	var adjectives []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == ADJECTIVE {
				var adjective Glossary
				adjective.First = predicate.First
				adjective.Second = predicate.Second
				adjective.Order = predicate.Order
				adjectives = append(adjectives, adjective)
			}
		}
	}
	var prepositions []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == PREPOSITION {
				var preposition Glossary
				preposition.First = predicate.First
				preposition.Second = predicate.Second
				preposition.Order = predicate.Order
				prepositions = append(prepositions, preposition)
			}
		}
	}
	var adverbials []Glossary = SplitAdjective(adjectives)
	var qualities []Recite = MountAdjective(adverbials)
	recites = append(recites, qualities...)
	var surrogates []Glossary = SplitPronoun(pronouns)
	nouns = append(nouns, surrogates...)
	var directs []Recite = MountNoun(nouns, prepositions, DIRECT)
	var indirects []Recite = MountNoun(nouns, prepositions, INDIRECT)
	recites = append(recites, directs...)
	recites = append(recites, indirects...)
	return recites
}

func PeriodCompound(glossaries []Glossary, relations []Vocabulary) []Recite {
	var recites []Recite
	var nouns []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == NOUN {
				var noun Glossary
				noun.First = predicate.First
				noun.Second = predicate.Second
				noun.Order = predicate.Order
				nouns = append(nouns, noun)
			}
		}
	}
	var pronouns []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == PRONOUN {
				var pronoun Glossary
				pronoun.First = predicate.First
				pronoun.Second = predicate.Second
				pronoun.Order = predicate.Order
				pronouns = append(pronouns, pronoun)
			}
		}
	}
	var adjectives []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == ADJECTIVE {
				var adjective Glossary
				adjective.First = predicate.First
				adjective.Second = predicate.Second
				adjective.Order = predicate.Order
				adjectives = append(adjectives, adjective)
			}
		}
	}
	var prepositions []Glossary
	for _, predicate := range glossaries {
		for _, term := range predicate.First.Etiology {
			if term == PREPOSITION {
				var preposition Glossary
				preposition.First = predicate.First
				preposition.Second = predicate.Second
				preposition.Order = predicate.Order
				prepositions = append(prepositions, preposition)
			}
		}
	}
	var adverbials []Glossary = SplitAdjective(adjectives)
	var qualities []Recite = MountAdjective(adverbials)
	recites = append(recites, qualities...)
	var kind Talk = KindConnection(relations)
	var connection Talk = TypeConnection(relations)
	for _, vocable := range kind.Etiology {
		if vocable == NUMERAL || vocable == ADJECTIVE || vocable == PRONOUN {
			var directs []Recite = MountNoun(nouns, prepositions, DIRECT)
			var indirects []Recite = MountNoun(nouns, prepositions, INDIRECT)
			var compounds []Recite
			if directs != nil {
				var first Recite = MountCompound(kind, prepositions, DIRECT)
				directs = append(directs, first)
				compounds = NounCompound(connection, directs, DIRECT)
			} else {
				var first Recite = MountCompound(kind, prepositions, INDIRECT)
				indirects = append(indirects, first)
				compounds = NounCompound(connection, indirects, INDIRECT)
			}
			recites = append(recites, compounds...)
		}
		if vocable == NOUN {
			var surrogates []Glossary = SplitPronoun(pronouns)
			nouns = append(nouns, surrogates...)
			var directs []Recite = MountNoun(nouns, prepositions, DIRECT)
			var indirects []Recite = MountNoun(nouns, prepositions, INDIRECT)
			var compounds []Recite
			if directs != nil {
				compounds = NounCompound(connection, directs, DIRECT)
			} else {
				compounds = NounCompound(connection, indirects, INDIRECT)
			}
			recites = append(recites, compounds...)
		}
	}
	return recites
}

func NextPeriod(relations []Vocabulary, ratio int) Talk {
	var talk Talk
	for index, relation := range relations {
		if index == ratio {
			var connections []Glossary = relation.Jargon
			var vocabularies []Vocabulary = MountVocabulary(connections, true, false)
			for indicator, vocabulary := range vocabularies {
				if indicator == 0 {
					talk = vocabulary.Jargon[0].First
				}
			}
		}
	}
	return talk
}

func MountRelation(vocabularies []Vocabulary) []Recite {
	var recites []Recite
	for _, relation := range vocabularies {
		var connections []Glossary = relation.Jargon
		var talks []Talk
		for _, connection := range connections {
			if connection.Order == 0 {
				talks = append(talks, connection.First)
			}
		}
		sort.Slice(talks, func(i, j int) bool {
			return talks[i].Order < talks[j].Order
		})
		var recite = Recite{
			Kind: CONNECTIVE,
			Talk: talks,
		}
		recites = append(recites, recite)
	}
	return recites
}

func MountPeriod(glossaries []Glossary, relations []Vocabulary) []Recite {
	var recites []Recite
	var connectives = MountRelation(relations)
	recites = append(recites, connectives...)
	for index, relation := range relations {
		var connections []Glossary = relation.Jargon
		if index == 0 {
			var first int = OrderFirst(connections)
			var lexicons []Glossary
			for _, glossary := range glossaries {
				if glossary.First.Order < first {
					lexicons = append(lexicons, glossary)
				}
			}
			var declaims []Recite = PeriodSample(lexicons)
			recites = append(recites, declaims...)
		}
		var initial int = OrderLast(connections)
		var last Talk
		var size int = len(relations)
		if size > index+1 {
			last = NextPeriod(relations, index+1)
		}
		var latest int = 0
		if last.Term != "" {
			latest = last.Order
		}
		var vocabulary []Glossary
		if latest == 0 {
			for _, glossary := range glossaries {
				if glossary.First.Order > initial {
					vocabulary = append(vocabulary, glossary)
				}
			}
		} else {
			for _, glossary := range glossaries {
				if glossary.First.Order > initial && glossary.First.Order < latest {
					vocabulary = append(vocabulary, glossary)
				}
			}
		}
		var declaims []Recite = PeriodSample(vocabulary)
		recites = append(recites, declaims...)
	}
	return recites
}

func SyntaxSubject(orations []Phrase, dome brand.Arbor, language string, rate int) Recite {
	var lexicons []Glossary = Oration(orations, SUBJECT, rate)
	var nouns []Glossary
	for _, subject := range lexicons {
		for _, term := range subject.First.Etiology {
			if term == NOUN {
				var noun Glossary
				noun.First = subject.First
				noun.Second = subject.Second
				noun.Order = subject.Order
				nouns = append(nouns, noun)
			}
		}
	}
	var pronouns []Glossary
	for _, subject := range lexicons {
		for _, term := range subject.First.Etiology {
			if term == PRONOUN {
				var pronoun Glossary
				pronoun.First = subject.First
				pronoun.Second = subject.Second
				pronoun.Order = subject.Order
				pronouns = append(pronouns, pronoun)
			}
		}
	}
	var prepositions []Glossary
	for _, subject := range lexicons {
		for _, term := range subject.First.Etiology {
			if term == PREPOSITION {
				var preposition Glossary
				preposition.First = subject.First
				preposition.Second = subject.Second
				preposition.Order = subject.Order
				prepositions = append(prepositions, preposition)
			}
		}
	}
	var relations []Vocabulary
	relations = MountConnection(lexicons, SUBJECT, dome, language)
	var recite Recite
	if relations != nil {
		var kind Talk = KindConnection(relations)
		var connection Talk = TypeConnection(relations)
		for _, vocable := range kind.Etiology {
			if vocable == NUMERAL || vocable == ADJECTIVE || vocable == PRONOUN {
				var first Recite = MountCompound(kind, prepositions, BOTH)
				var boths []Recite = MountNoun(nouns, prepositions, BOTH)
				boths = append(boths, first)
				var compounds []Recite = NounCompound(connection, boths, COMPOUND)
				recite = compounds[0]
			}
			if vocable == NOUN {
				var surrogates []Glossary = SplitPronoun(pronouns)
				nouns = append(nouns, surrogates...)
				var boths []Recite = MountNoun(nouns, prepositions, BOTH)
				var compounds []Recite = NounCompound(connection, boths, COMPOUND)
				recite = compounds[0]
			}
		}
	} else {
		var surrogates []Glossary = SplitPronoun(pronouns)
		nouns = append(nouns, surrogates...)
		var boths []Recite = MountNoun(nouns, prepositions, BOTH)
		if boths == nil {
			recite.Kind = MISSING
		} else {
			for index, both := range boths {
				if index == 0 {
					recite.Kind = SAMPLE
					recite.Talk = both.Talk
				}
			}
		}
	}
	return recite
}

func TypeAppositive(glossaries []Glossary) bool {
	var appositive bool = false
	for _, glossary := range glossaries {
		if glossary.First.Term == COLON {
			appositive = true
		}
	}
	return appositive
}

func MountAppositive(glossaries []Glossary) []Recite {
	var recites []Recite
	var order int = 0
	for _, glossary := range glossaries {
		if glossary.First.Term == COLON {
			order = glossary.First.Order
		}
	}
	var predicates []Glossary
	var appositives []Glossary
	for _, glossary := range glossaries {
		if glossary.First.Order < order {
			predicates = append(predicates, glossary)
		}
		if glossary.First.Order > order {
			appositives = append(appositives, glossary)
		}
	}
	var captions []Recite = PeriodSample(predicates)
	recites = append(recites, captions...)
	var conjunctions []Glossary
	for _, predicate := range appositives {
		for _, term := range predicate.First.Etiology {
			if term == CONJUNCTION {
				var conjunction Glossary
				conjunction.First = predicate.First
				conjunction.Second = predicate.Second
				conjunction.Order = predicate.Order
				conjunctions = append(conjunctions, conjunction)
			}
		}
	}
	var specials []Glossary
	for _, predicate := range appositives {
		if predicate.First.Term == COMMA || predicate.First.Term == SPOT {
			var special Glossary
			special.First = predicate.First
			special.Second = predicate.Second
			special.Order = predicate.Order
			specials = append(specials, special)
		}
	}
	var connections []Glossary
	connections = append(connections, conjunctions...)
	connections = append(connections, specials...)
	var talks []Talk
	var contact Talk = conjunctions[0].First
	for index, appositive := range appositives {
		if appositive.Order == 0 {
			var link bool = false
			if index == 0 {
				var special Talk
				special.Term = contact.Term
				special.Etiology = append(special.Etiology, contact.Etiology...)
				special.Pattern = append(special.Pattern, contact.Pattern...)
				special.Pattern = append(special.Pattern, APPOSITIVE)
				special.Order = order
				talks = append(talks, special)
				var talk Talk
				talk.Term = appositive.First.Term
				talk.Etiology = append(talk.Etiology, appositive.First.Etiology...)
				talk.Pattern = append(talk.Pattern, appositive.First.Pattern...)
				talk.Pattern = append(talk.Pattern, APPOSITIVE)
				talk.Order = appositive.First.Order
				talks = append(talks, talk)
				continue
			}
			for _, connection := range connections {
				if CompareTalk(appositive.First, connection.First) {
					var recipe = Recite{
						Kind: ENUMERATIVE,
						Talk: talks,
					}
					recites = append(recites, recipe)
					var talk Talk
					if appositive.First.Term == COMMA {
						talk.Term = conjunctions[0].First.Term
					} else {
						talk.Term = appositive.First.Term
					}
					talk.Etiology = append(talk.Etiology, appositive.First.Etiology...)
					talk.Pattern = append(talk.Pattern, appositive.First.Pattern...)
					talk.Pattern = append(talk.Pattern, APPOSITIVE)
					talk.Order = appositive.First.Order
					talks = nil
					talks = append(talks, talk)
					link = true
					break
				}
			}
			if link {
				continue
			}
			var talk Talk
			talk.Term = appositive.First.Term
			talk.Etiology = append(talk.Etiology, appositive.First.Etiology...)
			talk.Pattern = append(talk.Pattern, appositive.First.Pattern...)
			talk.Pattern = append(talk.Pattern, APPOSITIVE)
			talk.Order = appositive.First.Order
			talks = append(talks, talk)
		}
	}
	return recites
}

func SyntaxPredicate(orations []Phrase, dome brand.Arbor, language string, rate int) []Recite {
	var lexicons []Glossary = Oration(orations, PREDICATE, rate)
	var recites []Recite
	var verbs []Glossary
	for _, predicate := range lexicons {
		for _, term := range predicate.First.Etiology {
			if term == VERB {
				var verb Glossary
				verb.First = predicate.First
				verb.Second = predicate.Second
				verb.Order = predicate.Order
				verbs = append(verbs, verb)
			}
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
	var adverbials_verbs []Recite = MountVerb(verbs)
	recites = append(recites, adverbials_verbs...)
	var relations []Vocabulary
	relations = MountConnection(lexicons, PREDICATE, dome, language)
	var appositive bool = TypeAppositive(specials)
	if appositive {
		var declaim []Recite = MountAppositive(lexicons)
		recites = append(recites, declaim...)
		return recites
	}
	var fit int = len(adverbials_verbs)
	if fit > 1 {
		var declaim []Recite = MountPeriod(lexicons, relations)
		recites = append(recites, declaim...)
		return recites
	}
	if relations != nil {
		var declaim []Recite = PeriodCompound(lexicons, relations)
		recites = append(recites, declaim...)
	} else {
		var declaim []Recite = PeriodSample(lexicons)
		recites = append(recites, declaim...)
	}
	return recites
}

func Syntax(orations []Phrase, dome brand.Arbor, language string) []Recite {
	var rate int = RATE
	var subject Recite = SyntaxSubject(orations, dome, language, rate)
	var predicates []Recite = SyntaxPredicate(orations, dome, language, rate)
	var phrases []Recite
	phrases = append(phrases, subject)
	phrases = append(phrases, predicates...)
	return phrases
}

func Oration(orations []Phrase, sentence string, rate int) []Glossary {
	var words []Talk
	var order int = 1
	for _, oration := range orations {
		for _, word := range oration.Word {
			var talk Talk
			talk.Etiology = append(talk.Etiology, word.Class)
			talk.Pattern = append(talk.Pattern, word.Sentence)
			talk.Term = word.Term
			talk.Order = order
			words = append(words, talk)
			order++
		}
	}
	var filters []Talk
	for _, subject := range words {
		for _, term := range subject.Pattern {
			if term == sentence {
				filters = append(filters, subject)
			}
		}
	}
	var locutions []Recite
	var locution Recite
	locution.Talk = filters
	locutions = append(locutions, locution)
	var glossaries []Glossary = SetGlossary(locutions)
	var lexicons []Glossary = MountGlossary(glossaries, rate)
	return lexicons
}

func SetGlossary(word []Recite) []Glossary {
	var embendding []Glossary
	var count int = 0
	var size = len(word)
	for count < size {
		var spell Recite = word[count]
		var length = len(spell.Talk)
		var caption []Talk = spell.Talk
		for index := range spell.Talk {
			if index+1 == length {
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

func Right(glossaries []Glossary, order int, quantity int) []Glossary {
	var lenght int = len(glossaries)
	var sequence int = order + quantity
	if sequence >= lenght {
		sequence = lenght - 1
	}
	var lexicon []Glossary
	var next int = 1
	for count := order; count < sequence; count++ {
		var first Talk = glossaries[order].First
		var second Talk = glossaries[order+next].Second
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
		var first Talk = glossaries[order].First
		var second Talk = glossaries[order-next].First
		var dictionary Glossary
		dictionary.First = first
		dictionary.Second = second
		dictionary.Order = next * -1
		lexicon = append(lexicon, dictionary)
		next++
	}
	return lexicon
}
