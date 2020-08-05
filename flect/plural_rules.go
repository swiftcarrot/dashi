package flect

import (
	"regexp"
	"strings"
)

type inflection struct {
	regexp  *regexp.Regexp
	replace string
}

// Regular is a regexp find replace inflection
type Regular struct {
	find    string
	replace string
}

// Irregular is a hard replace inflection,
// containing both singular and plural forms
type Irregular struct {
	singular string
	plural   string
}

// RegularSlice is a slice of Regular inflections
type RegularSlice []Regular

// IrregularSlice is a slice of Irregular inflections
type IrregularSlice []Irregular

var pluralInflections = RegularSlice{
	{"([a-z])$", "${1}s"},
	{"s$", "s"},
	{"^(ax|test)is$", "${1}es"},
	{"(octop|vir)us$", "${1}i"},
	{"(octop|vir)i$", "${1}i"},
	{"(alias|status)$", "${1}es"},
	{"(bu)s$", "${1}ses"},
	{"(buffal|tomat)o$", "${1}oes"},
	{"([ti])um$", "${1}a"},
	{"([ti])a$", "${1}a"},
	{"sis$", "ses"},
	{"(?:([^f])fe|([lr])f)$", "${1}${2}ves"},
	{"(hive)$", "${1}s"},
	{"([^aeiouy]|qu)y$", "${1}ies"},
	{"(o|z|x|ch|ss|sh)$", "${1}es"},
	{"(matr|vert|ind)(?:ix|ex)$", "${1}ices"},
	{"^(m|l)ouse$", "${1}ice"},
	{"^(m|l)ice$", "${1}ice"},
	{"^(ox)$", "${1}en"},
	{"^(oxen)$", "${1}"},
	{"(quiz)$", "${1}zes"},
}

var singularInflections = RegularSlice{
	{"s$", ""},
	{"(ss)$", "${1}"},
	{"(n)ews$", "${1}ews"},
	{"([ti])a$", "${1}um"},
	{"((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$", "${1}sis"},
	{"(^analy)(sis|ses)$", "${1}sis"},
	{"([^f])ves$", "${1}fe"},
	{"(hive)s$", "${1}"},
	{"(tive)s$", "${1}"},
	{"([lr])ves$", "${1}f"},
	{"([^aeiouy]|qu)ies$", "${1}y"},
	{"(s)eries$", "${1}eries"},
	{"(m)ovies$", "${1}ovie"},
	{"(c)ookies$", "${1}ookie"},
	{"(z|x|ch|ss|sh)es$", "${1}"},
	{"^(m|l)ice$", "${1}ouse"},
	{"(bus)(es)?$", "${1}"},
	{"(o)es$", "${1}"},
	{"(shoe)s$", "${1}"},
	{"(cris|test)(is|es)$", "${1}is"},
	{"^(a)x[ie]s$", "${1}xis"},
	{"(octop|vir)(us|i)$", "${1}us"},
	{"(alias|status)(es)?$", "${1}"},
	{"^(ox)en", "${1}"},
	{"(vert|ind)ices$", "${1}ex"},
	{"(matr)ices$", "${1}ix"},
	{"(quiz)zes$", "${1}"},
	{"(database)s$", "${1}"},
}

var compiledPluralMaps []inflection
var compiledSingularMaps []inflection

// AddPlural adds a plural inflection
func AddPlural(find, replace string) {
	pluralInflections = append(pluralInflections, Regular{find, replace})
	compile()
}

// AddSingular adds a singular inflection
func AddSingular(find, replace string) {
	singularInflections = append(singularInflections, Regular{find, replace})
	compile()
}

var singleToPlural = map[string]string{
	"aircraft":    "aircraft",
	"alias":       "aliases",
	"alumna":      "alumnae",
	"alumnus":     "alumni",
	"analysis":    "analyses",
	"antenna":     "antennas",
	"antithesis":  "antitheses",
	"apex":        "apexes",
	"appendix":    "appendices",
	"axis":        "axes",
	"bacillus":    "bacilli",
	"bacterium":   "bacteria",
	"basis":       "bases",
	"beau":        "beaus",
	"bison":       "bison",
	"bureau":      "bureaus",
	"bus":         "buses",
	"cactus":      "cacti",
	"campus":      "campuses",
	"caucus":      "caucuses",
	"château":     "châteaux",
	"circus":      "circuses",
	"codex":       "codices",
	"concerto":    "concertos",
	"corpus":      "corpora",
	"crisis":      "crises",
	"criterion":   "criteria",
	"curriculum":  "curriculums",
	"datum":       "data",
	"dear":        "dear",
	"deer":        "deer",
	"diagnosis":   "diagnoses",
	"die":         "dice",
	"dwarf":       "dwarves",
	"ellipsis":    "ellipses",
	"equipment":   "equipment",
	"erratum":     "errata",
	"faux pas":    "faux pas",
	"fez":         "fezzes",
	"fish":        "fish",
	"focus":       "foci",
	"foo":         "foos",
	"foot":        "feet",
	"formula":     "formulas",
	"fungus":      "fungi",
	"gas":         "gasses",
	"genus":       "genera",
	"goose":       "geese",
	"graffito":    "graffiti",
	"grouse":      "grouse",
	"half":        "halves",
	"halo":        "halos",
	"hoof":        "hooves",
	"human":       "humans",
	"hypothesis":  "hypotheses",
	"index":       "indices",
	"information": "information",
	"larva":       "larvae",
	"libretto":    "librettos",
	"loaf":        "loaves",
	"locus":       "loci",
	"louse":       "lice",
	"matrix":      "matrices",
	"minutia":     "minutiae",
	"money":       "money",
	"moose":       "moose",
	"mouse":       "mice",
	"move":        "moves", // avoid moves => mofe
	"nebula":      "nebulae",
	"news":        "news",
	"nucleus":     "nuclei",
	"oasis":       "oases",
	"octopus":     "octopi",
	"offspring":   "offspring",
	"opus":        "opera",
	"ovum":        "ova",
	"ox":          "oxen",
	"parenthesis": "parentheses",
	"phenomenon":  "phenomena",
	"photo":       "photos",
	"phylum":      "phyla",
	"piano":       "pianos",
	"plus":        "pluses",
	"police":      "police",
	"portfolio":   "portfolios",
	"prize":       "prizes", // avoid prizes => priz
	"prognosis":   "prognoses",
	"prometheus":  "prometheuses",
	"quiz":        "quizzes",
	"radius":      "radiuses",
	"referendum":  "referendums",
	"ress":        "resses",
	"rice":        "rice",
	"salmon":      "salmon",
	"sex":         "sexes",
	"series":      "series",
	"sheep":       "sheep",
	"shrimp":      "shrimp",
	"species":     "species",
	"stimulus":    "stimuli",
	"stratum":     "strata",
	"swine":       "swine",
	"syllabus":    "syllabi",
	"symposium":   "symposiums",
	"synopsis":    "synopses",
	"tableau":     "tableaus",
	"testis":      "testes",
	"thesis":      "theses",
	"thief":       "thieves",
	"tooth":       "teeth",
	"trout":       "trout",
	"tuna":        "tuna",
	"vertebra":    "vertebrae",
	"vertix":      "vertices",
	"vita":        "vitae",
	"vortex":      "vortices",
	"wharf":       "wharves",
	"wife":        "wives",
	"wolf":        "wolves",
	"you":         "you",
}

// words can be suffix of other words
//  salesperson -> salespeople
var singularToPluralSuffix = map[string]string{
	"child":  "children",
	"jeans":  "jeans",
	"man":    "men",
	"person": "people",
	"shoe":   "shoes",
	"video":  "videos",
	"woman":  "women",
}

var pluralToSingle = map[string]string{}

func buildInflection(str string, replace string) *inflection {
	return &inflection{
		regexp:  regexp.MustCompile("([^0-9A-Za-z])" + str + "$"),
		replace: "${1}" + replace,
	}
}

func compile() {
	compiledPluralMaps = []inflection{}
	compiledSingularMaps = []inflection{}

	for k, v := range singleToPlural {
		pluralToSingle[v] = k

		compiledPluralMaps = append(compiledPluralMaps, []inflection{
			*buildInflection(strings.Title(k), strings.Title(v)),
			*buildInflection(strings.ToUpper(k), strings.ToUpper(v)),
			*buildInflection(k, v),

			*buildInflection(strings.Title(v), strings.Title(v)),
			*buildInflection(strings.ToUpper(v), strings.ToUpper(v)),
			*buildInflection(v, v),
		}...)

		compiledSingularMaps = append(compiledSingularMaps, []inflection{
			*buildInflection(strings.Title(v), strings.Title(k)),
			*buildInflection(strings.ToUpper(v), strings.ToUpper(k)),
			*buildInflection(v, k),
		}...)
	}

	for k, v := range singularToPluralSuffix {
		singleToPlural[k] = v
		pluralToSingle[v] = k

		compiledPluralMaps = append(compiledPluralMaps, []inflection{
			{regexp: regexp.MustCompile(strings.Title(k) + "$"), replace: strings.Title(v)},
			{regexp: regexp.MustCompile(strings.ToUpper(k) + "$"), replace: strings.ToUpper(v)},
			{regexp: regexp.MustCompile(k + "$"), replace: v},

			{regexp: regexp.MustCompile(strings.Title(v) + "$"), replace: strings.Title(v)},
			{regexp: regexp.MustCompile(strings.ToUpper(v) + "$"), replace: strings.ToUpper(v)},
			{regexp: regexp.MustCompile(v + "$"), replace: v},
		}...)

		compiledSingularMaps = append(compiledSingularMaps, []inflection{
			{regexp: regexp.MustCompile(strings.Title(v) + "$"), replace: strings.Title(k)},
			{regexp: regexp.MustCompile(strings.ToUpper(v) + "$"), replace: strings.ToUpper(k)},
			{regexp: regexp.MustCompile(v + "$"), replace: k},
		}...)
	}

	for i := len(pluralInflections) - 1; i >= 0; i-- {
		value := pluralInflections[i]
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			{regexp: regexp.MustCompile(value.find), replace: value.replace},
			{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}
		compiledPluralMaps = append(compiledPluralMaps, infs...)
	}

	for i := len(singularInflections) - 1; i >= 0; i-- {
		value := singularInflections[i]
		infs := []inflection{
			{regexp: regexp.MustCompile(strings.ToUpper(value.find)), replace: strings.ToUpper(value.replace)},
			{regexp: regexp.MustCompile(value.find), replace: value.replace},
			{regexp: regexp.MustCompile("(?i)" + value.find), replace: value.replace},
		}

		compiledSingularMaps = append(compiledSingularMaps, infs...)
	}
}

func init() {
	compile()
}
