package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {

		file := "standard.txt"
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		ascii := []string{}
		for scanner.Scan() {
			s := strings.ReplaceAll(scanner.Text(), "/n", "")
			ascii = append(ascii, s)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if checker(file) {
			argument := os.Args[1]
			if !check(argument) {
				return
			}
			For_Letters(ascii, argument)
		}
	}
}

func check(s string) bool {
	for _, i := range s {
		if i < 32 || i > 126 {
			return false
		}
	}
	return true
}

func checker(a string) bool {
	switch a {
	case "standard.txt":
		if hash(a) == "ac85e83127e49ec42487f272d9b9db8b" {
			return true
		}
	case "shadow.txt":
		if hash(a) == "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return true
		}

	case "thinkertoy.txt":
		if hash(a) == "eef471ad03be9d13027560644dda8359" {
			return true
		}
	}
	return false
}

func For_Letters(s []string, a string) {
	if len(a) > 0 {
		e := map[rune][]string{}
		var b string
		var q rune = 32
		count := 0
		for i := 1; i < len(s); i += 9 {
			e[q] = s[i : i+8]
			q++
		}
		for _, v := range a {
			if string(v) == "\\" {
				count++
			}
		}
		if count*2 == len(a) {
			for i := 0; i < count; i++ {
				fmt.Println()
			}
			return
		}
		k := strings.ReplaceAll(a, "\\n", "\n")
		l := strings.Split(k, "\n")

		for _, w := range l {
			if w == "" {
				b += "\n"
			} else {
				for i := 0; i < 8; i++ {
					for t := 0; t < len(w); t++ {
						if w[t] >= 32 && w[t] <= 126 {
							b += e[rune(w[t])][i]
						}
					}
					b += "\n"
				}
			}
		}
		fmt.Print(b)
	}
}

func hash(s string) string {
	h := md5.New()
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	a := fmt.Sprintf("%x", h.Sum(nil))
	return a
}
