package mapsynonyms

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readLine(reader *bufio.Reader) (string, bool) {
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", false
	}

	return line, true
}

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

	var lineCount int
	if _, err := fmt.Fscan(reader, &lineCount); err != nil {
		return
	}

	if _, ok := readLine(reader); !ok {
		return
	}

	lines := make([]string, lineCount)
	for i := range lineCount {
		line, ok := readLine(reader)
		if !ok {
			return
		}

		lines[i] = line
	}

	targetWord, ok := readLine(reader)
	if !ok {
		return
	}

	targetWord = strings.TrimSpace(targetWord)

	synonym, ok := MapSynonyms(lines, targetWord)
	if ok {
		if _, err := fmt.Fprintln(os.Stdout, synonym); err != nil {
			return
		}
	} else {
		if _, err := fmt.Fprintln(os.Stdout, "None"); err != nil {
			return
		}
	}
}
