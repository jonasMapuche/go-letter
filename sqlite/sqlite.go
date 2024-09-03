package sqlite

import (
	"database/sql"

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

func Noun() []grammar.Noun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM substantivo")
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

	rows, err := database.Query("SELECT distinct name, language FROM verbo")
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

	rows, err := database.Query("SELECT distinct name, language FROM pronome")
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

	rows, err := database.Query("SELECT distinct name, language FROM artigo")
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

	rows, err := database.Query("SELECT distinct name, language FROM adjetivo")
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

	rows, err := database.Query("SELECT distinct name, language FROM adverbio")
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

	rows, err := database.Query("SELECT distinct name, language FROM preposicao")
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

	rows, err := database.Query("SELECT distinct name, language FROM numeral")
	checkErr(err)

	var numerals []grammar.Numeral

	for rows.Next() {
		var numeral grammar.Numeral
		err = rows.Scan(&numeral.Name, &numeral.Language)
		checkErr(err)
		numerals = append(numerals, numeral)
	}

	return numerals
}

func Conjunction() []grammar.Conjunction {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM conjuncao")
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

	rows, err := database.Query("SELECT distinct name, language FROM interjeicao")
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
