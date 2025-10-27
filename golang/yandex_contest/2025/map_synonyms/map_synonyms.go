package mapsynonyms

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MapSynonyms(lines []string, targetWord string) (string, bool) {
	synonyms := make(map[string]string, len(lines)*2)
	for _, line := range lines {
		words := strings.Split(strings.TrimSpace(line), " ")
		word1, word2 := words[0], words[1]
		synonyms[word1] = word2
		synonyms[word2] = word1
	}

	synonym, ok := synonyms[targetWord]
	if ok {
		return synonym, true
	}

	return "", false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	reader.ReadString('\n')

	lines := make([]string, n)
	for i := range n {
		lines[i], _ = reader.ReadString('\n')
	}

	targetWord, _ := reader.ReadString('\n')
	targetWord = strings.TrimSpace(targetWord)

	synonym, ok := MapSynonyms(lines, targetWord)
	if ok {
		fmt.Println(synonym)
	} else {
		fmt.Println("None")
	}
}
