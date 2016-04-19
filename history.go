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
  filename := flag.String("filename", "bash_history", "a filename to parse")

  flag.Parse()

  fmt.Println(*filename)

  f, err := os.Open(*filename)
  check(err)

	scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    t := scanner.Text()
    fmt.Println(t[0:1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
