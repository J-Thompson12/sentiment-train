package main

import (
	"strings"

	snowball "github.com/kljensen/snowball"
)

const emojiiPattern = "[\\x{2712}\\x{2714}\\x{2716}\\x{271d}\\x{2721}\\x{2728}\\x{2733}\\x{2734}\\x{2744}\\x{2747}\\x{274c}\\x{274e}\\x{2753}-\\x{2755}\\x{2757}\\x{2763}\\x{2764}\\x{2795}-\\x{2797}\\x{27a1}\\x{27b0}\\x{27bf}\\x{2934}\\x{2935}\\x{2b05}-\\x{2b07}\\x{2b1b}\\x{2b1c}\\x{2b50}\\x{2b55}\\x{3030}\\x{303d}\\x{1f004}\\x{1f0cf}\\x{1f170}\\x{1f171}\\x{1f17e}\\x{1f17f}\\x{1f18e}\\x{1f191}-\\x{1f19a}\\x{1f201}\\x{1f202}\\x{1f21a}\\x{1f22f}\\x{1f232}-\\x{1f23a}\\x{1f250}\\x{1f251}\\x{1f300}-\\x{1f321}\\x{1f324}-\\x{1f393}\\x{1f396}\\x{1f397}\\x{1f399}-\\x{1f39b}\\x{1f39e}-\\x{1f3f0}\\x{1f3f3}-\\x{1f3f5}\\x{1f3f7}-\\x{1f4fd}\\x{1f4ff}-\\x{1f53d}\\x{1f549}-\\x{1f54e}\\x{1f550}-\\x{1f567}\\x{1f56f}\\x{1f570}\\x{1f573}-\\x{1f579}\\x{1f587}\\x{1f58a}-\\x{1f58d}\\x{1f590}\\x{1f595}\\x{1f596}\\x{1f5a5}\\x{1f5a8}\\x{1f5b1}\\x{1f5b2}\\x{1f5bc}\\x{1f5c2}-\\x{1f5c4}\\x{1f5d1}-\\x{1f5d3}\\x{1f5dc}-\\x{1f5de}\\x{1f5e1}\\x{1f5e3}\\x{1f5ef}\\x{1f5f3}\\x{1f5fa}-\\x{1f64f}\\x{1f680}-\\x{1f6c5}\\x{1f6cb}-\\x{1f6d0}\\x{1f6e0}-\\x{1f6e5}\\x{1f6e9}\\x{1f6eb}\\x{1f6ec}\\x{1f6f0}\\x{1f6f3}\\x{1f910}-\\x{1f918}\\x{1f980}-\\x{1f984}\\x{1f9c0}\\x{3297}\\x{3299}\\x{a9}\\x{ae}\\x{203c}\\x{2049}\\x{2122}\\x{2139}\\x{2194}-\\x{2199}\\x{21a9}\\x{21aa}\\x{231a}\\x{231b}\\x{2328}\\x{2388}\\x{23cf}\\x{23e9}-\\x{23f3}\\x{23f8}-\\x{23fa}\\x{24c2}\\x{25aa}\\x{25ab}\\x{25b6}\\x{25c0}\\x{25fb}-\\x{25fe}\\x{2600}-\\x{2604}\\x{260e}\\x{2611}\\x{2614}\\x{2615}\\x{2618}\\x{261d}\\x{2620}\\x{2622}\\x{2623}\\x{2626}\\x{262a}\\x{262e}\\x{262f}\\x{2638}-\\x{263a}\\x{2648}-\\x{2653}\\x{2660}\\x{2663}\\x{2665}\\x{2666}\\x{2668}\\x{267b}\\x{267f}\\x{2692}-\\x{2694}\\x{2696}\\x{2697}\\x{2699}\\x{269b}\\x{269c}\\x{26a0}\\x{26a1}\\x{26aa}\\x{26ab}\\x{26b0}\\x{26b1}\\x{26bd}\\x{26be}\\x{26c4}\\x{26c5}\\x{26c8}\\x{26ce}\\x{26cf}\\x{26d1}\\x{26d3}\\x{26d4}\\x{26e9}\\x{26ea}\\x{26f0}-\\x{26f5}\\x{26f7}-\\x{26fa}\\x{26fd}\\x{2702}\\x{2705}\\x{2708}-\\x{270d}\\x{270f}]|\\x{23}\\x{20e3}|\\x{2a}\\x{20e3}|\\x{30}\\x{20e3}|\\x{31}\\x{20e3}|\\x{32}\\x{20e3}|\\x{33}\\x{20e3}|\\x{34}\\x{20e3}|\\x{35}\\x{20e3}|\\x{36}\\x{20e3}|\\x{37}\\x{20e3}|\\x{38}\\x{20e3}|\\x{39}\\x{20e3}|\\x{1f1e6}[\\x{1f1e8}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f2}\\x{1f1f4}\\x{1f1f6}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fd}\\x{1f1ff}]|\\x{1f1e7}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ef}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1e8}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ee}\\x{1f1f0}-\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}-\\x{1f1ff}]|\\x{1f1e9}[\\x{1f1ea}\\x{1f1ec}\\x{1f1ef}\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1ff}]|\\x{1f1ea}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ed}\\x{1f1f7}-\\x{1f1fa}]|\\x{1f1eb}[\\x{1f1ee}-\\x{1f1f0}\\x{1f1f2}\\x{1f1f4}\\x{1f1f7}]|\\x{1f1ec}[\\x{1f1e6}\\x{1f1e7}\\x{1f1e9}-\\x{1f1ee}\\x{1f1f1}-\\x{1f1f3}\\x{1f1f5}-\\x{1f1fa}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1ed}[\\x{1f1f0}\\x{1f1f2}\\x{1f1f3}\\x{1f1f7}\\x{1f1f9}\\x{1f1fa}]|\\x{1f1ee}[\\x{1f1e8}-\\x{1f1ea}\\x{1f1f1}-\\x{1f1f4}\\x{1f1f6}-\\x{1f1f9}]|\\x{1f1ef}[\\x{1f1ea}\\x{1f1f2}\\x{1f1f4}\\x{1f1f5}]|\\x{1f1f0}[\\x{1f1ea}\\x{1f1ec}-\\x{1f1ee}\\x{1f1f2}\\x{1f1f3}\\x{1f1f5}\\x{1f1f7}\\x{1f1fc}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1f1}[\\x{1f1e6}-\\x{1f1e8}\\x{1f1ee}\\x{1f1f0}\\x{1f1f7}-\\x{1f1fb}\\x{1f1fe}]|\\x{1f1f2}[\\x{1f1e6}\\x{1f1e8}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1ff}]|\\x{1f1f3}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}-\\x{1f1ec}\\x{1f1ee}\\x{1f1f1}\\x{1f1f4}\\x{1f1f5}\\x{1f1f7}\\x{1f1fa}\\x{1f1ff}]|\\x{1f1f4}\\x{1f1f2}|\\x{1f1f5}[\\x{1f1e6}\\x{1f1ea}-\\x{1f1ed}\\x{1f1f0}-\\x{1f1f3}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fc}\\x{1f1fe}]|\\x{1f1f6}\\x{1f1e6}|\\x{1f1f7}[\\x{1f1ea}\\x{1f1f4}\\x{1f1f8}\\x{1f1fa}\\x{1f1fc}]|\\x{1f1f8}[\\x{1f1e6}-\\x{1f1ea}\\x{1f1ec}-\\x{1f1f4}\\x{1f1f7}-\\x{1f1f9}\\x{1f1fb}\\x{1f1fd}-\\x{1f1ff}]|\\x{1f1f9}[\\x{1f1e6}\\x{1f1e8}\\x{1f1e9}\\x{1f1eb}-\\x{1f1ed}\\x{1f1ef}-\\x{1f1f4}\\x{1f1f7}\\x{1f1f9}\\x{1f1fb}\\x{1f1fc}\\x{1f1ff}]|\\x{1f1fa}[\\x{1f1e6}\\x{1f1ec}\\x{1f1f2}\\x{1f1f8}\\x{1f1fe}\\x{1f1ff}]|\\x{1f1fb}[\\x{1f1e6}\\x{1f1e8}\\x{1f1ea}\\x{1f1ec}\\x{1f1ee}\\x{1f1f3}\\x{1f1fa}]|\\x{1f1fc}[\\x{1f1eb}\\x{1f1f8}]|\\x{1f1fd}\\x{1f1f0}|\\x{1f1fe}[\\x{1f1ea}\\x{1f1f9}]|\\x{1f1ff}[\\x{1f1e6}\\x{1f1f2}\\x{1f1fc}]"

// stopwords are words which have very little meaning
var stopwords = map[string]struct{}{
	"i": struct{}{}, "me": struct{}{}, "my": struct{}{}, "myself": struct{}{}, "we": struct{}{}, "our": struct{}{}, "ours": struct{}{},
	"ourselves": struct{}{}, "you": struct{}{}, "your": struct{}{}, "yours": struct{}{}, "yourself": struct{}{}, "yourselves": struct{}{},
	"he": struct{}{}, "him": struct{}{}, "his": struct{}{}, "himself": struct{}{}, "she": struct{}{}, "her": struct{}{}, "hers": struct{}{},
	"herself": struct{}{}, "it": struct{}{}, "its": struct{}{}, "itself": struct{}{}, "they": struct{}{}, "them": struct{}{}, "their": struct{}{},
	"theirs": struct{}{}, "themselves": struct{}{}, "what": struct{}{}, "which": struct{}{}, "who": struct{}{}, "whom": struct{}{}, "this": struct{}{},
	"that": struct{}{}, "these": struct{}{}, "those": struct{}{}, "am": struct{}{}, "is": struct{}{}, "are": struct{}{}, "was": struct{}{},
	"were": struct{}{}, "be": struct{}{}, "been": struct{}{}, "being": struct{}{}, "have": struct{}{}, "has": struct{}{}, "had": struct{}{},
	"having": struct{}{}, "do": struct{}{}, "does": struct{}{}, "did": struct{}{}, "doing": struct{}{}, "a": struct{}{}, "an": struct{}{},
	"the": struct{}{}, "and": struct{}{}, "but": struct{}{}, "if": struct{}{}, "or": struct{}{}, "because": struct{}{}, "as": struct{}{},
	"until": struct{}{}, "while": struct{}{}, "of": struct{}{}, "at": struct{}{}, "by": struct{}{}, "for": struct{}{}, "with": struct{}{},
	"about": struct{}{}, "against": struct{}{}, "between": struct{}{}, "float64o": struct{}{}, "through": struct{}{}, "during": struct{}{},
	"before": struct{}{}, "after": struct{}{}, "above": struct{}{}, "below": struct{}{}, "to": struct{}{}, "from": struct{}{}, "up": struct{}{},
	"down": struct{}{}, "in": struct{}{}, "out": struct{}{}, "on": struct{}{}, "off": struct{}{}, "over": struct{}{}, "under": struct{}{},
	"again": struct{}{}, "further": struct{}{}, "then": struct{}{}, "once": struct{}{}, "here": struct{}{}, "there": struct{}{}, "when": struct{}{},
	"where": struct{}{}, "why": struct{}{}, "how": struct{}{}, "all": struct{}{}, "any": struct{}{}, "both": struct{}{}, "each": struct{}{},
	"few": struct{}{}, "more": struct{}{}, "most": struct{}{}, "other": struct{}{}, "some": struct{}{}, "such": struct{}{}, "no": struct{}{},
	"nor": struct{}{}, "not": struct{}{}, "only": struct{}{}, "same": struct{}{}, "so": struct{}{}, "than": struct{}{}, "too": struct{}{},
	"very": struct{}{}, "can": struct{}{}, "will": struct{}{}, "just": struct{}{}, "don't": struct{}{}, "should": struct{}{}, "should've": struct{}{},
	"now": struct{}{}, "aren't": struct{}{}, "couldn't": struct{}{}, "didn't": struct{}{}, "doesn't": struct{}{}, "hasn't": struct{}{}, "haven't": struct{}{},
	"isn't": struct{}{}, "shouldn't": struct{}{}, "wasn't": struct{}{}, "weren't": struct{}{}, "won't": struct{}{}, "wouldn't": struct{}{},
}

// IsStopword does a thing
func IsStopword(w string) bool {
	_, ok := stopwords[w]
	return ok
}

// cleanup remove none-alnum characters and lowercasize them
func cleanup(sentence string) string {

	// just copying jake here, first lets coerce hashtags and mentions float64o their own words

	sentence = strings.ReplaceAll(sentence, "#", " #")
	sentence = strings.ReplaceAll(sentence, "@", " @")
	sentence = strings.ReplaceAll(sentence, ".", "")
	sentence = strings.ReplaceAll(sentence, "?", "")
	sentence = strings.ReplaceAll(sentence, "!", "")
	sentence = strings.ReplaceAll(sentence, ",", "")
	sentence = strings.ToLower(sentence)
	//fmt.Prfloat64ln(sentence)

	// re := regexp.MustCompile(emojiiPattern)
	// fmt.Prfloat64ln(re.FindAllString(sentence, -1))
	// re := regexp.MustCompile("[^a-zA-Z 0-9]")
	return sentence
	// return re.ReplaceAllString(strings.ToLower(sentence), "")
}

// tokenize create an array of words from a sentence
func tokenize(sentence string) []string {
	s := cleanup(sentence)
	words := strings.Fields(s)
	var tokens []string
	for _, w := range words {
		if !IsStopword(w) {
			w = stem(w)
			tokens = append(tokens, w)
		}
	}
	return tokens
}

func tokenizeMulti(sentence string, size int) []string {
	words := tokenize(sentence)
	var tokens []string
	for i := 0; i+size <= len(words); i++ {
		tokens = append(tokens, strings.Join(words[i:i+size], " "))
	}
	return tokens
}

// stem a word using the Snowball algorithm from https://github.com/snowballstem/snowball
func stem(word string) string {
	stemmed, err := snowball.Stem(word, "english", true)
	if err == nil {
		return stemmed
	}
	return word
}

func countWords(document string) (wordCount map[string]float64) {
	words := tokenize(document)
	words = append(words, tokenizeMulti(document, 2)...)
	wordCount = make(map[string]float64)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}
