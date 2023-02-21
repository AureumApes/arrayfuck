package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arr [3000][3000]byte
	var currFirst = 0
	var currSecond = 0
	var args = os.Args[1:]

	if len(args) != 1 {
		fmt.Print("Missing file argument")
		return
	}

	file, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < len(file); i++ {
		switch file[i] {
		case '+':
			arr[currFirst][currSecond] = arr[currFirst][currSecond] + 1
		case '-':
			arr[currFirst][currSecond] = arr[currFirst][currSecond] - 1
		case ',':
			reader := bufio.NewReader(os.Stdin)
			char, err := reader.ReadByte()
			if err != nil {
				fmt.Print(err)
				return
			}
			arr[currFirst][currSecond] = char
		case '.':
			fmt.Print(string(arr[currFirst][currSecond]))
		case '[':
			if arr[currFirst][currSecond] == 0 {
				var skips = 1
				for skips != 0 {
					i = i + 1
					if file[i] == 91 {
						skips = skips + 1
					}
					if file[i] == 93 {
						skips = skips - 1
					}
				}
			}
		case ']':
			if arr[currFirst][currSecond] != 0 {
				var skips = 1
				for skips != 0 {
					i = i - 1
					if file[i] == 93 {
						skips = skips + 1
					}
					if file[i] == 91 {
						skips = skips - 1
					}
				}
			}
		case '>':
			currFirst = currFirst + 1
		case '<':
			currFirst = currFirst - 1
		case '^':
			currFirst = 0
			currSecond = currSecond + 1
		case 'v':
			currFirst = 0
			currSecond = currSecond - 1
		case '*':
			fmt.Printf("%v", arr[currFirst])
		case 'u':
			arr[currFirst][currSecond+1] = arr[currFirst][currSecond]
		case 'd':
			arr[currFirst][currSecond+1] = arr[currFirst][currSecond]
		}
	}
}
