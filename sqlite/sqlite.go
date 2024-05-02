package sqlite

import (
	"database/sql"

	"letter.go/tree"
	_ "modernc.org/sqlite"
)

const file string = "sqlite/letter.db"

func Build() tree.Arbor {
	var verb []tree.Verb = Verb()
	var noun []tree.Noun = Noun()
	var preposition []tree.Preposition = Preposition()
	var article []tree.Article = Article()
	var pronoun []tree.Pronoun = Pronoun()
	var adjective []tree.Adjective = Adjective()
	var adverb []tree.Adverb = Adverb()
	var numeral []tree.Numeral = Numeral()
	var conjunction []tree.Conjunction = Conjunction()
	var interjection []tree.Interjection = Interjection()

	tree := tree.Arbor{
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

func Noun() []tree.Noun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM substantivo")
	checkErr(err)

	var nouns []tree.Noun

	for rows.Next() {
		var noun tree.Noun
		err = rows.Scan(&noun.Name, &noun.Language)
		checkErr(err)
		nouns = append(nouns, noun)
	}

	return nouns
}

func Verb() []tree.Verb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM verbo")
	checkErr(err)

	var verbs []tree.Verb

	for rows.Next() {

		var verb tree.Verb
		err = rows.Scan(&verb.Name, &verb.Language)
		checkErr(err)
		verbs = append(verbs, verb)
	}

	return verbs
}

func Pronoun() []tree.Pronoun {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM pronome")
	checkErr(err)

	var pronouns []tree.Pronoun

	for rows.Next() {
		var pronoun tree.Pronoun
		err = rows.Scan(&pronoun.Name, &pronoun.Language)
		checkErr(err)
		pronouns = append(pronouns, pronoun)
	}

	return pronouns
}

func Article() []tree.Article {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM artigo")
	checkErr(err)

	var articles []tree.Article

	for rows.Next() {
		var article tree.Article
		err = rows.Scan(&article.Name, &article.Language)
		checkErr(err)
		articles = append(articles, article)
	}

	return articles
}

func Adjective() []tree.Adjective {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM adjetivo")
	checkErr(err)

	var adjectives []tree.Adjective

	for rows.Next() {
		var adjective tree.Adjective
		err = rows.Scan(&adjective.Name, &adjective.Language)
		checkErr(err)
		adjectives = append(adjectives, adjective)
	}

	return adjectives
}

func Adverb() []tree.Adverb {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM adverbio")
	checkErr(err)

	var adverbs []tree.Adverb

	for rows.Next() {
		var adverb tree.Adverb
		err = rows.Scan(&adverb.Name, &adverb.Language)
		checkErr(err)
		adverbs = append(adverbs, adverb)
	}

	return adverbs
}

func Preposition() []tree.Preposition {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM preposicao")
	checkErr(err)

	var prepositions []tree.Preposition

	for rows.Next() {
		var preposition tree.Preposition
		err = rows.Scan(&preposition.Name, &preposition.Language)
		checkErr(err)
		prepositions = append(prepositions, preposition)
	}

	return prepositions
}

func Numeral() []tree.Numeral {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM numeral")
	checkErr(err)

	var numerals []tree.Numeral

	for rows.Next() {
		var numeral tree.Numeral
		err = rows.Scan(&numeral.Name, &numeral.Language)
		checkErr(err)
		numerals = append(numerals, numeral)
	}

	return numerals
}

func Conjunction() []tree.Conjunction {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM conjuncao")
	checkErr(err)

	var conjunctions []tree.Conjunction

	for rows.Next() {
		var conjunction tree.Conjunction
		err = rows.Scan(&conjunction.Name, &conjunction.Language)
		checkErr(err)
		conjunctions = append(conjunctions, conjunction)
	}

	return conjunctions
}

func Interjection() []tree.Interjection {
	database, err := sql.Open("sqlite", file)
	checkErr(err)

	rows, err := database.Query("SELECT distinct name, language FROM interjeicao")
	checkErr(err)

	var interjections []tree.Interjection

	for rows.Next() {
		var interjection tree.Interjection
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
