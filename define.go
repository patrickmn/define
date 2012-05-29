package main

import (
	"github.com/pmylund/go-wikimedia"

	// "encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

const (
	version     = "1.0"
	defaultUA   = "Define/" + version + " (http://patrickmylund.com/projects/define/)"
	defaultLang = "en"
)

var (
	langs = map[string]string{
		"aa":           "Afar",
		"ab":           "Abkhazian",
		"ace":          "Acehnese",
		"af":           "Afrikaans",
		"ak":           "Akan",
		"als":          "Alemannic",
		"am":           "Amharic",
		"an":           "Aragonese",
		"ang":          "Anglo-Saxon",
		"ar":           "Arabic",
		"arc":          "Assyrian Neo-Aramaic",
		"arz":          "Egyptian Arabic",
		"as":           "Assamese",
		"ast":          "Asturian",
		"av":           "Avar",
		"ay":           "Aymara",
		"az":           "Azerbaijani",
		"ba":           "Bashkir",
		"bar":          "Bavarian",
		"bat-smg":      "Samogitian",
		"bcl":          "Central_Bicolano",
		"be":           "Belarusian",
		"be-x-old":     "Belarusian (Tara�kievica)",
		"bg":           "Bulgarian",
		"bh":           "Bihari",
		"bi":           "Bislama",
		"bjn":          "Banjar",
		"bm":           "Bambara",
		"bn":           "Bengali",
		"bo":           "Tibetan",
		"bpy":          "Bishnupriya Manipuri",
		"br":           "Breton",
		"bs":           "Bosnian",
		"bug":          "Buginese",
		"bxr":          "Buryat (Russia)",
		"ca":           "Catalan",
		"cbk-zam":      "Zamboanga Chavacano",
		"cdo":          "Min Dong",
		"ce":           "Chechen",
		"ceb":          "Cebuano",
		"ch":           "Chamorro",
		"cho":          "Choctaw",
		"chr":          "Cherokee",
		"chy":          "Cheyenne",
		"ckb":          "Sorani",
		"co":           "Corsican",
		"cr":           "Cree",
		"crh":          "Crimean Tatar",
		"cs":           "Czech",
		"csb":          "Kashubian",
		"cu":           "Old Church Slavonic",
		"cv":           "Chuvash",
		"cy":           "Welsh",
		"da":           "Danish",
		"de":           "German",
		"diq":          "Zazaki",
		"dsb":          "Lower Sorbian",
		"dv":           "Divehi",
		"dz":           "Dzongkha",
		"ee":           "Ewe",
		"el":           "Greek",
		"eml":          "Emilian-Romagnol",
		"en":           "English",
		"eo":           "Esperanto",
		"es":           "Spanish",
		"et":           "Estonian",
		"eu":           "Basque",
		"ext":          "Extremaduran",
		"fa":           "Persian",
		"ff":           "Fula",
		"fi":           "Finnish",
		"fiu-vro":      "V�ro",
		"fj":           "Fijian",
		"fo":           "Faroese",
		"fr":           "French",
		"frp":          "Franco-Proven�al/Arpitan",
		"frr":          "North Frisian",
		"fur":          "Friulian",
		"fy":           "West Frisian",
		"ga":           "Irish",
		"gag":          "Gagauz",
		"gan":          "Gan",
		"gd":           "Scottish Gaelic",
		"gl":           "Galician",
		"glk":          "Gilaki",
		"gn":           "Guarani",
		"got":          "Gothic",
		"gu":           "Gujarati",
		"gv":           "Manx",
		"ha":           "Hausa",
		"hak":          "Hakka",
		"haw":          "Hawaiian",
		"he":           "Hebrew",
		"hi":           "Hindi",
		"hif":          "Fiji Hindi",
		"ho":           "Hiri Motu",
		"hr":           "Croatian",
		"hsb":          "Upper Sorbian",
		"ht":           "Haitian",
		"hu":           "Hungarian",
		"hy":           "Armenian",
		"hz":           "Herero",
		"ia":           "Interlingua",
		"id":           "Indonesian",
		"ie":           "Interlingue",
		"ig":           "Igbo",
		"ii":           "Sichuan Yi",
		"ik":           "Inupiak",
		"ilo":          "Ilokano",
		"io":           "Ido",
		"is":           "Icelandic",
		"it":           "Italian",
		"iu":           "Inuktitut",
		"ja":           "Japanese",
		"jbo":          "Lojban",
		"jv":           "Javanese",
		"ka":           "Georgian",
		"kaa":          "Karakalpak",
		"kab":          "Kabyle",
		"kbd":          "Kabardian Circassian",
		"kg":           "Kongo",
		"ki":           "Kikuyu",
		"kj":           "Kuanyama",
		"kk":           "Kazakh",
		"kl":           "Greenlandic",
		"km":           "Khmer",
		"kn":           "Kannada",
		"ko":           "Korean",
		"koi":          "Komi-Permyak",
		"kr":           "Kanuri",
		"krc":          "Karachay-Balkar",
		"ks":           "Kashmiri",
		"ksh":          "Ripuarian",
		"ku":           "Kurdish",
		"kv":           "Komi",
		"kw":           "Cornish",
		"ky":           "Kirghiz",
		"la":           "Latin",
		"lad":          "Ladino",
		"lb":           "Luxembourgish",
		"lbe":          "Lak",
		"lez":          "Lezgian",
		"lg":           "Luganda",
		"li":           "Limburgian",
		"lij":          "Ligurian",
		"lmo":          "Lombard",
		"ln":           "Lingala",
		"lo":           "Lao",
		"lt":           "Lithuanian",
		"ltg":          "Latgalian",
		"lv":           "Latvian",
		"map-bms":      "Banyumasan",
		"mdf":          "Moksha",
		"mg":           "Malagasy",
		"mh":           "Marshallese",
		"mhr":          "Meadow Mari",
		"mi":           "Maori",
		"mk":           "Macedonian",
		"ml":           "Malayalam",
		"mn":           "Mongolian",
		"mo":           "Moldovan",
		"mr":           "Marathi",
		"mrj":          "Hill Mari",
		"ms":           "Malay",
		"mt":           "Maltese",
		"mus":          "Muscogee",
		"mwl":          "Mirandese",
		"my":           "Burmese",
		"myv":          "Erzya",
		"mzn":          "Mazandarani",
		"na":           "Nauruan",
		"nah":          "Nahuatl",
		"nap":          "Neapolitan",
		"nds":          "Low Saxon",
		"nds-nl":       "Dutch Low Saxon",
		"ne":           "Nepali",
		"new":          "Newar / Nepal Bhasa",
		"ng":           "Ndonga",
		"nl":           "Dutch",
		"nn":           "Norwegian (Nynorsk)",
		"no":           "Norwegian (Bokm�l)",
		"nov":          "Novial",
		"nrm":          "Norman",
		"nso":          "Northern Sotho",
		"nv":           "Navajo",
		"ny":           "Chichewa",
		"oc":           "Occitan",
		"om":           "Oromo",
		"or":           "Oriya",
		"os":           "Ossetian",
		"pa":           "Punjabi",
		"pag":          "Pangasinan",
		"pam":          "Kapampangan",
		"pap":          "Papiamentu",
		"pcd":          "Picard",
		"pdc":          "Pennsylvania German",
		"pfl":          "Palatinate German",
		"pi":           "Pali",
		"pih":          "Norfolk",
		"pl":           "Polish",
		"pms":          "Piedmontese",
		"pnb":          "Western Panjabi",
		"pnt":          "Pontic",
		"ps":           "Pashto",
		"pt":           "Portuguese",
		"qu":           "Quechua",
		"rm":           "Romansh",
		"rmy":          "Romani",
		"rn":           "Kirundi",
		"ro":           "Romanian",
		"roa-rup":      "Aromanian",
		"roa-tara":     "Tarantino",
		"ru":           "Russian",
		"rue":          "Rusyn",
		"rw":           "Kinyarwanda",
		"sa":           "Sanskrit",
		"sah":          "Sakha",
		"sc":           "Sardinian",
		"scn":          "Sicilian",
		"sco":          "Scots",
		"sd":           "Sindhi",
		"se":           "Northern Sami",
		"sg":           "Sango",
		"sh":           "Serbo-Croatian",
		"si":           "Sinhalese",
		"simple":       "Simple English",
		"sk":           "Slovak",
		"sl":           "Slovenian",
		"sm":           "Samoan",
		"sn":           "Shona",
		"so":           "Somali",
		"sq":           "Albanian",
		"sr":           "Serbian",
		"srn":          "Sranan",
		"ss":           "Swati",
		"st":           "Sesotho",
		"stq":          "Saterland Frisian",
		"su":           "Sundanese",
		"sv":           "Swedish",
		"sw":           "Swahili",
		"szl":          "Silesian",
		"ta":           "Tamil",
		"te":           "Telugu",
		"tet":          "Tetum",
		"tg":           "Tajik",
		"th":           "Thai",
		"ti":           "Tigrinya",
		"tk":           "Turkmen",
		"tl":           "Tagalog",
		"tn":           "Tswana",
		"to":           "Tongan",
		"tpi":          "Tok Pisin",
		"tr":           "Turkish",
		"ts":           "Tsonga",
		"tt":           "Tatar",
		"tum":          "Tumbuka",
		"tw":           "Twi",
		"ty":           "Tahitian",
		"udm":          "Udmurt",
		"ug":           "Uyghur",
		"uk":           "Ukrainian",
		"ur":           "Urdu",
		"uz":           "Uzbek",
		"ve":           "Venda",
		"vec":          "Venetian",
		"vep":          "Vepsian",
		"vi":           "Vietnamese",
		"vls":          "West Flemish",
		"vo":           "Volap�k",
		"wa":           "Walloon",
		"war":          "Waray-Waray",
		"wo":           "Wolof",
		"wuu":          "Wu",
		"xal":          "Kalmyk",
		"xh":           "Xhosa",
		"xmf":          "Mingrelian",
		"yi":           "Yiddish",
		"yo":           "Yoruba",
		"za":           "Zhuang",
		"zea":          "Zeelandic",
		"zh":           "Chinese",
		"zh-classical": "Classical Chinese",
		"zh-min-nan":   "Min Nan",
		"zh-yue":       "Cantonese",
		"zu":           "Zulu",
	}
	def *definer
)

var (
	lang *string = flag.String("l", defaultLang, "language/dictionary to look in")
)

func showLanguages() {
	mk := make([]string, len(langs))
	i := 0
	for k, _ := range langs {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	for _, v := range mk {
		fmt.Printf("  %-15s %s\n", v, langs[v])
	}
}

type definer struct {
	w *wikimedia.Wikimedia
}

func (d *definer) resolveWord(w string) (string, error) {
	f := url.Values{
		"action":   {"query"},
		"list":     {"search"},
		"srsearch": {w},
		"exintro":  {"1"},
	}
	res, err := d.w.Query(f)
	if err != nil {
		return "", err
	}

	if len(res.Query.Search) == 0 {
		return "", fmt.Errorf("No match found for %s", w)
	}
	return res.Query.Search[0].Title, nil
}

type word struct {
	title string
	text  string
}

func (w *word) format() string {
	var out []string
	wasNewline := false
	for _, v := range strings.Split(w.text, "\n") {
		if len(v) > 1 && v[:1] == " " {
			v = fmt.Sprintf("\n%s\n-----", v[1:])
		}
		v = strings.Trim(v, " ")
		if v != "" && (!wasNewline && v != "\n") {
			out = append(out, v)
		}
	}
	return strings.Join(out, "\n")
}

func (d *definer) getWords(words []string) ([]word, error) {
	// http://en.wikipedia.org/w/api.php?action=query&prop=extracts&titles=Threadless&format=json&exintro=1
	f := url.Values{
		"action": {"query"},
		"prop":   {"extracts"},
		"titles": {strings.Join(words, "|")},
		// "exintro": {"1"},
	}
	res, err := d.w.Query(f)
	if err != nil {
		return nil, err
	}

	ws := make([]word, len(res.Query.Pages))
	i := 0
	for _, v := range res.Query.Pages {
		ws[i] = word{v.Title, v.Extract}
		i++
	}
	return ws, nil
}

func init() {
	flag.Parse()
	w, err := wikimedia.New(fmt.Sprintf("http://%s.wiktionary.org/w/api.php", *lang))
	if err != nil {
		fmt.Println("Error setting up Wikimedia library:", err)
		os.Exit(1)
	}
	w.StripHtml = true
	w.UserAgent = defaultUA
	def = &definer{
		w: w,
	}
}

func main() {
	if flag.NArg() == 0 {
		fmt.Println("Define", version)
		fmt.Println("http://patrickmylund.com/projects/define/")
		fmt.Println("-----")
		flag.Usage()
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println(" ", os.Args[0], "<word or sentence>")
		fmt.Println(" ", os.Args[0], "<word 1>,<word 2>")
		fmt.Println("")
		fmt.Println("To search for something, type", os.Args[0], "search <word or sentence>")
		fmt.Println("To see a list of available languages, type", os.Args[0], "languages")
	} else if flag.Arg(0) == "languages" {
		showLanguages()
	} else {
		in := strings.Split(strings.Join(flag.Args(), " "), ",")

		var to []string
		for _, v := range in {
			w, err := def.resolveWord(v)
			if err != nil {
				fmt.Printf("Error searching for %s: %v\n", v, err)
			} else {
				to = append(to, w)
			}
		}
		if len(to) == 0 {
			fmt.Println("No matches found")
			return
		}
		ws, err := def.getWords(to)
		if err != nil {
			fmt.Printf("Couldn't look up %s: %v\n", strings.Join(in, ", "), err)
			return
		}
		for _, v := range ws {
			fmt.Println(v.format())
		}
	}
}
