package polyglots

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type PolyglotsLanguages struct {
	commonLanguages    []string
	allUniqueLanguages []string
}

func Polyglots(studentLanguages [][]string) *PolyglotsLanguages {

	allLanguages := make(map[string]int)

	for i := range len(studentLanguages) {
		for _, language := range studentLanguages[i] {

			allLanguages[language]++
		}
	}

	commonLanguages := []string{}
	for language, count := range allLanguages {
		if count == len(studentLanguages) {
			commonLanguages = append(commonLanguages, language)
		}
	}

	allUniqueLanguages := []string{}
	for language := range allLanguages {
		allUniqueLanguages = append(allUniqueLanguages, language)
	}

	sort.Strings(commonLanguages)
	sort.Strings(allUniqueLanguages)

	polyglots := PolyglotsLanguages{
		commonLanguages:    commonLanguages,
		allUniqueLanguages: allUniqueLanguages,
	}

	return &polyglots
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	studentLanguages := make([][]string, n)

	for i := range n {
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		studentLanguages[i] = make([]string, m)
		for j := range m {
			scanner.Scan()
			language := scanner.Text()
			studentLanguages[i][j] = language
		}
	}

	polyglots := Polyglots(studentLanguages)

	fmt.Println(len(polyglots.commonLanguages))
	for _, language := range polyglots.commonLanguages {
		fmt.Println(language)
	}

	fmt.Println(len(polyglots.allUniqueLanguages))
	for _, language := range polyglots.allUniqueLanguages {
		fmt.Println(language)
	}
}
