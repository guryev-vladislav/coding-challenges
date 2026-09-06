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
		fmt.Println("Error:", err)
		return
	}

	scanner.Scan()
	input := scanner.Text()

	ok := PhoneNumberCurrent(input)
	if ok {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
