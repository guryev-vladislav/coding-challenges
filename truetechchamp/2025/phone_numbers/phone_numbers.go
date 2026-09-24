package phonenumbers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PhoneNumberCurrent(input string) bool {
	if len(input)%11 != 0 {
		return false
	}

	for i := 0; i < len(input); i += 11 {
		if input[i] != '7' || input[i+1] != '9' {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		if _, printErr := fmt.Fprintln(os.Stdout, "Error:", err); printErr != nil {
			return
		}

		return
	}

	if !scanner.Scan() {
		return
	}

	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		return
	}

	ok := PhoneNumberCurrent(input)
	if ok {
		if _, err := fmt.Fprintln(os.Stdout, "1"); err != nil {
			return
		}
	} else {
		if _, err := fmt.Fprintln(os.Stdout, "0"); err != nil {
			return
		}
	}
}
