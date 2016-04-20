package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
    "time"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func IsBufferTimestamp(buffer string) bool {

  if len(buffer) > 0 && buffer[0:1] == "#" {
    return true
  }

  return false
}

func main() {
  filename := flag.String("filename", "bash_history", "a filename to parse")

  flag.Parse()

  fmt.Println(*filename)

  f, err := os.Open(*filename)
  check(err)

	scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    buffer := scanner.Text()

    if IsBufferTimestamp(buffer) {
      timestamp, _ := strconv.ParseInt(buffer[1:len(buffer)], 10, 64)
      tm := time.Unix(timestamp, 0)
      fmt.Print(tm, " ")
    } else {
      fmt.Println(buffer)
    }
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
