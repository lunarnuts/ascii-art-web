package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(args []string) (string, int) {

	if len(args) == 0 {
		return "Error", 1
	}

	fontStyle := "standard.txt"
	if len(args) > 1 {
		fontStyle = args[1] + ".txt"
	}

	file, err := os.Open("./ascii/" + fontStyle)
	if err != nil {
		fmt.Println(err.Error())
		return "Error", 1
	}

	defer func() (string, int) {
		if err = file.Close(); err != nil {
			fmt.Println(err)
			return "Error", 1
		}
		return "", 200
	}()

	ascii := GetASCII(file)

	buf := make([]string, 8)

	str := strings.Split(args[0], string(byte(10)))

	res := ""
	ind := 0
	for _, a := range str {
		for _, n := range a {
			ind++
			if n < 32 || n > 126 {
				fmt.Println("Error: Message is not valid")
				fmt.Println(n)
				if n == 10 {
					continue
				}
				return "Error", 1
			}
			buf = AddCh(buf, ascii[n])
		}
		for i := range buf {
			res += buf[i] + "\n"
		}
		buf = make([]string, 8)
	}
	return res, 0
}

//GetFlag function gets value of flag and deletes it from args
func GetFlag(flag string, args []string) (string, []string) {
	value := ""
	l := len(flag)
	for i := 0; i < len(args); i++ {
		if len(args[i]) > l {
			if args[i][:l] == flag {
				value = args[i][l:]
				args = append(args[:i], args[i+1:]...)
				i--
			}
		}
	}
	return value, args
}

//GetASCII fuction writes all characters to array
func GetASCII(file *os.File) map[rune][]string {
	ascii := make(map[rune][]string, 95)
	scanner := bufio.NewScanner(file)
	buf := make([]string, 8)
	charLine := 0
	asciiChar := 32
	for scanner.Scan() {
		if scanner.Text() == "" {
			charLine = 0
			buf = nil
			continue
		} else {
			buf = append(buf, scanner.Text())
			if asciiChar == 127 { //break the loop when, 96 char is read, as there are only 95 chars
				asciiChar = 0
				break
			}
			if charLine == 7 {
				ascii[rune(asciiChar)] = buf
				buf = nil
				charLine = 0
				asciiChar++
				continue
			}
			charLine++
		}
	}
	return ascii
}

//AddCh function adds characters one by one
func AddCh(buf, new []string) []string {
	for i := range buf {
		buf[i] = buf[i] + new[i]
	}
	return buf
}

//Prints function prints ASCII message by lines
func Prints(buf string) {
	for i := range buf {
		fmt.Print(string(buf[i]))
	}
}

//GetLines function separates lines given in file
func GetLines(content []byte) []string {
	lines := []string{}
	currLine := ""
	for i := 0; i < len(content); i++ {
		currLine += string(content[i])
		if content[i] == '\n' {
			lines = append(lines, currLine)
			currLine = ""
		}
	}
	return lines
}

func toLower(str string) string {
	r := []rune(str)
	for n := range r {
		if (r[n] > 64) && (r[n] < 91) {
			r[n] = r[n] + rune(32)
		}
	}
	return string(r)
}
