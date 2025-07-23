package brand

type Adverb struct {
	Name     string
	Language string
	Type     string
}

type Pronoun struct {
	Name     string
	Language string
	Type     string
	Number   string
	Person   int
	Gender   string
	Context  string
}

type Article struct {
	Name     string
	Language string
	Type     string
	Number   string
	Gender   string
}

type Conjunction struct {
	Name     string
	Language string
	Type     string
}

type Numeral struct {
	Name     string
	Initial  int
	Language string
	Type     string
}

type Preposition struct {
	Name     string
	Language string
	Type     string
}

type Verb struct {
	Name     string
	Language string
	Model    string
	Mode     string
	Pronoun  string
}

type Auxiliary struct {
	Name     string
	Language string
	Model    string
	Prefix   string
	Preverb  string
	Premode  string
}

type Adjective struct {
	Name     string
	Lesson   string
	Language string
}

type Noun struct {
	Name     string
	Lesson   string
	Language string
}

type Arbor struct {
	Adverb      []Adverb
	Pronoun     []Pronoun
	Article     []Article
	Conjunction []Conjunction
	Numeral     []Numeral
	Preposition []Preposition
	Verb        []Verb
	Auxiliary   []Auxiliary
	Adjective   []Adjective
	Noun        []Noun
}
