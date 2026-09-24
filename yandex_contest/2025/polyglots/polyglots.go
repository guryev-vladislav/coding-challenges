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

	for i := range studentLanguages {
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

	allUniqueLanguages := make([]string, 0, len(allLanguages))
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

func readStudents(scanner *bufio.Scanner, studentCount int) ([][]string, bool) {
	studentLanguages := make([][]string, studentCount)
	for studentIndex := range studentCount {
		if !scanner.Scan() {
			return nil, false
		}

		languageCount, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, false
		}

		studentLanguages[studentIndex] = make([]string, languageCount)
		for languageIndex := range languageCount {
			if !scanner.Scan() {
				return nil, false
			}

			studentLanguages[studentIndex][languageIndex] = scanner.Text()
		}
	}

	return studentLanguages, true
}

func writePolyglots(writer *bufio.Writer, polyglots *PolyglotsLanguages) error {
	if _, err := fmt.Fprintln(writer, len(polyglots.commonLanguages)); err != nil {
		return err
	}

	for _, language := range polyglots.commonLanguages {
		if _, err := fmt.Fprintln(writer, language); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(writer, len(polyglots.allUniqueLanguages)); err != nil {
		return err
	}

	for _, language := range polyglots.allUniqueLanguages {
		if _, err := fmt.Fprintln(writer, language); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	studentCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	studentLanguages, ok := readStudents(scanner, studentCount)
	if !ok {
		return
	}

	if err := scanner.Err(); err != nil {
		return
	}

	polyglots := Polyglots(studentLanguages)

	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			return
		}
	}()

	if err := writePolyglots(writer, polyglots); err != nil {
		return
	}
}
