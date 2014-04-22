

Fuzzy
=====

Fuzzy is a very fast spell checker and query suggester written in Golang. 

Motivation:
- Sajari uses very large queries (hundreds of words) but needs to respond sub-second to these queries where possible. Common spell check algorithms are quite resource hungry
- The aim was to achieve spell checks in sub 100usec per word (10kHz single core) with at least 70% accuracy. 
- Currently we can get sub 200usec per word and 68% accuracy for a Levenshtein distance of 2 chars (test set comes from Peter Norvig's article, see http://norvig.com/spell-correct.html). 
- A 500 word query can be spell checked in ~0.1 sec / cpu cores, which is good enough for us.

Notes:
- It is currently executed as a single goroutine per lookup, so undoubtedly this could be much faster using multiple cores, but currently the speed is quite good.
- Accuracy is hit slightly because several words either appear once or don't appear in the training text (data/big.txt).
- Fuzzy is a "Symmetric Delete Spelling Corrector", which relates to some blogs by Wolf Garbe at Faroo.com (see http://blog.faroo.com/2012/06/07/improved-edit-distance-based-spelling-correction/)


Usage:
See the fuzzy_test.go file to train with the big.text file or use the single word and multiword training functions to integrate with your application.


```go
package main 

import(
	"github.com/sajari/fuzzy"
	"fmt"
)

func main() {
	model := fuzzy.NewModel()

	// For testing only, this is not advisable on production
	model.SetThreshold(1)

	// This expands the area searched, but costs more resources. For spell checking, 2 is typically best, for suggestions this can be higher
	model.SetDepth(5)

	// Train multiple words simultaneously
	words := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}
	model.Train(words)
	
	// Train word by word
	model.TrainWord("single")

	// Check Spelling
	fmt.Println("\nSPELL CHECKS")
	fmt.Println("	Deletion test (yor) : ", model.SpellCheck("yor"))
	fmt.Println("	Swap test (uncel) : ", model.SpellCheck("uncel"))
	fmt.Println("	Replace test (dynemite) : ", model.SpellCheck("dynemite"))
	fmt.Println("	Insert test (dellicate) : ", model.SpellCheck("dellicate"))
	fmt.Println("	Two char test (dellicade) : ", model.SpellCheck("dellicade"))

	// Suggest completions
	fmt.Println("\nQUERY SUGGESTIONS")
	fmt.Println("	\"bigge\". Did you mean?: ", model.Suggestions("bigge", false))
	fmt.Println("	\"bo\". Did you mean?: ", model.Suggestions("bo", false))
	fmt.Println("	\"dyn\". Did you mean?: ", model.Suggestions("dyn", false))

}
```