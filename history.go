package main

import (
    "bufio"
    "fmt"
    "os"
    "flag"
    "time"
    "strconv"
)

type History struct {
  Timestamp time.Time
  Command string
}

func (h History) String() string {
    return fmt.Sprintf("%s: %s", h.Timestamp, h.Command)
}

type ByTimestamp []History

func (a ByTimestamp) Len() int           { return len(a) }
func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp.Before(a[j].Timestamp) }

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

  f, err := os.Open(*filename)
  check(err)

	scanner := bufio.NewScanner(f)

  i := 0
  item := History{}

  //@todo fix this fixed length of 10000
  items := [10000]History{}

  for scanner.Scan() {
    buffer := scanner.Text()

    if IsBufferTimestamp(buffer) {
      timestamp, _ := strconv.ParseInt(buffer[1:len(buffer)], 10, 64)
      tm := time.Unix(timestamp, 0)

      item.Timestamp = tm
    } else {
      item.Command = buffer
      items[i] = item
      i++
    }
	}

  sort.Sort(ByTimestamp(items))
  fmt.Println(items)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
