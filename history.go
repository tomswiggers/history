package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
  filename := flag.String("filename", "history", "a filename to parse")

  flag.Parse()

  fmt.Println(*filename)

  f, err := os.Open(*filename)
  check(err)

	scanner := bufio.NewScanner(f)

  for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
