package sqlite

import (
	"database/sql"

	"letter.go/brand"
	"letter.go/grammar"
	_ "modernc.org/sqlite"
)

const file string = "letter.db"

func Build() grammar.Arbor {
	var verb []grammar.Verb = Verb()
	var noun []grammar.Noun = Noun()
	var preposition []grammar.Preposition = Preposition()
	var article []grammar.Article = Article()
	var pronoun []grammar.Pronoun = Pronoun()
	var adjective []grammar.Adjective = Adjective()
	var adverb []grammar.Adverb = Adverb()
	var numeral []grammar.Numeral = Numeral()
	var conjunction []grammar.Conjunction = Conjunction()
	var interjection []grammar.Interjection = Interjection()

	tree := grammar.Arbor{
		Noun:         noun,
		Verb:         verb,
		Preposition:  preposition,
		Article:      article,
		Pronoun:      pronoun,
		Adjective:    adjective,
		Adverb:       adverb,
		Numeral:      numeral,
		Conjunction:  conjunction,
		Interjection: interjection,
	}

	return tree
}

func Forge() brand.Arbor {
	var adverb []brand.Adverb = Adverbs()
	var pronoun []brand.Pronoun = Pronouns()
	var article []brand.Article = Articles()
	var conjunction []brand.Conjunction = Conjunctions()
	var numeral []brand.Numeral = Numerals()
	var preposition []brand.Preposition = Prepositions()
	var verb []brand.Verb = Verbs()
	var adjective []brand.Adjective = Adjectives()
	var noun []brand.Noun = Nouns()
	var sentence []brand.Sentence = Sentences()
	var auxiliary []brand.Auxiliary = Auxiliaries()

	tree := brand.Arbor{
		Adverb:      adverb,
		Pronoun:     pronoun,
		Article:     article,
		Conjunction: conjunction,
		Numeral:     numeral,
		Preposition: preposition,
		Verb:        verb,
		Adjective:   adjective,
		Noun:        noun,
		Sentence:    sentence,
		Auxiliary:   auxiliary,
	}

	return tree
}

func Adverbs() []brand.Adverb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) AS language, LOWER(type) AS type FROM adverbios")
	checkErr(err)

	var adverbs []brand.Adverb

	for rows.Next() {
		var adverb brand.Adverb
		err = rows.Scan(&adverb.Name, &adverb.Language, &adverb.Type)
		checkErr(err)
		adverbs = append(adverbs, adverb)
	}

	return adverbs
}

func Pronouns() []brand.Pronoun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) as language, LOWER(type) AS type, LOWER(number) AS number, person, LOWER(gender) gender, LOWER(context) context FROM pronomes")
	checkErr(err)

	var pronouns []brand.Pronoun

	for rows.Next() {
		var pronoun brand.Pronoun
		err = rows.Scan(&pronoun.Name, &pronoun.Language, &pronoun.Type, &pronoun.Number, &pronoun.Person, &pronoun.Gender, &pronoun.Context)
		checkErr(err)
		pronouns = append(pronouns, pronoun)
	}

	return pronouns
}

func Articles() []brand.Article {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) AS language, LOWER(type) AS type, LOWER(number) AS number, LOWER(gender) AS gender FROM artigos")
	checkErr(err)

	var articles []brand.Article

	for rows.Next() {
		var article brand.Article
		err = rows.Scan(&article.Name, &article.Language, &article.Type, &article.Number, &article.Gender)
		checkErr(err)
		articles = append(articles, article)
	}

	return articles
}

func Conjunctions() []brand.Conjunction {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) AS language, LOWER(type) AS type FROM conjuncoes")
	checkErr(err)

	var conjunctions []brand.Conjunction

	for rows.Next() {
		var conjunction brand.Conjunction
		err = rows.Scan(&conjunction.Name, &conjunction.Language, &conjunction.Type)
		checkErr(err)
		conjunctions = append(conjunctions, conjunction)
	}

	return conjunctions
}

func Numerals() []brand.Numeral {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(initial) AS initial, LOWER(language) language, LOWER(type) AS type FROM numerais")
	checkErr(err)

	var numerals []brand.Numeral

	for rows.Next() {
		var numeral brand.Numeral
		err = rows.Scan(&numeral.Name, &numeral.Initial, &numeral.Language, &numeral.Type)
		checkErr(err)
		numerals = append(numerals, numeral)
	}

	return numerals
}

func Prepositions() []brand.Preposition {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) AS language, LOWER(type) AS type FROM preposicoes")
	checkErr(err)

	var prepositions []brand.Preposition

	for rows.Next() {
		var preposition brand.Preposition
		err = rows.Scan(&preposition.Name, &preposition.Language, &preposition.Type)
		checkErr(err)
		prepositions = append(prepositions, preposition)
	}

	return prepositions
}

func Verbs() []brand.Verb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) AS language, LOWER(model) AS model, LOWER(mode) AS mode, LOWER(pronoun) AS pronoun FROM verbos")
	checkErr(err)

	var verbs []brand.Verb

	for rows.Next() {
		var verb brand.Verb
		err = rows.Scan(&verb.Name, &verb.Language, &verb.Model, &verb.Mode, &verb.Pronoun)
		checkErr(err)
		verbs = append(verbs, verb)
	}

	return verbs
}

func Adjectives() []brand.Adjective {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(lesson) AS lesson, LOWER(language) AS language FROM adjetivo")
	checkErr(err)

	var adjectives []brand.Adjective

	for rows.Next() {
		var adjective brand.Adjective
		err = rows.Scan(&adjective.Name, &adjective.Lesson, &adjective.Language)
		checkErr(err)
		adjectives = append(adjectives, adjective)
	}

	return adjectives
}

func Nouns() []brand.Noun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(lesson) AS lesson, LOWER(language) AS language FROM substantivo")
	checkErr(err)

	var nouns []brand.Noun

	for rows.Next() {
		var noun brand.Noun
		err = rows.Scan(&noun.Name, &noun.Lesson, &noun.Language)
		checkErr(err)
		nouns = append(nouns, noun)
	}

	return nouns
}

func Sentences() []brand.Sentence {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(language) AS language, LOWER(impulse) AS impulse, LOWER(rest) AS rest FROM sentencas")
	checkErr(err)

	var sentences []brand.Sentence

	for rows.Next() {
		var sentence brand.Sentence
		err = rows.Scan(&sentence.Language, &sentence.Impulse, &sentence.Rest)
		checkErr(err)
		sentences = append(sentences, sentence)
	}

	return sentences
}

func Auxiliaries() []brand.Auxiliary {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, LOWER(language) as language, LOWER(mode) AS mode, LOWER(prefix) AS prefix, LOWER(preverb) AS preverb, LOWER(premode) AS premode FROM auxiliares")
	checkErr(err)

	var auxiliaries []brand.Auxiliary

	for rows.Next() {
		var auxiliary brand.Auxiliary
		err = rows.Scan(&auxiliary.Name, &auxiliary.Language, &auxiliary.Mode, &auxiliary.Prefix, &auxiliary.Preverb, &auxiliary.Premode)
		checkErr(err)
		auxiliaries = append(auxiliaries, auxiliary)
	}

	return auxiliaries
}

func Noun() []grammar.Noun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM substantivo")
	checkErr(err)

	var nouns []grammar.Noun

	for rows.Next() {
		var noun grammar.Noun
		err = rows.Scan(&noun.Name, &noun.Language)
		checkErr(err)
		nouns = append(nouns, noun)
	}

	return nouns
}

func Verb() []grammar.Verb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM verbos")
	checkErr(err)

	var verbs []grammar.Verb

	for rows.Next() {

		var verb grammar.Verb
		err = rows.Scan(&verb.Name, &verb.Language)
		checkErr(err)
		verbs = append(verbs, verb)
	}

	return verbs
}

func Pronoun() []grammar.Pronoun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM pronomes")
	checkErr(err)

	var pronouns []grammar.Pronoun

	for rows.Next() {
		var pronoun grammar.Pronoun
		err = rows.Scan(&pronoun.Name, &pronoun.Language)
		checkErr(err)
		pronouns = append(pronouns, pronoun)
	}

	return pronouns
}

func Article() []grammar.Article {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM artigos")
	checkErr(err)

	var articles []grammar.Article

	for rows.Next() {
		var article grammar.Article
		err = rows.Scan(&article.Name, &article.Language)
		checkErr(err)
		articles = append(articles, article)
	}

	return articles
}

func Adjective() []grammar.Adjective {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM adjetivo")
	checkErr(err)

	var adjectives []grammar.Adjective

	for rows.Next() {
		var adjective grammar.Adjective
		err = rows.Scan(&adjective.Name, &adjective.Language)
		checkErr(err)
		adjectives = append(adjectives, adjective)
	}

	return adjectives
}

func Adverb() []grammar.Adverb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM adverbios")
	checkErr(err)

	var adverbs []grammar.Adverb

	for rows.Next() {
		var adverb grammar.Adverb
		err = rows.Scan(&adverb.Name, &adverb.Language)
		checkErr(err)
		adverbs = append(adverbs, adverb)
	}

	return adverbs
}

func Preposition() []grammar.Preposition {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM preposicoes")
	checkErr(err)

	var prepositions []grammar.Preposition

	for rows.Next() {
		var preposition grammar.Preposition
		err = rows.Scan(&preposition.Name, &preposition.Language)
		checkErr(err)
		prepositions = append(prepositions, preposition)
	}

	return prepositions
}

func Numeral() []grammar.Numeral {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, initial AS initial, language FROM numerais")
	checkErr(err)

	var numerals []grammar.Numeral

	for rows.Next() {
		var numeral grammar.Numeral
		err = rows.Scan(&numeral.Name, &numeral.Initial, &numeral.Language)
		checkErr(err)
		numerals = append(numerals, numeral)
	}

	return numerals
}

func Conjunction() []grammar.Conjunction {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM conjuncoes")
	checkErr(err)

	var conjunctions []grammar.Conjunction

	for rows.Next() {
		var conjunction grammar.Conjunction
		err = rows.Scan(&conjunction.Name, &conjunction.Language)
		checkErr(err)
		conjunctions = append(conjunctions, conjunction)
	}

	return conjunctions
}

func Interjection() []grammar.Interjection {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct LOWER(name) AS name, language FROM interjeicao")
	checkErr(err)

	var interjections []grammar.Interjection

	for rows.Next() {
		var interjection grammar.Interjection
		err = rows.Scan(&interjection.Name, &interjection.Language)
		checkErr(err)
		interjections = append(interjections, interjection)
	}

	return interjections
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
